package db_test

import (
	"encoding/json"
	"testing"

	"github.com/goexl/db"
	"github.com/stretchr/testify/assert"
)

func TestTypeMarshalJSON(t *testing.T) {
	testcases := []struct {
		in       db.Type
		expected string
	}{{
		in:       db.TypeMysql,
		expected: "mysql",
	}, {
		in:       db.TypePostgres,
		expected: "postgres",
	}, {
		in:       db.TypeSQLite,
		expected: "sqlite",
	}, {
		in:       db.TypeSQLite3,
		expected: "sqlite3",
	}, {
		in:       db.TypeOracle,
		expected: "oracle",
	}, {
		in:       db.TypeSQLServer,
		expected: "sqlserver",
	}}
	for _, testcase := range testcases {
		if result, me := json.Marshal(testcase.in); nil == me {
			assert.Equal(t, result, result)
		}
	}
}

func TestTypeUnmarshalJSONString(t *testing.T) {
	testcases := []struct {
		in       string
		expected db.Type
	}{{
		in:       "mysql",
		expected: db.TypeMysql,
	}, {
		in:       "postgres",
		expected: db.TypePostgres,
	}, {
		in:       "sqlite",
		expected: db.TypeSQLite,
	}, {
		in:       "sqlite3",
		expected: db.TypeSQLite3,
	}, {
		in:       "oracle",
		expected: db.TypeOracle,
	}, {
		in:       "sqlserver",
		expected: db.TypeSQLServer,
	}}
	for _, testcase := range testcases {
		result := new(db.Type)
		if ue := json.Unmarshal([]byte(testcase.in), result); nil == ue {
			assert.Equal(t, result, result)
		}
	}
}

func TestTypeUnmarshalJSONUint(t *testing.T) {
	testcases := []struct {
		in       uint8
		expected db.Type
	}{{
		in:       1,
		expected: db.TypeMysql,
	}, {
		in:       2,
		expected: db.TypePostgres,
	}, {
		in:       3,
		expected: db.TypeSQLite,
	}, {
		in:       4,
		expected: db.TypeSQLite3,
	}, {
		in:       5,
		expected: db.TypeOracle,
	}, {
		in:       6,
		expected: db.TypeSQLServer,
	}}
	for _, testcase := range testcases {
		result := new(db.Type)
		if ue := json.Unmarshal([]byte{testcase.in}, result); nil == ue {
			assert.Equal(t, result, result)
		}
	}
}
