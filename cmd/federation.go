// Copyright © 2018 Tino Rusch
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/federation"
)

// federationCmd represents the federation command
var federationCmd = &cobra.Command{
	Use:   "federation",
	Short: "get account id for a federation address",
	Long:  `get account id for a federation address`,
	RunE: func(cmd *cobra.Command, args []string) error {
		address, _ := cmd.Flags().GetString("address")
		if address == "" && len(args) > 0 {
			address = args[0]
		}
		addr, err := ResolveAddress(address)
		if err != nil {
			return err
		}
		fmt.Println(addr)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(federationCmd)
	federationCmd.Flags().String("address", "", "federation address")
}

func ResolveAddress(address string) (string, error) {
	if len(address) == 56 && address[0] == 'G' {
		return address, nil
	}
	resp, err := federation.DefaultPublicNetClient.LookupByAddress(address)
	if err != nil {
		return "", err
	}
	return resp.AccountID, nil
}

func MustResolveAddress(address string) string {
	addr, err := ResolveAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	return addr
}
