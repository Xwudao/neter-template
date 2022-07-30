package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
)

func NewConfig() (*koanf.Koanf, error) {
	var k = koanf.New(".")

	//default
	setDefault(k)

	f := file.Provider("config.yml")
	if err := k.Load(f, yaml.Parser()); err != nil {
		return nil, err
	}

	_ = f.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("config file changed")
		k = koanf.New(".")
		_ = k.Load(f, yaml.Parser())
		k.Print()
	})

	return k, nil
}

func setDefault(k *koanf.Koanf) {
	/*
		app:
		  port: 8080
		  mode: debug

		log:
		  level: debug
		  format: json
		  linkName: current.log
		  path: ./logs

			cors:
		  allowOrigin:
		    - http://localhost:*
		    - http://127.0.0.1:*
		  allowCredentials: true
		  maxAge: 24h
	*/
	_ = k.Load(confmap.Provider(map[string]interface{}{
		"cors.allowOrigin":      []string{"http://localhost:*", "http://127.0.0.1:*"},
		"cors.allowCredentials": true,
		"cors.maxAge":           "24h",
		"app.port":              8080,
		"app.mode":              "debug",
		"log.level":             "debug",
		"log.format":            "json",
		"log.linkName":          "current.log",
		"log.path":              "./logs",
	}, "."), nil)
}
