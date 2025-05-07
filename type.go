package db

import (
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
)

const (
	TypeMySQL     Type = "mysql"
	TypePostgres  Type = "postgres"
	TypeSQLite    Type = "sqlite"
	TypeSQLite3   Type = "sqlite3"
	TypeOracle    Type = "oracle"
	TypeSQLServer Type = "sqlserver"
)

// Type 数据库类型
type Type string

func (t *Type) String() (result string) {
	switch *t {
	case TypeMySQL:
		result = "mysql"
	case TypePostgres:
		result = "postgres"
	case TypeSQLite:
		result = "sqlite"
	case TypeSQLite3:
		result = "sqlite3"
	case TypeOracle:
		result = "oracle"
	case TypeSQLServer:
		result = "sqlserver"
	default:
		result = "unknown"
	}

	return
}

func (t *Type) MarshalJSON() (bytes []byte, err error) {
	switch *t {
	case TypeMySQL:
		bytes = []byte("mysql")
	case TypePostgres:
		bytes = []byte("postgres")
	case TypeSQLite:
		bytes = []byte("sqlite")
	case TypeSQLite3:
		bytes = []byte("sqlite3")
	case TypeOracle:
		bytes = []byte("oracle")
	case TypeSQLServer:
		bytes = []byte("sqlserver")
	default:
		err = exception.New().Message("未被支持的数据库").Field(field.New("type", *t)).Build()
	}

	return
}

func (t *Type) Unmarshal(bytes []byte) error {
	return t.UnmarshalJSON(bytes)
}

func (t *Type) UnmarshalJSON(bytes []byte) (err error) {
	switch strings.ToLower(string(bytes)) {
	case "mysql":
		*t = TypeMySQL
	case "postgres":
		*t = TypePostgres
	case "sqlite":
		*t = TypeSQLite
	case "sqlite3":
		*t = TypeSQLite3
	case "oracle":
		*t = TypeOracle
	case "sqlserver":
		*t = TypeSQLServer
	default:
		err = exception.New().Message("未被支持的数据库").Field(field.New("type", *t)).Build()
	}

	return
}
