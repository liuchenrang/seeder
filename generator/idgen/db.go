package idgen

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"seeder/config"
	"fmt"
	"seeder/logger"
)

type DBGen struct {
	maxId     uint64
	db        *sql.DB
	cacheStep uint64
	lock      *sync.Mutex
	Fin       chan<- int
	config    config.SeederConfig

	SeederLogger.Logger
}

var (
	db *sql.DB
)

func (dbgen *DBGen) GenerateSegment(bizTag string) (uint64, uint64, error) {
	dbgen.lock.Lock()
	defer dbgen.lock.Unlock()
	dbgen.flush(bizTag)
	dbgen.Debug("DBGen GenerateSegment %+v", dbgen)

	return dbgen.maxId, dbgen.cacheStep, nil
}
func (dbgen *DBGen) flush(bizTag string) {
	dbgen.find(bizTag)
	dbgen.UpdateStep(bizTag)
}
func (dbgen *DBGen) find(bizTag string) {
	dbgen.Debug("DBGen Find %+v", *dbgen)

	tx, errBegin := dbgen.db.Begin()
	dbgen.Debug("DBGen find concif %+v ", dbgen.config)

	sqlSelect := "SELECT currentId,cacheStep from " + dbgen.config.Database.Account.Table + " where keyName= ? FOR UPDATE"
	dbgen.Debug("DBGen find ", sqlSelect)
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
	dbgen.cacheStep = cacheStep
	dbgen.maxId = currentId + 1
}
func (dbgen *DBGen) UpdateStep(bizTag string) (int64, error) {
	dbgen.Debug("DBGen UpdateStep %+v", dbgen)

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
func init() {

}
func NewDBGen(bizTag string, config config.SeederConfig) IDGen {
	if db == nil {
		var errOpen error;
		//
		//dsn := fmt.Sprintf(
		//	"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		//	config.Database.Account.Name,
		//	config.Database.Account.Password,
		//	config.Database.Master[0].Host,
		//	config.Database.Master[0].Port,
		//	config.Database.Account.DBName,
		//)
		dsn := "root:tortdh_gogo888!@tcp(10.10.106.218:3306)/maindb?charset=utf8"
		fmt.Printf(dsn)
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
	dbGen := &DBGen{db: db, lock: &sync.Mutex{}, config: config}
	return dbGen
}
