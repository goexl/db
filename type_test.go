package db_test

import (
	"testing"

	"github.com/goexl/db"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.Equal(t, db.Type(1), db.TypeMysql)
	assert.Equal(t, db.Type(2), db.TypePostgres)
	assert.Equal(t, db.Type(3), db.TypeSQLite)
	assert.Equal(t, db.Type(4), db.TypeOracle)
}
