package main

import (
	"github.com/spf13/cobra"
)

const (
	unicodeSmallA   = 97
	unicodeSmallZ   = 122
	unicodeCapitalA = 65
	unicodeCapitalZ = 90
)

// RootCmd does it
var RootCmd = &cobra.Command{
	Use:  "entityscrape",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		entity := args[0]
		url := args[1]

		Associate(entity, url, aliases)

		return nil
	},
}
var aliases []string

func init() {
	RootCmd.Flags().StringSliceVarP(&aliases, "aliases", "a", []string{}, "")
}
