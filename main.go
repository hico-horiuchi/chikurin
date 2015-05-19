package main

import (
	"fmt"

	"./chikurin"
	"github.com/spf13/cobra"
)

var version string

func main() {
	rootCmd := &cobra.Command{
		Use:   "chikurin",
		Short: "Sensu status page by golang",
		Long:  "Sensu status page by golang\nhttps://github.com/hico-horiuchi/chikurin",
		Run: func(cmd *cobra.Command, args []string) {
			chikurin.Serve()
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print and check version of chikurin",
		Long:  "Print and check version of chikurin",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(chikurin.Version(version))
		},
	})

	rootCmd.Execute()
}
