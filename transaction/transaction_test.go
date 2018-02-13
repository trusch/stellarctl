package transaction_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	yaml "gopkg.in/yaml.v2"

	. "github.com/trusch/stellarctl/transaction"
)

var _ = Describe("Transaction", func() {
	It("should be possible to load a transaction from YAML", func() {
		bs, err := ioutil.ReadFile("test-transaction.yaml")
		Expect(err).NotTo(HaveOccurred())
		_, err = FromYAML(bs)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should be possible to sign a transaction", func() {
		bs, err := ioutil.ReadFile("test-transaction.yaml")
		Expect(err).NotTo(HaveOccurred())
		tx, err := FromYAML(bs)
		Expect(err).NotTo(HaveOccurred())
		Expect(tx.Sign("SAMXFRJPYERQS42BAL7DSYP7IRFQW4S2G53RHQNCTYKIKFFBBTP3PKDD")).To(Succeed())
		Expect(len(tx.Signatures)).To(Equal(2))
	})

	It("should be possible to encode a signed transaction as XDR", func() {
		bs, err := ioutil.ReadFile("test-transaction.yaml")
		Expect(err).NotTo(HaveOccurred())
		tx, err := FromYAML(bs)
		Expect(err).NotTo(HaveOccurred())
		Expect(tx.Sign("SAMXFRJPYERQS42BAL7DSYP7IRFQW4S2G53RHQNCTYKIKFFBBTP3PKDD")).To(Succeed())
		xdr, err := tx.ToXDR()
		Expect(err).NotTo(HaveOccurred())
		Expect(xdr).NotTo(BeEmpty())
		yml, err := yaml.Marshal(tx)
		Expect(err).NotTo(HaveOccurred())
		fmt.Print(string(yml))
	})
})
