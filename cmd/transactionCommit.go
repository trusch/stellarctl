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
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/trusch/stellarctl/transaction"
)

// transactionCommitCmd represents the transactionCommit command
var transactionCommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit a transaction to the network",
	Long:  `commit a transaction to the network`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := getClient()
		input, _ := cmd.Flags().GetString("input")
		bs, err := ioutil.ReadFile(input)
		if err != nil {
			return err
		}
		tx, err := transaction.FromYAML(bs)
		if err != nil {
			return err
		}
		xdr, err := tx.ToXDR()
		if err != nil {
			return err
		}
		_, err = cli.SubmitTransaction(xdr)
		return err
	},
}

func init() {
	transactionCmd.AddCommand(transactionCommitCmd)
	transactionCommitCmd.Flags().String("input", "", "transaction file")
}
