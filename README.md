# influxdb简单使用示例
读取mysql数据库实例信息写入到influxdb
## master分支是写入到influxdb2
## v1分支是写入到influxdb1
# 使用方法
## 1. 创建数据库db_info, db_info表
```sql
DROP TABLE IF EXISTS `db_info`;
CREATE TABLE `db_info`  (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `port` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
```
## 2. 创建数据库账号并授权，可以直接使用root账号
用户名密码自己根据实际情况设置，跟DbSize2Influx里面的能对上就行。
```sql
grant all on db_info.* to 'db_info'@'%' identified by 'w6R2gK4R8hPN';
flush privileges;
```
## 3. 写入需要采集的数据库连接信息到db_info表里
## 4. docker启动influxdb2
## 5. 运行本程序
```sh
go run main.go
```
