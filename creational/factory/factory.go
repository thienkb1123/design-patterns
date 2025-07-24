/*
	Context: The application requires two database instances: postgres and mysql.
	Requirement:
		- Each type implements the database interface.
		- The Factory takes the database name as input and returns corresponding instance.
*/

package factory

import (
	"database/sql"
	"errors"
)

type DBType string

const (
	MySQL      DBType = "mysql"
	PostgreSQL DBType = "postgres"
)

type Database interface {
	Client() (sql.DB, error)
}

func NewDatabase(dbType DBType) (Database, error) {
	switch dbType {
	case MySQL:
		return NewMysql(), nil
	case PostgreSQL:
		return NewPostgres(), nil
	default:
		return nil, errors.New("db type not supported")
	}
}

type Postgres struct {
}

func NewPostgres() Database {
	return &Postgres{}
}

func (p *Postgres) Client() (sql.DB, error) {
	return sql.DB{}, nil
}

type Mysql struct {
}

func NewMysql() Database {
	return &Mysql{}
}

func (p *Mysql) Client() (sql.DB, error) {
	return sql.DB{}, nil
}
