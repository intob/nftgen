/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"

	"github.com/intob/nftgen/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
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
		count := viper.GetInt("count")
		if count < 1 {
			fmt.Println("no count defined")
			os.Exit(1)
		}
		possibleCount, err := gen.CountPossibleMappings(traits)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if count > possibleCount {
			fmt.Printf("configured count of %v is greater than number of possible variants: %v\r\n", count, possibleCount)
			os.Exit(1)
		}
		m := make(map[string]map[string]gen.Variant)
		hash := fnv.New128()
		for len(m) < count {
			r, err := gen.RandomTraitMapping(traits)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			b, err := json.Marshal(r)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			hash.Write(b)
			h := hash.Sum(nil)
			hash.Reset()
			hStr := hex.EncodeToString(h)
			m[hStr] = r
			fmt.Print("#")
		}
		fmt.Print("\r\n")
		yml, err := yaml.Marshal(m)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if outputPath == "" {
			outputPath = "gen.yml"
		}
		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		f.Write(yml)
		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
