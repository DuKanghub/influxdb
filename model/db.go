package model

type DbInfo struct {
	Name string	//实例名字
	Host string
	Port string
	User string
	PassWord string
}

type DbSize struct {
	DbName    string
	TableName string
	DataSize  int
	IndexSize int
	TotalSize int
}
