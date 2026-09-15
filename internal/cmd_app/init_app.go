package cmd_app

import (
	"log"
	"os"

	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/knadh/koanf/parsers/yaml"
)

type InitApp struct {
}

func NewInitApp() *InitApp {
	return &InitApp{}
}

func (a *InitApp) Run() {

}
func (a *InitApp) Config() {
	f := "config.yml"

	ff, err := os.OpenFile(f, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	defer ff.Close()
	parser := yaml.Parser()
	koanf, err := config.NewKoanf()
	if err != nil {
		log.Fatalf("new koanf err: %s", err)
	}
	// _, _ = config.NewConfig(koanf)
	data, err := koanf.Marshal(parser)
	if err != nil {
		log.Fatalf("marshal config err: %s", err)
	}
	log.Println(string(data))
	_, err = ff.Write(data)
	if err != nil {
		log.Fatalf("write config err: %s", err)
	} else {
		log.Println("write config!")
	}
}
