/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/spf13/cobra"
)

type initCmd struct {
	cmd *cobra.Command
}

func (c *initCmd) run(cmd *cobra.Command, args []string) {

}

func (c *initCmd) getCmd() *cobra.Command {
	return c.cmd
}

func newInitCmd(confCmd *configCmd) *initCmd {
	m := &cobra.Command{
		Use:   "init",
		Short: "init something that the system need",
	}
	rtn := &initCmd{
		cmd: m,
	}
	m.AddCommand(confCmd.getCmd())

	return rtn
}

type configCmd struct {
	cmd *cobra.Command
}

func (c *configCmd) getCmd() *cobra.Command {
	return c.cmd
}
func (c *configCmd) run(cmd *cobra.Command, args []string) {
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
func newConfigCmd() *configCmd {
	m := &cobra.Command{
		Use:   "config",
		Short: "config something that the system need",
	}
	rtn := &configCmd{
		cmd: m,
	}
	m.Run = rtn.run

	return rtn
}
