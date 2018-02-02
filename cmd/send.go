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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send assets",
	Long:  `send assets to another account`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seed, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		amount, _ := cmd.Flags().GetString("amount")
		assetCode, _ := cmd.Flags().GetString("asset-code")
		assetIssuer, _ := cmd.Flags().GetString("asset-issuer")
		memo, _ := cmd.Flags().GetString("memo")
		cli := getClient()
		txArgs := []build.TransactionMutator{
			build.SourceAccount{AddressOrSeed: seed},
			build.AutoSequence{SequenceProvider: cli},
		}
		if viper.GetBool("testnet") {
			txArgs = append(txArgs, build.TestNetwork)
		}

		paymentArgs := []interface{}{
			build.Destination{AddressOrSeed: to},
		}
		if assetCode == "" {
			paymentArgs = append(paymentArgs, build.NativeAmount{Amount: amount})
		} else {
			paymentArgs = append(paymentArgs,
				build.CreditAmount{
					Code:   assetCode,
					Issuer: assetIssuer,
					Amount: amount,
				},
			)
		}
		txArgs = append(txArgs, build.Payment(paymentArgs...))
		if memo != "" {
			txArgs = append(txArgs, build.MemoText{Value: memo})
		}
		tx, err := build.Transaction(txArgs...)
		if err != nil {
			return err
		}
		txe, err := tx.Sign(seed)
		if err != nil {
			return err
		}
		txeB64, err := txe.Base64()
		if err != nil {
			return err
		}
		_, err = cli.SubmitTransaction(txeB64)
		if err != nil {
			if e, ok := err.(*horizon.Error); ok {
				fmt.Println(string(e.Problem.Extras["result_codes"]))
			}
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().String("from", "", "source account seed ")
	sendCmd.Flags().String("to", "", "destination account address")
	sendCmd.Flags().String("amount", "0", "amount")
	sendCmd.Flags().String("asset-code", "", "asset code")
	sendCmd.Flags().String("asset-issuer", "", "asset issuer")
	sendCmd.Flags().String("memo", "", "transaction memo")

}
