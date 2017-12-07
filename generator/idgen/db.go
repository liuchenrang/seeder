package idgen

import (
	"database/sql"
	"fmt"
	"log"
	"seeder/bootstrap"
	"sync"

	"github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
)

type DBGen struct {
	db   *sql.DB
	lock *sync.Mutex
	Fin  chan<- int

	application *bootstrap.Application
}

var (
	db *sql.DB
)

func (this *DBGen) GenerateSegment(bizTag string) (currentId uint64, cacheSteop uint64, step uint64, e error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	currentId, cacheSteop, step, e = this.find(bizTag)
	return currentId, cacheSteop, step, e
}
func (this *DBGen) flush(bizTag string) {
	this.UpdateStep(bizTag)
}
func (this *DBGen) getLogger() log4go.Logger {
	return this.application.Get("globalLogger").(log4go.Logger)
}
func (this *DBGen) find(bizTag string) (currentId uint64, cacheStep uint64, step uint64, e error) {

	tx, errBegin := this.db.Begin()

	sqlSelect := "SELECT currentId,cacheStep,step from " + this.application.GetConfig().Database.Account.Table + " where keyName= ? FOR UPDATE"
	stmt, errPrepare := this.db.Prepare(sqlSelect)
	defer stmt.Close()
	if errPrepare != nil {
		log.Fatal(errBegin.Error())
	}
	stmt.Exec(bizTag)
	if errBegin != nil {
		log.Fatal(errBegin.Error())
	}
	errQuery := stmt.QueryRow(bizTag).Scan(&currentId, &cacheStep, &step)
	if errQuery != nil {
		panic(errQuery.Error()) // proper error handling instead of panic in your app
	}
	tx.Commit()
	this.getLogger().Debug("DBGen find ", sqlSelect, "currentId", currentId, "cacheStep", cacheStep)
	return currentId, cacheStep, step, errQuery
}
func (this *DBGen) UpdateStep(bizTag string) (int64, error) {

	stmt, errPrepare := this.db.Prepare("UPDATE " + this.application.GetConfig().Database.Account.Table + " SET currentId = currentId + cacheStep where keyName= ? ")
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
		var errOpen error
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
