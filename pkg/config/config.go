package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
)

type AppConfig struct {
	App struct {
		Port int    `koanf:"port,omitempty"`
		Mode string `koanf:"mode,omitempty"`
	} `koanf:"app,omitempty"`
	Log struct {
		Level    string `koanf:"level,omitempty"`
		Format   string `koanf:"format,omitempty"`
		LinkName string `koanf:"linkName,omitempty"`
		Path     string `koanf:"path,omitempty"`
	} `koanf:"log,omitempty"`
	Db struct {
		Dialect     string `koanf:"dialect,omitempty"`
		Host        string `koanf:"host,omitempty"`
		Port        int    `koanf:"port,omitempty"`
		Username    string `koanf:"username,omitempty"`
		Password    string `koanf:"password,omitempty"`
		Database    string `koanf:"database,omitempty"`
		TablePrefix string `koanf:"tablePrefix,omitempty"`
		DevDsn      string `koanf:"devDsn,omitempty"`
	} `koanf:"db,omitempty"`
	Jwt struct {
		Secret string `koanf:"secret,omitempty"`
	} `koanf:"jwt,omitempty"`
	Cors struct {
		AllowOrigin      []string `koanf:"allowOrigin,omitempty"`
		AllowCredentials bool     `koanf:"allowCredentials,omitempty"`
		MaxAge           string   `koanf:"maxAge,omitempty"`
	} `koanf:"cors,omitempty"`
	Redis struct {
		Addr     string `koanf:"addr,omitempty"`
		Password string `koanf:"password,omitempty"`
		Db       int    `koanf:"db,omitempty"`
	} `koanf:"redis,omitempty"`
	Mail struct {
		From           string        `koanf:"from,omitempty"`
		Host           string        `koanf:"host,omitempty"`
		Port           int           `koanf:"port,omitempty"`
		Username       string        `koanf:"username,omitempty"`
		Password       string        `koanf:"password,omitempty"`
		KeepAlive      bool          `koanf:"keepAlive,omitempty"`
		ConnectTimeout time.Duration `koanf:"connectTimeout,omitempty"`
		SendTimeout    time.Duration `koanf:"sendTimeout,omitempty"`
	} `koanf:"mail,omitempty"`
}

func setDefault(k *koanf.Koanf) {
	_ = k.Load(confmap.Provider(map[string]any{
		"cors.allowOrigin":      []string{"http://localhost:*", "http://127.0.0.1:*"},
		"cors.allowCredentials": true,
		"cors.maxAge":           "24h",

		"app.port":  8080,
		"app.mode":  "debug",
		"app.pprof": false,

		"log.level":    "debug",
		"log.format":   "json",
		"log.linkName": "current.log",
		"log.path":     "./logs",

		"db.dialect":     "mysql",
		"db.host":        "127.0.0.1",
		"db.port":        3306,
		"db.username":    "root",
		"db.password":    "root",
		"db.database":    "ent-demo",
		"db.tablePrefix": "",
		"db.devDsn":      "mysql://root:root@:3306/dev-ent",

		"jwt.secret": "secret",

		"redis.addr":     "127.0.0.1:6379",
		"redis.password": "",
		"redis.db":       0,

		"mail.from":           "xx",
		"mail.host":           "xx",
		"mail.port":           466,
		"mail.username":       "xx",
		"mail.password":       "xx",
		"mail.keepAlive":      true,
		"mail.connectTimeout": "10s",
		"mail.sendTimeout":    "10s",
	}, "."), nil)
}

func NewConfig(k *koanf.Koanf) (*AppConfig, error) {
	var c = &AppConfig{}
	if err := k.Unmarshal("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func NewKoanf() (*koanf.Koanf, error) {
	var k = koanf.New(".")

	//default
	setDefault(k)

	var err error

	var files = []string{"config.yml", "config.dev.yml"}
	for _, fn := range files {
		f := file.Provider(fn)
		if err = k.Load(f, yaml.Parser()); err != nil {
			log.Println("load file " + fn + " failed")
			continue
		}

		_ = f.Watch(func(event any, err error) {
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("config file changed")
			k = koanf.New(".")
			_ = k.Load(f, yaml.Parser())
			k.Print()
		})
	}

	return k, nil
}

func NewTestConfig() (*koanf.Koanf, error) {

	dir, _ := os.Getwd()
	var k = koanf.New(".")
	paths := []string{
		filepath.Join(dir, "config.yml"),
		filepath.Join(dir, "config.json"),
		filepath.Join(dir, "testdata/config.yml"),
		filepath.Join(dir, "testdata/config.json"),
	}
	var success = 0
	for _, path := range paths {
		var err error
		switch filepath.Ext(path) {
		case ".yml", ".yaml":
			err = k.Load(file.Provider(path), yaml.Parser())
		case ".json":
			err = k.Load(file.Provider(path), json.Parser())
		}
		if err == nil {
			success++
		}
	}
	if success == 0 {
		log.Fatalf("no config file found in %v", paths)
	}

	//if err := k.Load(file.Provider(p), json.Parser()); err != nil {
	//	log.Fatalf("error loading config: %v", err)
	//}

	return k, nil
}
