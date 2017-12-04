package idgen

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"seeder/config"
	"fmt"
)

type DBGen struct {
	maxId     uint64
	db        *sql.DB
	cacheStep uint64
	lock      *sync.Mutex
	Fin       chan<- int
	config config.SeederConfig


}

var (
	db *sql.DB
)

func (dbgen *DBGen) GenerateSegment(bizTag string) (uint64, uint64, error) {
	return dbgen.maxId, dbgen.cacheStep, nil
}
func (dbgen *DBGen) flush(bizTag string) {
	dbgen.find(bizTag)
}
func (dbgen *DBGen) find(bizTag string) {
	dbgen.lock.Lock()
	defer dbgen.lock.Unlock()

	tx, errBegin := dbgen.db.Begin()
	stmt, errPrepare := dbgen.db.Prepare("SELECT currentId,cacheStep from " + dbgen.config.Database.Account.Table + " where keyName= ? FOR UPDATE")
	defer stmt.Close()
	if errPrepare != nil {
		log.Fatal(errBegin)
	}
	stmt.Exec(bizTag)
	if errBegin != nil {
		log.Fatal(errBegin)
	}
	var currentId, cacheStep uint64
	errQuery := stmt.QueryRow(bizTag).Scan(&currentId, &cacheStep)
	if errQuery != nil {
		panic(errQuery) // proper error handling instead of panic in your app
	}
	tx.Commit()
	dbgen.cacheStep = cacheStep
	dbgen.maxId = currentId + cacheStep
}
func (dbgen *DBGen) UpdateStep(bizTag string) (int64, error) {
	dbgen.lock.Lock()
	defer dbgen.lock.Unlock()

	stmt, errPrepare := dbgen.db.Prepare("UPDATE " + dbgen.config.Database.Account.Table + " SET currentId = currentId + cacheStep where keyName= ? ")
	var errorUpdate error
	defer stmt.Close()
	if errPrepare != nil {
		errorUpdate = errPrepare
	}
	result, errorExec := stmt.Exec(bizTag)
	if errorExec != nil {
		errorUpdate = errorExec
		log.Fatal(errorExec)
	}
	affected, errorRwos := result.RowsAffected()
	if errorRwos != nil {
		errorUpdate = errorRwos
	}
	return affected, errorUpdate
}
func init()  {

}
func NewDBGen(bizTag string, config config.SeederConfig) IDGen {
	if db == nil {
		var errOpen error;

		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8",
			config.Database.Account.Name,
				config.Database.Account.Password,
				config.Database.Master[0].Host,
				config.Database.Master[0].Port,
				config.Database.Account.DBName,
			)
		db, errOpen = sql.Open("mysql", "root:tortdh_gogo888!@tcp(10.10.106.218:3306)/maindb?charset=utf8") //
		if db == nil {
			if errOpen != nil {
				log.Fatal(errOpen)
			}
		}
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
	}
	dbGen := &DBGen{db: db, lock: &sync.Mutex{}, config: config}
	dbGen.find(bizTag)
	return dbGen
}
