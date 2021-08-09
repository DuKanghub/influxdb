package service

import (
	"database/sql"
	"fmt"
	"w2influx/model"
	v1 "w2influx/service/influx/v1"
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
		//fmt.Printf("dbName: %s, tableName: %s, dataSize: %d, indexSize: %d, totalSize: %d\n", dbs.DbName, dbs.TableName, dbs.DataSize, dbs.IndexSize, dbs.TotalSize)
		//v1.WritePoint(dbInfo, dbs) //一次写一条记录
		dbSize = append(dbSize, dbs)
	}
	// 批量写
	v1.WritePoints(dbInfo, dbSize)

	return
}
