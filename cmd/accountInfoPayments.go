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
	"net/http"

	"github.com/spf13/cobra"
)

// accountInfoPaymentsCmd represents the accountInfoPayments command
var accountInfoPaymentsCmd = &cobra.Command{
	Use:   "payments",
	Short: "list account payments",
	Long:  `list account payments`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("id")
		id = MustResolveAddress(id)
		cursor, _ := cmd.Flags().GetInt("cursor")
		limit, _ := cmd.Flags().GetInt("limit")
		order, _ := cmd.Flags().GetString("order")
		url := getClientURL()
		url += "/accounts/" + id + "/payments?"
		if cursor != -1 {
			url += fmt.Sprintf("cursor=%v&", cursor)
		}
		if limit != -1 {
			url += fmt.Sprintf("limit=%v&", limit)
		}
		if order != "" {
			url += fmt.Sprintf("order=%v&", order)
		}
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return errors.New("http status: " + resp.Status)
		}
		defer resp.Body.Close()
		return print(resp.Body)
	},
}

func init() {
	accountInfoCmd.AddCommand(accountInfoPaymentsCmd)
	accountInfoPaymentsCmd.Flags().Int("cursor", -1, "list cursor")
	accountInfoPaymentsCmd.Flags().Int("limit", -1, "list limit")
	accountInfoPaymentsCmd.Flags().String("order", "desc", "list order")
}
