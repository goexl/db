package db

import (
	"encoding/json"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
)

const (
	TypeMySQL Type = iota + 1
	TypePostgres
	TypeSQLite
	TypeSQLite3
	TypeOracle
	TypeSQLServer
)

// Type 数据库类型
type Type uint8

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

func (t *Type) UnmarshalJSON(bytes []byte) (err error) {
	num := new(uint8)
	if ne := json.Unmarshal(bytes, num); nil == ne {
		err = t.unmarshalUint(*num)
	} else {
		err = t.unmarshalString(string(bytes))
	}

	return
}

func (t *Type) unmarshalUint(value uint8) (err error) {
	switch value {
	case 1:
		*t = TypeMySQL
	case 2:
		*t = TypePostgres
	case 3:
		*t = TypeSQLite
	case 4:
		*t = TypeSQLite3
	case 5:
		*t = TypeOracle
	case 6:
		*t = TypeSQLServer
	default:
		err = exception.New().Message("未被支持的数据库").Field(field.New("type", *t)).Build()
	}

	return
}

func (t *Type) unmarshalString(value string) (err error) {
	switch strings.ToLower(value) {
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
