package main

import (
	"fmt"

	"./chikurin"
	"github.com/VividCortex/godaemon"
	"github.com/spf13/cobra"
)

var version string

func main() {
	rootCmd := &cobra.Command{
		Use:   "chikurin",
		Short: "Sensu status page by golang",
		Long:  "Sensu status page by golang\nhttps://github.com/hico-horiuchi/chikurin",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			chikurin.LoadConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			chikurin.Serve()
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start chikurin daemon",
		Long:  "Start chikurin daemon",
		Run: func(cmd *cobra.Command, args []string) {
			if chikurin.Status() > -1 {
				fmt.Println("chikurin is already running")
			} else {
				fmt.Println("Starting chikurin")
				godaemon.MakeDaemon(&godaemon.DaemonAttr{})
				chikurin.Start()
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stop chikurin daemon",
		Long:  "Stop chikurin daemon",
		Run: func(cmd *cobra.Command, args []string) {
			if chikurin.Status() > -1 {
				fmt.Println("Stopping chikurin")
				chikurin.Stop()
			} else {
				fmt.Println("chikurin is stopped")
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Show status of chikurin daemon",
		Long:  "Show status of chikurin daemon",
		Run: func(cmd *cobra.Command, args []string) {
			pid := chikurin.Status()
			if pid > -1 {
				fmt.Printf("chikurin (pid %d) is runnning\n", pid)
			} else {
				fmt.Println("chikurin is stopped")
			}
		},
	})

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
