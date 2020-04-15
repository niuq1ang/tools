package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/bangwork/bang-api/migration/common/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivpusic/grpool"
	"gopkg.in/gorp.v1"
)

const (
	DSNPattern = `%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FShanghai&charset=utf8mb4`
)

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Verbose  bool
}

func (c *MySQLConfig) BindFlag() {
	flag.StringVar(&c.Host, "host", "127.0.0.1", "MySQL host")
	flag.IntVar(&c.Port, "port", 3306, "MySQL port")
	flag.StringVar(&c.User, "user", "", "MySQL user")
	flag.StringVar(&c.Password, "password", "", "MySQL password")
	flag.StringVar(&c.DBName, "db", "bang", "MySQL database name")
	flag.BoolVar(&c.Verbose, "verbose", false, "Enable verbose log")
}

type MySQLDB struct {
	DB  *sql.DB
	DBM *gorp.DbMap
}

func InitMySQL(c *MySQLConfig) (*MySQLDB, error) {
	dsn := fmt.Sprintf(DSNPattern, c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	dbm := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}
	if c.Verbose {
		dbm.TraceOn("[SQL]", log.New(os.Stdout, "", log.Lmicroseconds))
	} else {
		dbm.TraceOff()
	}
	result := &MySQLDB{
		DB:  db,
		DBM: dbm,
	}
	return result, nil
}

func MySQLTransact(dbm *gorp.DbMap, txFunc func(tx *gorp.Transaction) error) (err error) {
	tx, err := dbm.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			log.Printf("%s: %s\n", p, debug.Stack())
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	return txFunc(tx)
}

func MySQLSimpleFlow(f func(tx *gorp.Transaction) error) {
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
	err = MySQLTransact(sqldb.DBM, f)
	if err != nil {
		log.Fatalln(err)
	}
}

func MySQLInterruptableMigration(alter string, migrate func(tx *gorp.Transaction) error, interFunc func(db *gorp.DbMap)) {
	var onceFunc sync.Once
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
	if len(alter) > 0 {
		err = MySQLTransact(sqldb.DBM, func(tx *gorp.Transaction) error {
			return MySQLExecMultipleStatements(tx, alter)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}

	utils.KillSignal(func() {
		onceFunc.Do(func() {
			if interFunc != nil {
				interFunc(sqldb.DBM)
			}
			log.Fatalln("kill signal.")
		})
	})

	err = MySQLTransact(sqldb.DBM, func(tx *gorp.Transaction) error {
		err = migrate(tx)
		return err
	})
	if err != nil {
		onceFunc.Do(func() {
			if interFunc != nil {
				interFunc(sqldb.DBM)
			}
		})
		log.Fatalln(err)
	}
}

func MySQLMigrationFlow(alter string, migrate func(tx *gorp.Transaction) error) {
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
	if len(alter) > 0 {
		err = MySQLTransact(sqldb.DBM, func(tx *gorp.Transaction) error {
			return MySQLExecMultipleStatements(tx, alter)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = MySQLTransact(sqldb.DBM, migrate)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewDBM() *gorp.DbMap {
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
	return sqldb.DBM
}

// MySQLMigrationsFlow 把迁移拆分为多个事物进行迁移
// 注意：使用查分事物的方式进行迁移时，要尽量保证迁移是可重复执行的
// routineCount 并发协程数， 非特殊情况下设置为1， 如果有需要建议可以设置为2或4
func MySQLMigrationsFlow(dbm *gorp.DbMap, routineCount int, alter string, migrates ...func(tx *gorp.Transaction) error) {
	if len(alter) > 0 {
		err := MySQLTransact(dbm, func(tx *gorp.Transaction) error {
			return MySQLExecMultipleStatements(tx, alter)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	pool := grpool.NewPool(routineCount, 100)
	defer pool.Release()
	count := len(migrates)
	pool.WaitCount(count)
	for _, migrate := range migrates {
		mySQLTransactWithPool(pool, dbm, migrate)
	}
	pool.WaitAll()
}

func mySQLTransactWithPool(pool *grpool.Pool, dbm *gorp.DbMap, migrate func(tx *gorp.Transaction) error) {
	pool.JobQueue <- func() {
		defer pool.JobDone()
		err := MySQLTransact(dbm, migrate)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MySQLExecMultipleStatements(tx *gorp.Transaction, rawquery string) error {
	// 处理 SQL 字符串中带有分号的特殊情况
	re := regexp.MustCompile(`[^\\];`)
	locs := re.FindAllStringIndex(rawquery, -1)
	prev := 0
	for _, loc := range locs {
		end := loc[1]
		query := strings.TrimSpace(rawquery[prev:end])
		if query == "" || query == ";" {
			continue
		}
		if _, err := tx.Exec(query); err != nil {
			return err
		}
		prev = end
	}
	return nil
}

func MySQLExecMultipleStatementsOnNative(tx *gorp.DbMap, rawquery string) error {
	// 处理 SQL 字符串中带有分号的特殊情况
	re := regexp.MustCompile(`[^\\];`)
	locs := re.FindAllStringIndex(rawquery, -1)
	prev := 0
	for _, loc := range locs {
		end := loc[1]
		query := strings.TrimSpace(rawquery[prev:end])
		if query == "" || query == ";" {
			continue
		}
		if _, err := tx.Exec(query); err != nil {
			return err
		}
		prev = end
	}
	return nil
}

// 无事务
func MySqlExecutor(sqlFunc func(src *MySQLDB) error) error {
	c := new(MySQLConfig)
	c.BindFlag()
	flag.Parse()
	sqldb, err := InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
	return sqlFunc(sqldb)
}
