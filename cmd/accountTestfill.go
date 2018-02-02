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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// accountTestfillCmd represents the accountTestfill command
var accountTestfillCmd = &cobra.Command{
	Use:   "testfill",
	Short: "testfill a newly created account",
	Long:  `testfill a newly created account, only works on the testnet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("idess")
		if id == "" && len(args) > 0 {
			id = args[0]
		}
		if id == "" {
			return errors.New("you must specify an address")
		}
		url := "https://horizon-testnet.stellar.org/friendbot?addr=" + id
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		return err
	},
}

func init() {
	accountCmd.AddCommand(accountTestfillCmd)
	accountTestfillCmd.Flags().String("id", "", "accounts id")
}
