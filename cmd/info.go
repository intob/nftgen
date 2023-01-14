/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/intob/nftgen/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get stats about your project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath != "" {
			viper.SetConfigFile(configPath)
		}
		viper.ReadInConfig()
		traits := make(map[string]gen.Trait)
		err := viper.UnmarshalKey("traits", &traits)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		n, err := gen.CountPossibleMappings(traits)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("There are %v possible variants\r\n", n)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
