package db

import (
	"context"
	"database/sql"
	"go-gql-sqlboiler-psql/config"
	"time"

	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

type DbUtils interface {
	GetDao(context.Context) boil.ContextExecutor
	Error(error) error
}

type dbUtil struct {
	db *sql.DB
}

var txKey struct{}

func Init() {
	connection, err := pq.ParseURL(config.Conf.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=" + config.Conf.SSLMode
	DB, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	if err := DB.Ping(); err != nil {
		panic("not db connection")
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxLifetime(300 * time.Second)

	boil.DebugMode = true
	boil.SetDB(DB)
}

func NewDbUtil(db *sql.DB) DbUtils {
	return &dbUtil{db: db}
}

func (u *dbUtil) GetDao(ctx context.Context) boil.ContextExecutor {
	tx, ok := ctx.Value(txKey).(*sql.Tx)
	if ok {
		return tx
	}
	return u.db
}

func (u *dbUtil) Error(err error) error {
	if err == nil || err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DoInTx(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, txKey, tx)
	result, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, nil
}
