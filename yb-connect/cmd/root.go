/*
Copyright © 2021 Yugabyte Support

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yugabyte/yb-tools/yb-connect/pkg/cmdutil"
)

var (
	cfgFile string

	Version = "DEV"
)

var rootCmd = RootInit()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ybconnect" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ybconnect")
	}

	viper.SetEnvPrefix("YC")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func RootInit() *cobra.Command {
	globalOptions := &cmdutil.YCGlobalOptions{}

	cmd := &cobra.Command{
		Use:     "yb-connect",
		Short:   "Connect to one or more Yugabyte resources",
		Version: Version,
	}

	cmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ybconnect.yaml)")
	globalOptions.AddFlags(cmd)

	ctx := cmdutil.NewCommandContext().
		WithGlobalOptions(globalOptions)

	// Top level commands
	cmd.AddCommand(ExecCmd(ctx))

	return cmd
}
