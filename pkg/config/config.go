package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
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
