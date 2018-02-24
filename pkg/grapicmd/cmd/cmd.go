package cmd

import (
	"os"
	"path/filepath"

	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// NewGrapiCommand creates a new command object.
func NewGrapiCommand(cfg grapicmd.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cfg.AppName(),
		Short: "JSON API framework implemented with gRPC and Gateway",
		Long:  "",
	}

	var cfgFile string
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "./"+cfg.AppName()+".toml", "config file")
	cfg.Init(cfgFile)

	cmd.AddCommand(newInitCommand(cfg.OutWriter()))

	udCmds := make([]*cobra.Command, 0)
	wd, err := os.Getwd()
	if err == nil {
		paths, err := afero.Glob(cfg.Fs(), filepath.Join(wd, "cmd/*/run.go"))
		if err == nil {
			for _, path := range paths {
				udCmds = append(udCmds, newUserDefinedCommand(path))
			}
		}
	}
	if len(udCmds) > 0 {
		cmd.AddCommand(udCmds...)
	}

	return cmd
}