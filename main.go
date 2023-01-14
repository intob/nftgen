/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/intob/thinggen/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	cmd.Execute()
}
