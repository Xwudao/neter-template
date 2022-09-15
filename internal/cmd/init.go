/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init something that the system need",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config something that the system need",
	Run: func(cmd *cobra.Command, args []string) {
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
		//_, _ = config.NewConfig(koanf)
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
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
