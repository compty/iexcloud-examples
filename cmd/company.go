// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/goinvest/iexcloud-examples/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(companyCmd)
}

var companyCmd = &cobra.Command{
	Use:   "company [stock]",
	Short: "Retrieve the company data for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		company, err := client.Company(context.Background(), stock)
		fmt.Printf("%q\n", company)
	},
}
