package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"gomodules.xyz/kglog"
	"os"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	kglog.ParseFlags()

	rootCmd.AddCommand(NewGetCmd())
	return rootCmd
}

func NewGetCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "hugo get",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	cmd.Flags().String("n", "tamal", "name")

	return cmd
}


func main() {
	kglog.InitLogs()
	defer kglog.FlushLogs()

	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
