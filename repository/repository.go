package repository

import (
	"context"
	gosql "database/sql"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/migrate"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewEnt() (*gosql.DB, *ent.Client) {
	// parseTime=true가 없을 시
	// Error: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	// ref: https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	drv, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=%s",
			config.Config.DB.MySQL.User,
			config.Config.DB.MySQL.Password,
			config.Config.DB.MySQL.Host,
			config.Config.DB.MySQL.DatabaseName,
			"Asia%2FSeoul",
		))
	if err != nil {
		log.Panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)

	conn, err := db.Conn(context.TODO())
	if err != nil {
		log.Panic(err)
	}
	conn.Close()

	db.SetConnMaxLifetime(time.Hour)
	ent.Debug()
	ent.Log(func(i ...interface{}) {
		log.Warn(i...)
	})
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	err = client.Schema.Create(context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(false),
		migrate.WithForeignKeys(true),
	)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db, client
}

type SynchronousCacheWrite bool
