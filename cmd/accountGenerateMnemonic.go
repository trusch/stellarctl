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

	bip39 "github.com/bartekn/go-bip39"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stellar/go/exp/crypto/derivation"
	"github.com/stellar/go/keypair"
)

// accountGenerateMnemonicCmd represents the accountGenerateMnemonic command
var accountGenerateMnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "generate account keys from a mnemonic passphrase (sep-0005)",
	Long:  `generate account keys from a mnemonic passphrase (sep-0005)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic, _ := cmd.Flags().GetString("mnemonic")
		var err error
		if mnemonic == "" {
			mnemonic, err = generateNewMnemonic(cmd)
			if err != nil {
				return err
			}
		}
		accounts, err := generateAccounts(cmd, mnemonic)
		if err != nil {
			return err
		}
		fmt.Printf("Mnemonic: %v\n", mnemonic)
		for idx, acc := range accounts {
			fmt.Printf("Account %v:\nSeed: %v\nAddr: %v\n---\n", idx, acc.Seed(), acc.Address())
		}
		return nil
	},
}

func init() {
	accountGenerateCmd.AddCommand(accountGenerateMnemonicCmd)
	accountGenerateMnemonicCmd.Flags().String("mnemonic", "", "your mnemonic words")
	accountGenerateMnemonicCmd.Flags().Int("entropy", 256, "size of entropy when generating a new mnemonic")
	accountGenerateMnemonicCmd.Flags().Int("count", 1, "number of accounts to generate from the mnemonic")
	accountGenerateMnemonicCmd.Flags().String("password", "", "password for the accounts")
}

func generateNewMnemonic(cmd *cobra.Command) (string, error) {
	entropyBits, _ := cmd.Flags().GetInt("entropy")
	entropy, err := bip39.NewEntropy(entropyBits)
	if err != nil {
		return "", errors.Wrap(err, "Error generating entropy")
	}

	return bip39.NewMnemonic(entropy)
}

func generateAccounts(cmd *cobra.Command, mnemonic string) ([]*keypair.Full, error) {
	password, _ := cmd.Flags().GetString("password")
	count, _ := cmd.Flags().GetInt("count")
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, errors.New("Invalid words or checksum")
	}

	masterKey, err := derivation.DeriveForPath(derivation.StellarAccountPrefix, seed)
	if err != nil {
		return nil, errors.Wrap(err, "Error deriving master key")
	}

	keys := make([]*keypair.Full, count)
	for i := 0; i < count; i++ {
		key, err := masterKey.Derive(derivation.FirstHardenedIndex + uint32(i))
		if err != nil {
			return nil, errors.Wrap(err, "Error deriving child key")
		}

		kp, err := keypair.FromRawSeed(key.RawSeed())
		if err != nil {
			return nil, errors.Wrap(err, "Error creating key pair")
		}
		keys[i] = kp
	}
	return keys, nil
}
