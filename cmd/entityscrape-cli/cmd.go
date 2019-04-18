package main

import (
	cli "github.com/ndabAP/entityscrape/internal/entityscrape-cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "entityscrape",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		entity := args[0]
		url := args[1]

		err := cli.Make(entity, url, aliases)
		if err != nil {
			return err
		}

		return nil
	},
}

var aliases []string

func init() {
	rootCmd.Flags().StringSliceVarP(&aliases, "aliases", "a", []string{}, "")
}
