package persistence

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/url"
	"strings"
	"time"
)

const ulrConnectionString = "postgresql://%s:%s@%s:%s/%s?sslmode=%s"

func GetPostgresClient(ctx context.Context, pgProp *PostgresProp) *PgClient {
	verifyPgProp(pgProp)

	urlCnx := fmt.Sprintf(
		ulrConnectionString, url.QueryEscape(pgProp.DbUser), url.QueryEscape(pgProp.DbPassword),
		url.QueryEscape(pgProp.DbHost), url.QueryEscape(pgProp.DbPort), url.QueryEscape(pgProp.DbName),
		url.QueryEscape(pgProp.DbSslMode),
	)

	cnxCfg, err := pgxpool.ParseConfig(urlCnx)
	if err != nil {
		log.Fatalf("Error creating pool config: %v\n", err)
	}
	cnxCfg.ConnConfig.RuntimeParams["timezone"] = GetTimeZone(pgProp.Country)
	var maxCnx = 30
	if pgProp.DbMaxCnx > 0 {
		maxCnx = pgProp.DbMaxCnx
	}
	cnxCfg.MaxConns = int32(maxCnx)
	dbPool, err := pgxpool.ConnectConfig(ctx, cnxCfg)

	if err != nil {
		log.Fatalf("Error creating connection: %v\n", err)
	}

	return &PgClient{ctx, dbPool}
}

func verifyPgProp(pgProp *PostgresProp) {
	if len(pgProp.DbUser) == 0 {
		log.Fatalf("User property is mandatory")
	}
	if len(pgProp.DbHost) == 0 {
		log.Fatal("DB_HOST not found")
	}
	if len(pgProp.DbPort) == 0 {
		log.Fatal("DB_PORT not found")
	}
	if len(pgProp.DbUser) == 0 {
		log.Fatal("DB_USER not found")
	}
	if len(pgProp.DbPassword) == 0 {
		log.Fatal("DB_PASSWORD not found")
	}
	if len(pgProp.DbName) == 0 {
		log.Fatal("DB_NAME not found")
	}
	v := pgProp.DbSslMode
	if len(v) > 0 && (!strings.EqualFold(v, "disable") && !strings.EqualFold(v, "require") &&
		!strings.EqualFold(v, "verify-ca") && !strings.EqualFold(v, "verify-full")) {
		log.Fatalf("DB_SSL_MODE invalid value: %s", v)
	} else if len(v) == 0 {
		pgProp.DbSslMode = "require"
	}
	if pgProp.DbSelectLimit == 0 {
		pgProp.DbSelectLimit = 5000
	}
	if pgProp.DbMaxCnx == 0 {
		pgProp.DbMaxCnx = 30
	}
	if pgProp.DbWaitAfterQuery == 0 {
		pgProp.DbWaitAfterQuery = time.Second * 10
	}
}

type PgClient struct {
	ctx     context.Context
	poolCnx *pgxpool.Pool
}

func (db *PgClient) RunQueryRows(query string) (pgx.Rows, error) {
	rows, err := db.poolCnx.Query(db.ctx, query)
	if err != nil {
		log.Printf("Error running query (%s): %v\n", query, err)
	}
	return rows, err
}

func (db *PgClient) RunQueryArgs(query string, args []interface{}) (pgx.Rows, error) {
	rows, err := db.poolCnx.Query(db.ctx, query, args...)
	if err != nil {
		log.Printf("Error running query (%s): %v\n", query, err)
	}
	return rows, err
}

func (db *PgClient) RunQueryRow(query string) *pgx.Row {
	row := db.poolCnx.QueryRow(db.ctx, query)
	return &row
}
