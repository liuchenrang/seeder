package idgen

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"fmt"
	"seeder/bootstrap"
	"github.com/alecthomas/log4go"
)

type DBGen struct {
	maxId     uint64
	db        *sql.DB
	cacheStep uint64
	lock      *sync.Mutex
	Fin       chan<- int

	application *bootstrap.Application
	 
}

var (
	db *sql.DB
)

func (dbgen *DBGen) GenerateSegment(bizTag string) (uint64, uint64, error) {
	dbgen.lock.Lock()
	defer dbgen.lock.Unlock()
	dbgen.find(bizTag)

	return dbgen.maxId, dbgen.cacheStep, nil
}
func (dbgen *DBGen) flush(bizTag string) {
	dbgen.UpdateStep(bizTag)
}
func (dbgen *DBGen)  getLogger() log4go.Logger {
	return dbgen.application.Get("globalLogger").(log4go.Logger)
}
func (dbgen *DBGen) find(bizTag string) {

	tx, errBegin := dbgen.db.Begin()

	sqlSelect := "SELECT currentId,cacheStep from " + dbgen.application.GetConfig().Database.Account.Table + " where keyName= ? FOR UPDATE"
	stmt, errPrepare := dbgen.db.Prepare(sqlSelect)
	defer stmt.Close()
	if errPrepare != nil {
		log.Fatal(errBegin.Error())
	}
	stmt.Exec(bizTag)
	if errBegin != nil {
		log.Fatal(errBegin.Error())
	}
	var currentId, cacheStep uint64
	errQuery := stmt.QueryRow(bizTag).Scan(&currentId, &cacheStep)
	if errQuery != nil {
		panic(errQuery.Error()) // proper error handling instead of panic in your app
	}
	tx.Commit()
	dbgen.getLogger().Debug("DBGen find ", sqlSelect,"currentId", currentId, "cacheStep", cacheStep)

	dbgen.cacheStep = cacheStep
	dbgen.maxId = currentId + 1
}
func (dbgen *DBGen) UpdateStep(bizTag string) (int64, error) {

	stmt, errPrepare := dbgen.db.Prepare("UPDATE " + dbgen.application.GetConfig().Database.Account.Table + " SET currentId = currentId + cacheStep where keyName= ? ")
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
func init() {

}
func NewDBGen(bizTag string, application *bootstrap.Application) IDGen {
	if db == nil {
		var errOpen error;
		//
		config := application.GetConfig()
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8",
			config.Database.Account.Name,
			config.Database.Account.Password,
			config.Database.Master[0].Host,
			config.Database.Master[0].Port,
			config.Database.Account.DBName,
		)
		//dsn := "root:tortdh_gogo888!@tcp(10.10.106.218:3306)/maindb?charset=utf8"
		db, errOpen = sql.Open("mysql", dsn) //
		if db == nil {
			if errOpen != nil {
				fmt.Println("error open")
				log.Fatal(errOpen)
			}
		}
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
	}
	dbGen := &DBGen{db: db, lock: &sync.Mutex{}, application: application}
	return dbGen
}
