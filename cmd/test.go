/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/intob/nftgen/gen"
	"github.com/intob/nftgen/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Generate a single NFT",
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
		traitMapping := make(map[string]gen.Variant)
		for _, arg := range args {
			kv := strings.Split(arg, "=")
			v := gen.Variant{}
			key := fmt.Sprintf("traits.%s.%s", kv[0], kv[1])
			err := viper.UnmarshalKey(key, &v)
			if err != nil {
				fmt.Println("failed to unmarshal variant", err)
				os.Exit(1)
			}
			traitMapping[kv[0]] = v
		}
		baseImg, err := util.DecodeImageFromPath(viper.GetString("img"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if outputPath == "" {
			outputPath = "output.jpg"
		}
		gen.Render(traitMapping, baseImg, outputPath)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
