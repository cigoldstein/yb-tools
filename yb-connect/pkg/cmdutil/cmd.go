package cmdutil

import (
	"context"
	"fmt"

	"github.com/blang/vfs"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/yugabyte/yb-tools/pkg/flag"
	"github.com/yugabyte/yb-tools/pkg/log"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client"
)

type CommandOptions interface {
	AddFlags(cmd *cobra.Command)
	Validate() error
}

type YBConnectContext struct {
	context.Context
	Log            logr.Logger
	Cmd            *cobra.Command
	GlobalOptions  *YCGlobalOptions
	CommandOptions CommandOptions
	Fs             vfs.Filesystem

	Client *client.YugawareClient
}

func NewCommandContext() *YBConnectContext {
	return &YBConnectContext{
		Context: context.Background(),
	}
}

func (ctx *YBConnectContext) WithGlobalOptions(options *YCGlobalOptions) *YBConnectContext {
	ctx.GlobalOptions = options
	return ctx
}

func (ctx *YBConnectContext) WithCmd(cmd *cobra.Command) *YBConnectContext {
	ctx.Cmd = cmd
	return ctx
}

func (ctx *YBConnectContext) WithOptions(options CommandOptions) *YBConnectContext {
	ctx.CommandOptions = options
	return ctx
}

func (ctx *YBConnectContext) Setup() error {
	if ctx.Cmd == nil ||
		ctx.GlobalOptions == nil {
		panic("command context is not set")
	}

	setupError := func(err error) error {
		return fmt.Errorf("failed to setup %s command: %w", ctx.Cmd.Name(), err)
	}

	var err error

	ctx.Log, err = log.GetLogger(ctx.Cmd.Name(), ctx.GlobalOptions.Debug)
	if err != nil {
		return setupError(err)
	}

	err = ctx.complete()
	if err != nil {
		return setupError(err)
	}

	err = flag.ValidateRequiredFlags(ctx.Cmd.Flags())
	if err != nil {
		return err
	}

	err = ctx.GlobalOptions.Validate()
	if err != nil {
		return err
	}

	if ctx.CommandOptions != nil {
		err = ctx.CommandOptions.Validate()
		if err != nil {
			return err
		}
	}

	ctx.Cmd.SilenceUsage = true

	err = ctx.Connect()
	if err != nil {
		return setupError(err)
	}

	return nil
}

func (ctx *YBConnectContext) Connect() error {
	c, err := ConnectToYugaware(ctx)
	ctx.Client = c
	return err
}

func (ctx *YBConnectContext) complete() error {
	flag.BindFlags(ctx.Cmd.Flags())

	err := flag.MergeConfigFile(ctx.Log, ctx.GlobalOptions)
	if err != nil {
		return err
	}

	if ctx.CommandOptions != nil {
		err := flag.MergeConfigFile(ctx.Log, ctx.CommandOptions)
		if err != nil {
			return err
		}
	}

	return nil
}

var _ CommandOptions = &YCGlobalOptions{}

type YCGlobalOptions struct {
	Debug                bool   `mapstructure:"debug"`
	Output               string `mapstructure:"output"`
	Hostname             string `mapstructure:"hostname"`
	DialTimeout          int    `mapstructure:"dialtimeout"`
	SkipHostVerification bool   `mapstructure:"skiphostverification"`
	CACert               string `mapstructure:"cacert"`
	ClientCert           string `mapstructure:"client_cert"`
	ClientKey            string `mapstructure:"client_key"`
}

func (o *YCGlobalOptions) AddFlags(cmd *cobra.Command) {
	// Global configuration flags
	flags := cmd.PersistentFlags()
	flags.BoolVar(&o.Debug, "debug", false, "debug mode")
	flags.StringVarP(&o.Output, "output", "o", "table", "Output options as one of: [table, json, yaml]")
	flags.StringVar(&o.Hostname, "hostname", "localhost:8080", "hostname of yugaware")
	flags.IntVar(&o.DialTimeout, "dialtimeout", 10, "number of seconds for dial timeouts")
	flags.BoolVar(&o.SkipHostVerification, "skiphostverification", false, "skip tls host verification")
	flags.StringVarP(&o.CACert, "cacert", "c", "", "the path to the CA certificate")
	flags.StringVar(&o.ClientCert, "client-cert", "", "the path to the client certificate")
	flags.StringVar(&o.ClientKey, "client-key", "", "the path to the client key file")
}

func (o *YCGlobalOptions) Validate() error {
	return nil
}

func ConnectToYugaware(ctx *YBConnectContext) (*client.YugawareClient, error) {
	return client.New(ctx, ctx.Log, ctx.GlobalOptions.Hostname).
		TLSOptions(&client.TLSOptions{
			SkipHostVerification: ctx.GlobalOptions.SkipHostVerification,
			CaCertPath:           ctx.GlobalOptions.CACert,
			CertPath:             ctx.GlobalOptions.ClientCert,
			KeyPath:              ctx.GlobalOptions.ClientKey,
		}).TimeoutSeconds(ctx.GlobalOptions.DialTimeout).Connect()
}
