package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/influxdata/influxdb-client-go/v2"
	"log"
	"time"
)

var (
	Db *sql.DB
)

type User struct {
	DbName    string
	TableName string
	DataSize  int
	IndexSize int
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	//fmt.Println("connnect success", db.Ping())


	sqlStr := "select table_schema as 'DbName',table_name as 'TableName',data_length as 'DataSize',index_length as 'IndexSize' from information_schema.tables where table_schema='mysql' order by data_length desc, index_length desc;"
	//sqlStr := "select * from information_schema.tables;"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	// 关闭rows释放持有的数据库链接
	defer rows.Close()
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "fsDpF0BLJcg4MxSoffbLoBGJZRzmZT5JuYapUm96eG35vplSprOYz8kaDvCCovG24iOIJiVJ6y2CdM8cNYAPhw=="
	const bucket = "dukang-bucket"
	const org = "dukanghub"
	client := influxdb2.NewClient("http://192.168.128.128:8086", token)
	// always close client at the end
	defer client.Close()
	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	p := influxdb2.NewPointWithMeasurement("db-int")
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.DbName, &u.TableName, &u.DataSize, &u.IndexSize)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("DbName:%s TableName:%s DataSize:%d IndexSize:%d\n", u.DbName, u.TableName, u.DataSize, u.IndexSize)
		p.AddTag("dbName", u.DbName).AddTag("tbName", u.TableName).
		AddField("dataSize", u.DataSize).AddField("indexSize", u.IndexSize).
		SetTime(time.Now())
		writeAPI.WritePoint(p)
		writeAPI.Flush()
	}


}