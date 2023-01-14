/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
	rootCmd    = &cobra.Command{
		Use:   "nftgen",
		Short: "Generate variants of a base image for NFTs",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		//Run: func(cmd *cobra.Command, args []string) {},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Point to a specific config file")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output", "", "Path for generated output")
}
