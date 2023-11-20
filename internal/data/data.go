package data

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Client *ent.Client
}

func NewData(conf *koanf.Koanf) (*Data, error) {
	dialect := conf.String("db.dialect")
	host := conf.String("db.host")
	port := conf.Int("db.port")
	username := conf.String("db.username")
	password := conf.String("db.password")
	dbname := conf.String("db.database")
	autoMigrate := conf.Bool("db.autoMigrate")
	isDebug := conf.String("app.mode") == gin.DebugMode

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", username, password, host, port, dbname)

	client, err := ent.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	if autoMigrate {
		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, err
		}
	}

	if isDebug {
		client = client.Debug()
	}

	return &Data{Client: client}, nil
}
