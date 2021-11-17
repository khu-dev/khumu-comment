package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/errorz"
	"github.com/sirupsen/logrus"
	"time"
)

func NewEnt() *ent.Client {
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
		logrus.Panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	conn, err := db.Conn(context.TODO())
	if err != nil {
		logrus.Panic(err)
	}
	conn.Close()

	db.SetConnMaxLifetime(time.Hour)
	ent.Debug()
	ent.Log(func(i ...interface{}) {
		logrus.Warn(i...)
	})
	client := ent.NewClient(ent.Driver(drv))
	return client
}

// NotFound에 대한 EntError를 우리 도메인의 에러 타입으로 변경
// 만약 여기서 감지되지 않은 에러 타입 케이스는 그냥 그대로 반환됨
func WrapEntError(entErr error) error {
	if entErr != nil {
		if ent.IsNotFound(entErr) {
			return errorz.ErrResourceNotFound
		}
	}
	return entErr
}

type SynchronousCacheWrite bool
