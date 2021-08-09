package v1

import (
	"fmt"
	"log"
	"time"
	"github.com/influxdata/influxdb1-client/v2"
	"w2influx/model"
)

func ConnectInflux() (cli client.Client, err error){

	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.159.136:8087",
		Username: "admin",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return
}

func WritePoint(db model.DbInfo, data model.DbSize) (err error) {
	c, err := ConnectInflux()
	if err != nil {
		return err
	}
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "db_size",
		Precision: "ms",   //时间精度，记录时间戳的精度，有秒s,毫秒ms,微秒us,一般s和ms基本够用了。
	})

	tags := map[string]string{
		"dbInstance": db.Name,
		"dbName": data.DbName,
		"tableName": data.TableName,
	}
	fields := map[string]interface{}{
		"dataSize": data.DataSize,
		"indexSize": data.IndexSize,
		"totalSize": data.TotalSize,
	}
	pt, err := client.NewPoint("db_size", tags, fields, time.Now())
	if err != nil {
		fmt.Println(err)
		return err
	}
	bp.AddPoint(pt)
	c.Write(bp)
	return err
}

func WritePoints(db model.DbInfo, datas []model.DbSize) (err error) {
	c, err := ConnectInflux()
	if err != nil {
		return err
	}
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "db_size",
		Precision: "ms",	//时间精度，记录时间戳的精度，有秒s,毫秒ms,微秒us,一般s和ms基本够用了。
	})
	var pts []*client.Point
	for _, data := range datas {
		tags := map[string]string{
			"dbInstance": db.Name,
			"dbName": data.DbName,
			"tableName": data.TableName,
		}
		fields := map[string]interface{}{
			"dataSize": data.DataSize,
			"indexSize": data.IndexSize,
			"totalSize": data.TotalSize,
		}
		pt, _ := client.NewPoint("db_size", tags, fields, time.Now())
		pts = append(pts, pt)

	}

	bp.AddPoints(pts)
	c.Write(bp)
	return err
}