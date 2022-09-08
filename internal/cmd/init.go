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

		var ff *os.File
		_, err := os.Stat(f)
		if err != nil {
			ff, err = os.Create(f)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("config file already exist")
		}
		defer ff.Close()
		if err != nil {
			log.Fatal(err)
		}
		parser := yaml.Parser()
		koanf, err := config.NewKoanf()
		if err != nil {
			log.Fatal(err)
		}
		data, err := koanf.Marshal(parser)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ff.Write(data)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("config file created")
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
