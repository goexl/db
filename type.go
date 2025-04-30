package db

const (
	TypeMysql Type = iota + 1
	TypePostgres
	TypeSQLite
	TypeOracle
)

// Type 数据库类型
type Type uint8
