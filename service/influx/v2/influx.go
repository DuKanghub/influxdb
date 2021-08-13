package v2

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
	"w2influx/model"
)

const Token = "fsDpF0BLJcg4MxSoffbLoBGJZRzmZT5JuYapUm96eG35vplSprOYz8kaDvCCovG24iOIJiVJ6y2CdM8cNYAPhw=="
const Bucket = "db_size"
const Org = "dukanghub"
const ServerUrl = "http://192.168.128.128:8086"


func WritePoint(db model.DbInfo, data model.DbSize) {
	client := influxdb2.NewClient(ServerUrl, Token)
	defer client.Close()
	// get non-blocking write client
	writeAPI := client.WriteAPI(Org, Bucket)
	p := influxdb2.NewPointWithMeasurement("db_size")
	p.AddTag("dbInstance", db.Name).AddTag("dbName", data.DbName).AddTag("tableName", data.TableName).
		AddField("dataSize", data.DataSize).AddField("indexSize", data.IndexSize).AddField("totalSize", data.TotalSize).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()

}
