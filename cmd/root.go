/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "SimplestDummyFile",
	Short: "Simpliest way of creating a dummy file",
	Long: `A cobra go program for creating very primitive dummy files.
The files are extremely compressable, that is because all of the bytes of the file are esentially 0.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		name, _ := cmd.Flags().GetString("name")
		dummy(name, size)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "dummy", "The name of the file to write to (If no file exists, it will be created)")
	rootCmd.Flags().IntP("size", "s", 1000, "Size (In Kilobytes)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dummy(name string, filesize int) {
	timeStart := time.Now().UnixMilli()
	var arr = make([]byte, filesize*1000, filesize*1000)

	os.WriteFile(name, arr, os.ModePerm)
	timeEnd := time.Now().UnixMilli()
	fmt.Printf("Process completed in %dms", timeEnd-timeStart)

}
