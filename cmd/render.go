/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/intob/thinggen/gen"
	"github.com/intob/thinggen/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	renderInputPath string
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render your Things using generated outout",
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
		input, err := os.ReadFile(renderInputPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		things := make(map[string]map[string]gen.Variant)
		err = yaml.Unmarshal(input, &things)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		base, err := util.DecodeImageFromPath(viper.GetString("img"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		i := 0
		if outputPath == "" {
			outputPath = "render"
		}
		os.RemoveAll(outputPath)
		err = os.Mkdir(outputPath, fs.ModePerm)
		if err != nil {
			fmt.Printf("failed to create output directory: %s\r\n", err)
			os.Exit(1)
		}
		for _, thing := range things {
			i++
			name := fmt.Sprintf("%s%v.jpg", viper.GetString("name"), i)
			o := path.Join(outputPath, name)
			err = gen.Render(thing, base, o)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringVar(&renderInputPath, "input", "gen.yml", "Path to generated output")
}
