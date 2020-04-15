package restore

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bangwork/bang-api/migration/common/db"
	"github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)

type DataStore struct {
	dataWriterSet map[string]*csv.Writer
	storeSet      map[string]*os.File
}

func NewDataStore(tableNames []string) (*DataStore, error) {
	ds := new(DataStore)
	ds.storeSet = make(map[string]*os.File)
	ds.dataWriterSet = make(map[string]*csv.Writer)

	for _, tbName := range tableNames {
		filePath := tbName
		of, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, errors.Trace(err)
		}
		ds.storeSet[filePath] = of
		ds.dataWriterSet[filePath] = csv.NewWriter(of)
	}
	return ds, nil
}

func (d *DataStore) Close() {
	for _, of := range d.storeSet {
		of.Close()
	}
}

func (d *DataStore) Flush() {
	for _, w := range d.dataWriterSet {
		w.Flush()
	}
}

type RecordFunc func() []string

func (d *DataStore) StoreFile(tableName string, fn RecordFunc) error {
	filePath := tableName
	if _, ok := d.storeSet[filePath]; !ok {
		return errors.New("not found store file: " + filePath)
	}
	record := fn()
	err := d.dataWriterSet[filePath].Write(record)
	return errors.Trace(err)
}

func (d *DataStore) Restore(sqldb *db.MySQLDB) error {

	for f, _ := range d.storeSet {
		mysql.RegisterLocalFile(f)
		defer mysql.DeregisterLocalFile(f)
	}

	restoreFunc := func(tbName string) error {
		absFilePath, err := filepath.Abs(tbName)
		if err != nil {
			log.Printf("get %s abs path has err: %v \n", tbName, err)
			return errors.Trace(err)
		}
		if !d.hasData(absFilePath) {
			log.Printf("%s no data.\n", absFilePath)
			return nil
		}

		log.Printf("load data %s.\n", absFilePath)
		defer func() {
			log.Printf("load data %s done.\n", absFilePath)
		}()

		query := `LOAD DATA LOCAL INFILE '%s' INTO TABLE %s `
		query += `FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n';`
		query = fmt.Sprintf(query, tbName, tbName)
		if _, err := sqldb.DB.Exec(query); err != nil {
			return errors.Trace(err)
		}
		return nil
	}

	for f, _ := range d.storeSet {
		restoreFunc(f)
	}
	return nil
}

func (d *DataStore) hasData(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return info.Size() > 0
}
