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

// setOptionsCmd represents the setOptions command
var setOptionsCmd = &cobra.Command{
	Use:   "set-options",
	Short: "set account options",
	Long:  `set account options`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seed, _ := cmd.Flags().GetString("seed")
		cli := getClient()
		txArgs := []build.TransactionMutator{
			build.SourceAccount{AddressOrSeed: seed},
			build.AutoSequence{SequenceProvider: cli},
		}
		if viper.GetBool("testnet") {
			txArgs = append(txArgs, build.TestNetwork)
		}
		if inf, _ := cmd.Flags().GetString("inflation-destination"); inf != "" {
			txArgs = append(txArgs, build.InflationDest(inf))
		}
		if domain, _ := cmd.Flags().GetString("home-domain"); domain != "" {
			txArgs = append(txArgs, build.HomeDomain(domain))
		}
		if b, _ := cmd.Flags().GetBool("set-auth-required"); b {
			txArgs = append(txArgs, build.SetAuthRequired())
		}
		if b, _ := cmd.Flags().GetBool("set-auth-revocable"); b {
			txArgs = append(txArgs, build.SetAuthRevocable())
		}
		if b, _ := cmd.Flags().GetBool("set-auth-immutable"); b {
			txArgs = append(txArgs, build.SetAuthImmutable())
		}
		if b, _ := cmd.Flags().GetBool("clear-auth-required"); b {
			txArgs = append(txArgs, build.ClearAuthRequired())
		}
		if b, _ := cmd.Flags().GetBool("clear-auth-revocable"); b {
			txArgs = append(txArgs, build.ClearAuthRevocable())
		}
		if b, _ := cmd.Flags().GetBool("clear-auth-immutable"); b {
			txArgs = append(txArgs, build.ClearAuthImmutable())
		}
		if n, _ := cmd.Flags().GetInt("master-weight"); n >= 0 {
			txArgs = append(txArgs, build.MasterWeight(n))
		}
		if thresh, _ := cmd.Flags().GetIntSlice("thresholds"); len(thresh) == 3 {
			txArgs = append(txArgs, build.SetThresholds(uint32(thresh[0]), uint32(thresh[1]), uint32(thresh[2])))
		}
		if signer, _ := cmd.Flags().GetString("add-signer"); signer != "" {
			weight, _ := cmd.Flags().GetInt("signer-weight")
			txArgs = append(txArgs, build.AddSigner(signer, uint32(weight)))
		}
		if signer, _ := cmd.Flags().GetString("remove-signer"); signer != "" {
			txArgs = append(txArgs, build.RemoveSigner(signer))
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
	accountCmd.AddCommand(setOptionsCmd)
	setOptionsCmd.Flags().String("seed", "", "account seed")
	setOptionsCmd.Flags().String("inflation-destination", "", "set the inflation destination")
	setOptionsCmd.Flags().String("home-domain", "", "set the home domain")
	setOptionsCmd.Flags().Bool("set-auth-required", false, "set the auth-required option")
	setOptionsCmd.Flags().Bool("set-auth-revocable", false, "set the auth-revocable option")
	setOptionsCmd.Flags().Bool("set-auth-immutable", false, "set the auth-immutable option")
	setOptionsCmd.Flags().Bool("clear-auth-required", false, "unset the auth-required option")
	setOptionsCmd.Flags().Bool("clear-auth-revocable", false, "unset the auth-revocable option")
	setOptionsCmd.Flags().Bool("clear-auth-immutable", false, "unset the auth-immutable option")
	setOptionsCmd.Flags().Int("master-weight", -1, "set the master weight")
	setOptionsCmd.Flags().IntSlice("thresholds", []int{}, "set the account thresholds")
	setOptionsCmd.Flags().String("add-signer", "", "add a signer to this account")
	setOptionsCmd.Flags().Int("signer-weight", 1, "weight of the signer which will be added")
	setOptionsCmd.Flags().String("remove-signer", "", "remove a signer from this account")
}
