package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gomodules.xyz/kglog"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
	"os"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
			klog.InitFlags(klogFlags)

			// Sync the glog and klog flags.
			cmd.Flags().VisitAll(func(f1 *pflag.Flag) {
				f2 := klogFlags.Lookup(f1.Name)
				if f2 != nil {
					value := f1.Value.String()
					// Ignore error. klog's -log_backtrace_at flag throws error when set to empty string.
					// Unfortunately, there is no way to tell if a flag was set to empty string or left unset on command line.
					_ = f2.Value.Set(value)
				}
			})
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	//pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	// kglog.ParseFlags()
	// pflag.Parse()

	rootCmd.AddCommand(NewGetCmd())
	return rootCmd
}

func NewGetCmd() *cobra.Command {
	var name string
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "hugo get",
		Run: func(cmd *cobra.Command, args []string) {
			klog.V(3).Info("name = ", name)
		},
	}
	cmd.Flags().StringVarP(&name, "name","n", "tamal", "name")

	return cmd
}


func main() {
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	// ParseFlags()

	kglog.InitLogs()
	defer kglog.FlushLogs()

	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ParseFlags() {
	flag.Parse()

	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)

	// Sync the glog and klog flags.
	flag.CommandLine.VisitAll(func(f1 *flag.Flag) {
		f2 := klogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			// Ignore error. klog's -log_backtrace_at flag throws error when set to empty string.
			// Unfortunately, there is no way to tell if a flag was set to empty string or left unset on command line.
			_ = f2.Value.Set(value)
		}
	})
}
