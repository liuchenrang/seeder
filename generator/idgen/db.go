package idgen

import (
	"database/sql"
	"fmt"
	"log"
	"seeder/bootstrap"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DBGen struct {
	db   *sql.DB
	lock *sync.Mutex
	Fin  chan<- int

	application *bootstrap.Application
}

var (
	DB   *sql.DB
	muDB sync.Mutex
)

func getDB(application *bootstrap.Application) *sql.DB {
	muDB.Lock()
	defer muDB.Unlock()
	config := application.GetConfig()
	ll := len(config.Database.Master)
	fmt.Println(ll)
	if DB != nil {
		error := DB.Ping()
		if error != nil {
			DB = nil
		}
	}
	if DB == nil {
		var errOpen error
		for _, mst := range config.Database.Master {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8",
				config.Database.Account.Name,
				config.Database.Account.Password,
				mst.Host,
				mst.Port,
				config.Database.Account.DBName,
			)
			DB, errOpen = sql.Open("mysql", dsn) //
			error := DB.Ping()
			if error == nil {
				break
			}
			if DB == nil {
				if errOpen != nil {
					log.Fatal(errOpen)
				}
			}
		}

		DB.SetMaxOpenConns(config.Database.ConnectionInfo.MaxOpenConns)
		DB.SetMaxIdleConns(config.Database.ConnectionInfo.MaxIdleConns)
	}
	return DB
}
func (this *DBGen) GenerateSegment(bizTag string) (currentId uint64, cacheSteop uint64, step uint64, e error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	currentId, cacheSteop, step, e = this.Find(bizTag)
	return currentId, cacheSteop, step, e
}
func (this *DBGen) flush(bizTag string) {
	this.UpdateStep(bizTag)
}

func (this *DBGen) Find(bizTag string) (currentId uint64, cacheStep uint64, step uint64, e error) {

	tx, errBegin := this.db.Begin()
	defer tx.Commit()

	sqlSelect := "SELECT currentId,cacheStep,step from " + this.application.GetConfig().Database.Account.Table + " where keyName= ? FOR UPDATE"
	stmt, errPrepare := this.db.Prepare(sqlSelect)
	defer stmt.Close()

	if errPrepare != nil {
		this.application.GetLogger().Error(errBegin.Error())
		log.Fatal(errBegin.Error())
	}
	stmt.Exec(bizTag)
	if errBegin != nil {
		this.application.GetLogger().Error(errBegin.Error())

		log.Fatal(errBegin.Error())
	}
	errQuery := stmt.QueryRow(bizTag).Scan(&currentId, &cacheStep, &step)
	if errQuery != nil {
		this.application.GetLogger().Error(errQuery.Error())
		panic(errQuery.Error()) // proper error handling instead of panic in your app
	}
	this.UpdateStep(bizTag)
	this.application.GetLogger().Debug("DBGen Find ", sqlSelect, "currentId", currentId, "cacheStep", cacheStep, "bizTag", bizTag)
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

	dbGen := &DBGen{db: getDB(application), lock: &sync.Mutex{}, application: application}
	return dbGen
}
