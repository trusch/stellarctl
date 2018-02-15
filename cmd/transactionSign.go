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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/trusch/stellarctl/transaction"
	yaml "gopkg.in/yaml.v2"
)

// transactionSignCmd represents the transactionSign command
var transactionSignCmd = &cobra.Command{
	Use:   "sign",
	Short: "add your signature to a transaction file",
	Long:  `add your signature to a transaction file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		if output == "" {
			output = input
		}
		seed, _ := cmd.Flags().GetString("seed")
		bs, err := ioutil.ReadFile(input)
		if err != nil {
			return err
		}
		tx, err := transaction.FromYAML(bs)
		if err != nil {
			return errors.Wrap(err, "can not parse transaction file")
		}
		if err = tx.Sign(seed); err != nil {
			return err
		}
		bs, err = yaml.Marshal(tx)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(output, bs, 0644)
	},
}

func init() {
	transactionCmd.AddCommand(transactionSignCmd)
	transactionSignCmd.Flags().String("input", "", "transaction file")
	transactionSignCmd.Flags().String("output", "", "signed output file")
	transactionSignCmd.Flags().String("seed", "", "account seed")
}
