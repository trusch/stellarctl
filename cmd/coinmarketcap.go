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
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// coinmarketcapCmd represents the coinmarketcap command
var coinmarketcapCmd = &cobra.Command{
	Use:   "coinmarketcap",
	Short: "get coinmarketcap.com data",
	Long:  `get coinmarketcap.com data`,
	RunE: func(cmd *cobra.Command, args []string) error {
		coin, _ := cmd.Flags().GetString("coin")
		convert, _ := cmd.Flags().GetString("convert")
		limit, _ := cmd.Flags().GetString("limit")
		url := "https://api.coinmarketcap.com/v1/ticker/"
		if coin != "all" {
			url += coin + "/"
			if convert != "" {
				url += "?convert=" + convert
			}
		} else {
			if limit != "" {
				url += "?limit=" + limit
			}
			if convert != "" {
				url += "&convert=" + convert
			}
		}
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			bs, _ := ioutil.ReadAll(resp.Body)
			err = errors.New(string(bs))
			return err
		}
		doc := make([]map[string]interface{}, 0)
		if err = json.NewDecoder(resp.Body).Decode(&doc); err != nil {
			return err
		}
		printWithFormat(doc)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(coinmarketcapCmd)
	coinmarketcapCmd.Flags().String("coin", "stellar", "which coin")
	coinmarketcapCmd.Flags().String("convert", "EUR", "include converted value")
	coinmarketcapCmd.Flags().String("limit", "", "limit output when query for 'all' coins")
}
