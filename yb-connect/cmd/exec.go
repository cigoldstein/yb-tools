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

	"github.com/spf13/cobra"
	"github.com/yugabyte/yb-tools/pkg/flag"
	"github.com/yugabyte/yb-tools/yb-connect/client"
	"github.com/yugabyte/yb-tools/yb-connect/pkg/cmdutil"
	"github.com/yugabyte/yb-tools/yb-connect/validation"
)

func ExecCmd(ctx *cmdutil.YBConnectContext) *cobra.Command {
	options := &ExecOptions{}
	cmd := &cobra.Command{
		Use:   "exec --universe <UNIVERSE> <COMMAND>",
		Short: "Execute a command against multiple Yugabyte hosts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := ctx.WithCmd(cmd).WithOptions(options).Setup()
			if err != nil {
				return err
			}

			// Positional argument
			options.Command = args[0]

			//TODO: replace with general file permissions check instead of root user check
			err = validation.VerifyRootUser()
			if err != nil {
				return err
			}

			return exec(ctx, options)
		},
	}

	options.AddFlags(cmd)

	return cmd
}

func exec(ctx *cmdutil.YBConnectContext, options *ExecOptions) error {
	universe, err := ctx.Client.GetUniverseByName(options.UniverseName)
	if err != nil {
		return err
	}

	if universe == nil {
		return fmt.Errorf("universe does not exist: %s", options.UniverseName)
	}

	ctx.Log.V(1).Info("got universe", "universe", universe)

	return client.SshToNodes(ctx.Log, universe, options.Command)
}

type ExecOptions struct {
	Command string

	//TODO: add env variable(s)?
	UniverseName string
	Tservers     bool
	Masters      bool
}

// My headset died -- brb

var _ cmdutil.CommandOptions = &ExecOptions{}

func (o *ExecOptions) AddFlags(cmd *cobra.Command) {
	flags := cmd.Flags()

	flags.StringVar(&o.UniverseName, "universe", "", "Specify universe name")
	flags.BoolVar(&o.Tservers, "tservers", false, "Connect to only tserver nodes")
	flags.BoolVar(&o.Masters, "masters", false, "Connect to only master nodes")

	flag.MarkFlagRequired("universe", flags)
}

func (o *ExecOptions) Validate() error {
	return nil
}
