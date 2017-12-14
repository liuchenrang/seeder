package idgen

import (
	"database/sql"
	"fmt"
	"log"
	"seeder/bootstrap"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strings"
	"strconv"
)

type DBGen struct {
	muDB sync.Mutex
	db   *sql.DB

	application *bootstrap.Application
}
func GoId() int {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)     }
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
var (
	DB   *sql.DB
	muDB sync.Mutex

	muDBGen sync.Mutex

	dbGen *DBGen
)

func getDB(application *bootstrap.Application) *sql.DB {
	muDB.Lock()
	defer muDB.Unlock()
	config := application.GetConfig()
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

			if DB == nil {
				if errOpen != nil {
					log.Fatal(errOpen)
				}
			}
			if error == nil {
				break
			}
		}

		DB.SetMaxOpenConns(config.Database.ConnectionInfo.MaxOpenConns)
		DB.SetMaxIdleConns(config.Database.ConnectionInfo.MaxIdleConns)
	}
	return DB
}
func (this *DBGen) GenerateSegment(bizTag string) (currentId uint64, cacheSteop uint64, step uint64, e error) {
	this.muDB.Lock()
	defer this.muDB.Unlock()
	currentId, cacheSteop, step, e = this.Find(bizTag)
	return currentId, cacheSteop, step, e
}


func (this *DBGen) Find(bizTag string) (currentId uint64, cacheStep uint64, step uint64, e error) {

	tx, errBegin := this.db.Begin()
	defer tx.Commit()

	sqlSelect := "SELECT currentId,cacheStep,step from " + this.application.GetConfig().Database.Account.Table + " where keyName= ? FOR UPDATE"
	stmt, errPrepare := tx.Prepare(sqlSelect)
	defer stmt.Close()
	if errPrepare != nil {
		this.application.GetLogger().Error("DBGEN", errBegin.Error())
		log.Fatal(errBegin.Error())
	}
	stmt.Exec(bizTag)
	if errBegin != nil {
		this.application.GetLogger().Error("DBGEN",errBegin.Error())
	}
	errQuery := stmt.QueryRow(bizTag).Scan(&currentId, &cacheStep, &step)
	if errQuery != nil {
		this.application.GetLogger().Warn("DBGEN",errQuery.Error())
		return 0,0,0,nil
	}
	affected , e := this.UpdateStep(tx, bizTag)
	this.application.GetLogger().Info("DBGen Find ", sqlSelect, "currentId", currentId, "cacheStep", cacheStep, "bizTag", bizTag)
	if cacheStep > 0 {
		if affected  > 0 {
			return currentId, cacheStep, step, errQuery
		}else{
			panic(e)
		}
	}else{
		this.application.GetLogger().Error("DBGen UpdateStep Fail ", sqlSelect, "currentId", currentId, "cacheStep", cacheStep, "bizTag", bizTag)
		return currentId, cacheStep, step, errQuery
	}


}

func (this *DBGen) UpdateStep(tx *sql.Tx, bizTag string) (int64, error) {

	stmt, errPrepare := tx.Prepare("UPDATE " + this.application.GetConfig().Database.Account.Table + " SET currentId = currentId + cacheStep where keyName= ? ")
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

func NewDBGen(bizTag string, application *bootstrap.Application) IDGen {

	return &DBGen{db: getDB(application), application: application}
}
