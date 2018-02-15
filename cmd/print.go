package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stellar/go/clients/horizon"
	yaml "gopkg.in/yaml.v2"
)

func print(arg interface{}) error {
	obj := make(map[string]interface{})
	switch arg.(type) {
	case horizon.Account:
		{
			acc := arg.(horizon.Account)
			obj = map[string]interface{}{
				"id":                     acc.ID,
				"sequence":               acc.Sequence,
				"subentry_count":         acc.SubentryCount,
				"inflations_destination": acc.InflationDestination,
				"home_domain":            acc.HomeDomain,
				"thresholds":             acc.Thresholds,
				"flags":                  acc.Flags,
				"balances":               acc.Balances,
				"signers":                acc.Signers,
				"data":                   acc.Data,
			}
		}
	case io.Reader:
		{
			reader := json.NewDecoder(arg.(io.Reader))
			if err := reader.Decode(&obj); err != nil {
				return err
			}
		}
	case horizon.OrderBookSummary:
		{
			book := arg.(horizon.OrderBookSummary)
			obj["asks"] = book.Asks
			obj["bids"] = book.Bids
			obj["buying"] = book.Buying
			obj["selling"] = book.Selling
		}
	case horizon.OffersPage:
		{
			offers := arg.(horizon.OffersPage)
			obj["offers"] = offers.Embedded.Records
		}
	default:
		{
			if err := mapstructure.Decode(arg, obj); err != nil {
				return err
			}
		}
	}
	removeLinks(obj)
	return printWithFormat(removeEmbeddedRecordsPreamble(obj))
}

func printWithFormat(arg interface{}) error {
	format := viper.GetString("format")
	switch format {
	case "yaml":
		{
			bs, err := yaml.Marshal(arg)
			if err != nil {
				return err
			}
			fmt.Print(string(bs))
		}
	case "json":
		{
			bs, err := json.Marshal(arg)
			if err != nil {
				return err
			}
			fmt.Print(string(bs))
		}
	case "jsonpretty":
		{
			bs, err := json.MarshalIndent(arg, "", "  ")
			if err != nil {
				return err
			}
			fmt.Print(string(bs))
		}
	default:
		{
			return fmt.Errorf("unknown format '%v'", format)
		}
	}
	return nil
}

func removeLinks(arg interface{}) {
	switch arg.(type) {
	case map[string]interface{}:
		{
			m := arg.(map[string]interface{})
			delete(m, "links")
			delete(m, "_links")
			for key := range m {
				removeLinks(m[key])
			}
		}
	case []interface{}:
		{
			a := arg.([]interface{})
			for key := range a {
				removeLinks(a[key])
			}
		}
	}
}

func removeEmbeddedRecordsPreamble(arg map[string]interface{}) interface{} {
	if embedded, ok := arg["_embedded"]; ok {
		if embeddedMap, ok := embedded.(map[string]interface{}); ok {
			if list, ok := embeddedMap["records"]; ok {
				return list
			}
		}
	}
	return arg
}
