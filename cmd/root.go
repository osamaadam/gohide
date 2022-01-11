package cmd

import (
	"path/filepath"
	"strings"

	"github.com/osamaadam/gohide/hide"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var opts struct {
	Glob   string
	Unhide bool
}

var rootCmd = &cobra.Command{
	Use:   "gohide files-to-hide...",
	Short: "small utility to hide, and unhide files",
	RunE:  rootRun,
}

func rootRun(cmd *cobra.Command, args []string) error {
	files := args
	opts.Glob = strings.TrimSpace(opts.Glob)

	if len(opts.Glob) > 0 {
		globFiles, err := filepath.Glob(opts.Glob)
		if err != nil {
			return errors.WithStack(err)
		}
		files = append(files, globFiles...)
	}

	if opts.Unhide {
		if err := hide.Unhide(files...); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := hide.Hide(files...); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&opts.Glob, "glob", "g", "", "specify glob to match")
	rootCmd.Flags().BoolVarP(&opts.Unhide, "unhide", "H", false, "unhide files")
}
