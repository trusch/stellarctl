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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stellar/go/keypair"
)

// accountAddressCmd represents the accountAddress command
var accountAddressCmd = &cobra.Command{
	Use:   "address",
	Short: "get account address for a given seed",
	Long:  `get account address for a given seed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seed, _ := cmd.Flags().GetString("seed")
		if seed == "" && len(args) > 0 {
			seed = args[0]
		}
		if seed == "" {
			return errors.New("you must specify a seed")
		}
		kp, err := keypair.Parse(seed)
		if err != nil {
			return err
		}
		fmt.Println(kp.Address())
		return nil
	},
}

func init() {
	accountCmd.AddCommand(accountAddressCmd)
	accountAddressCmd.Flags().String("seed", "", "account seed")
}
