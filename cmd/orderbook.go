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
	"strings"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizon"
	"github.com/trusch/stellarctl/transaction"
)

// orderbookCmd represents the orderbook command
var orderbookCmd = &cobra.Command{
	Use:   "orderbook",
	Short: "get orderbook of an SDEX trading pair",
	Long:  `get orderbook of an SDEX trading pair`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sellingCode, _ := cmd.Flags().GetString("selling-code")
		buyingCode, _ := cmd.Flags().GetString("buying-code")
		sellingIssuer, _ := cmd.Flags().GetString("selling-issuer")
		buyingIssuer, _ := cmd.Flags().GetString("buying-issuer")
		limit, _ := cmd.Flags().GetInt("limit")
		selling := horizon.Asset{
			Code:   sellingCode,
			Issuer: sellingIssuer,
			Type:   "credit_alphanum4",
		}
		buying := horizon.Asset{
			Code:   buyingCode,
			Issuer: buyingIssuer,
			Type:   "credit_alphanum4",
		}
		if selling.Code == "" || strings.ToUpper(selling.Code) == transaction.NativeAssetCode {
			selling.Type = "native"
		} else if len(selling.Code) > 4 {
			selling.Type = "credit_alphanum12"
		}
		if buying.Code == "" || strings.ToUpper(buying.Code) == transaction.NativeAssetCode {
			buying.Type = "native"
		} else if len(buying.Code) > 4 {
			buying.Type = "credit_alphanum12"
		}
		orderBookArgs := []interface{}{}
		if limit > 0 {
			orderBookArgs = append(orderBookArgs, horizon.Limit(limit))
		}

		cli := getClient()
		book, err := cli.LoadOrderBook(selling, buying)
		if err != nil {
			return err
		}
		print(book)
		// url := fmt.Sprintf("%v/order_book?selling_asset_type=%v&selling_asset_code=%v&selling_asset_issuer=%v&buying_asset_type=%v&buying_asset_code=%v&buying_asset_issuer=%v",
		// 	getClientURL(),
		// 	selling.Type, selling.Code, sellingIssuer,
		// 	buying.Type, buying.Code, buyingIssuer,
		// )
		// if limit > 0 {
		// 	url += fmt.Sprintf("&limit=%v", limit)
		// }
		// fmt.Println(url)
		// resp, err := http.Get(url)
		// if err != nil {
		// 	return err
		// }
		// defer resp.Body.Close()
		// io.Copy(os.Stdout, resp.Body)
		// doc := make(map[string]interface{})
		// if err = json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		// 	return err
		// }
		// print(doc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(orderbookCmd)
	orderbookCmd.Flags().String("selling-code", "", "selling asset code")
	orderbookCmd.Flags().String("selling-issuer", "", "selling asset issuer")
	orderbookCmd.Flags().String("buying-code", "", "buying asset code")
	orderbookCmd.Flags().String("buying-issuer", "", "buying asset issuer")
	orderbookCmd.Flags().Int("limit", 0, "list limit")
}
