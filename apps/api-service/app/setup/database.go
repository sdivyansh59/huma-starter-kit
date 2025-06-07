package setup

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"github.com/sdivyansh59/huma-project-starter/app/internal-lib/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

func InitializeDatabase() (*bun.DB, error) {
	dsn := utils.GetEnvOr("POSTGRES_DB_URL", "")
	var sqldb *sql.DB

	if dsn != "" {
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	} else {
		// Todo: set these value via environment variables
		pgconn := pgdriver.NewConnector(
			pgdriver.WithNetwork("tcp"),
			pgdriver.WithAddr("localhost:5437"),
			pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
			pgdriver.WithUser("test"),
			pgdriver.WithPassword("test"),
			pgdriver.WithDatabase("test"),
			pgdriver.WithApplicationName("myapp"),
			pgdriver.WithTimeout(5*time.Second),
			pgdriver.WithDialTimeout(5*time.Second),
			pgdriver.WithReadTimeout(5*time.Second),
			pgdriver.WithWriteTimeout(5*time.Second),
			pgdriver.WithConnParams(map[string]interface{}{
				"search_path": "my_search_path",
			}),
		)
		sqldb = sql.OpenDB(pgconn)
		if sqldb == nil {
			return nil, errors.New("failed to open database connection")
		}
	}
	// Configure connection pool settings
	sqldb.SetMaxOpenConns(20) // Adjust based on your needs
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxLifetime(time.Hour)

	db := bun.NewDB(sqldb, pgdialect.New())
	if db == nil {
		return nil, errors.New("failed to create bun DB instance")
	}

	return db, nil
}
