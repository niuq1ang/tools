package main

import (
	"flag"
	"log"
	"strings"

	"github.com/CodesInvoker/tools/common/db"
	"github.com/CodesInvoker/tools/common/timestamp"
	"github.com/CodesInvoker/tools/common/uuid"

	"gopkg.in/gorp.v1"
)

/*
	往表里插入数据工具
	用例：go run source.go -user root -password liuyexing -db project_f6005 -table team -count 10000
	说明：往project_f6005库的team表插入10000条假数据
*/

var (
	sqldb     *db.MySQLDB
	dbName    string
	tableName string
	columns   []string

	recordCount int

	now        = timestamp.GetSec()
	batchCount = 10000
)

func init() {
	c := new(db.MySQLConfig)
	c.BindFlag()
	flag.StringVar(&tableName, "table", "", "MySQL table name")
	flag.IntVar(&recordCount, "count", 1, "build record count")

	flag.Parse()
	dbName = c.DBName
	var err error
	sqldb, err = db.InitMySQL(c)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	tableSchemas, err := getTableSchema(sqldb.DBM)
	if err != nil {
		log.Fatalln("get table schema error", err)
	}
	for _, s := range tableSchemas {
		columns = append(columns, s.ColumnName)
	}
	records := buildRecords(tableSchemas)
	var index int
	var finishCount int
	for i := 0; i < len(records)/batchCount; i++ {
		index = i * batchCount
		parts := records[index:(index + batchCount)]
		err = addRecordInTransaction(parts)
		if err != nil {
			log.Fatalln("transact error", err)
		}
		finishCount += len(parts)
	}
	if len(records) > finishCount {
		parts := records[finishCount:]
		err = addRecordInTransaction(parts)
		if err != nil {
			log.Fatalln("transact error", err)
		}
	}
}

type TableSchema struct {
	ColumnName    string      `db:"column_name"`
	DataType      string      `db:"data_type"`
	ColumnDefault interface{} `db:"column_default"`
}

func getTableSchema(src gorp.SqlExecutor) ([]*TableSchema, error) {
	result := make([]*TableSchema, 0)
	sql := `SELECT column_name, data_type, column_default FROM information_schema.columns WHERE table_schema=? AND table_name=?;`
	_, err := src.Select(&result, sql, dbName, tableName)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func buildRecords(tableSchemas []*TableSchema) [][]interface{} {
	records := make([][]interface{}, 0)
	for i := 0; i < recordCount; i++ {
		record := make([]interface{}, 0)
		for _, s := range tableSchemas {
			if strings.HasSuffix(s.ColumnName, "uuid") {
				record = append(record, uuid.UUID())
			} else if strings.HasSuffix(s.ColumnName, "time") {
				record = append(record, now)
			} else {
				switch s.DataType {
				case "varchar", "text", "longtext":
					record = append(record, "test")
				case "tinyint", "int", "bigint":
					record = append(record, 1)
				default:

					record = append(record, s.ColumnDefault)
				}
			}
		}
		records = append(records, record)
	}
	return records
}

func addRecordInTransaction(records [][]interface{}) error {
	return db.MySQLTransact(sqldb.DBM, func(tx *gorp.Transaction) error {
		return addRecord(tx, records)
	})
}

func addRecord(tx *gorp.Transaction, records [][]interface{}) error {
	err := db.SqlBatchInsert(tx, tableName, columns, false, len(records), func(index int) []interface{} {
		return records[index]
	})
	if err != nil {
		return err
	}
	return nil
}
