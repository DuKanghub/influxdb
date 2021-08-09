package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"w2influx/model"
	"w2influx/service"
)


func main() {

	dbm, err := sql.Open("mysql", "db_info:w6R2gK4R8hPN@tcp(127.0.0.1:3306)/db_info?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	sqlStr := "select name,host,port,user,password from db_info"
	rows, err := dbm.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	defer dbm.Close()
	defer rows.Close()
	var db model.DbInfo
	for rows.Next() {
		rows.Scan(&db.Name, &db.Host, &db.Port, &db.User, &db.PassWord)
		service.GetDbSize(db)

	}

}