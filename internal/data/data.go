package data

import (
	"fmt"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/knadh/koanf"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	client *ent.Client
}

func NewData(conf *koanf.Koanf) (*Data, error) {
	dialect := conf.String("db.dialect")
	host := conf.String("db.host")
	port := conf.Int("db.port")
	username := conf.String("db.username")
	password := conf.String("db.password")
	dbname := conf.String("db.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", username, password, host, port, dbname)

	client, err := ent.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	return &Data{client: client}, nil
}
