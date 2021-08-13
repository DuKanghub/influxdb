package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"w2influx/model"
	v2 "w2influx/service/influx/v2"
)

func GetDbSize(dbInfo model.DbInfo) (dbSize []model.DbSize, err error) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True", dbInfo.User, dbInfo.PassWord, dbInfo.Host, dbInfo.Port)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return []model.DbSize{}, err
	}
	err = db.Ping()
	if err != nil {
		return []model.DbSize{}, err
	}
	sqlStr := "select table_schema as 'DbName',table_name as 'TableName',data_length as 'DataSize',index_length as 'IndexSize', data_length + index_length as 'TotalSize' FROM information_schema.TABLES where table_schema not in ('sys','information_schema','performance_schema') order by data_length + index_length desc"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return []model.DbSize{}, err
	}
	defer db.Close()
	defer rows.Close()
	var dbs model.DbSize
	for rows.Next() {
		rows.Scan(&dbs.DbName, &dbs.TableName, &dbs.DataSize, &dbs.IndexSize, &dbs.TotalSize)
		v2.WritePoint(dbInfo, dbs) //一次写一条记录
		dbSize = append(dbSize, dbs)
	}

	return
}

func DbSize2Influx() {
	dbm, err := sql.Open("mysql", "db_info:w6R2gK4R8hPN@tcp(127.0.0.1:3306)/db_info?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	sqlStr := "select name,host,port,user,password from db_info"
	rows, err := dbm.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbm.Close()
	defer rows.Close()
	var db model.DbInfo
	for rows.Next() {
		rows.Scan(&db.Name, &db.Host, &db.Port, &db.User, &db.PassWord)
		GetDbSize(db)
	}
}