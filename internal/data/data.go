package data

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/payloads"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Client *ent.Client
}

func NewData(conf *koanf.Koanf, dbConf *payloads.DBConfig) (*Data, error) {
	isDebug := conf.String("app.mode") == gin.DebugMode

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	)

	client, err := ent.Open(dbConf.Dialect, dsn)
	if err != nil {
		return nil, err
	}

	if dbConf.AutoMigrate {
		if err = client.Schema.Create(context.Background()); err != nil {
			return nil, err
		}
	}

	if isDebug {
		client = client.Debug()
	}

	return &Data{Client: client}, nil
}
