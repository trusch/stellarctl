`stellarctl`
==========

The long missing stellar command line utility! Finally a swiss army knife every lumenaut should have in their pockets.

## Features

### high level

* setup multi sign wallets with mnemonic codes according do SEP-0005
* send and receive assets without the need to trust a third party or browser
* buy and sell any asset on SDEX
* efficently inspect the stellar network state including balances and transaction details
* craft arbitary transactions using human readable yaml files
* query coinmarketcap from the commandline ;)

### full list

* create accounts
* send lumens and other assets
* manage trust
* issue assets
* manage SDEX offers
* declarative transaction build system
* generate keys
* generate mnemonic codes and accounts according to SEP-0005
* set account options
  * inflation address
  * thresholds
  * signers
  * flags
  * home domain
  * master weight
* inspect ledger for account details
  * transactions
  * operations
  * effects
* inspect transaction details
  * operations
  * effects
* create test accounts using friendbot

### non-functional

* single go binary
* uses official go SDK as far as possible
* simple command line syntax

### Todo

* implement watch commands

## Install
To install `stellarctl` go to the [Releases Page](https://github.com/trusch/stellarctl/releases) of this project and download the binary matching your operating system and architecture. Place this executable somewhere where your shell will recognize it. On linux and darwin (sane operating systems ;)) just move it somewhere inside your PATH. No you are ready to use the tool!

## Examples

### Create a test account and fund it using friendbot:
```bash
> stellarctl account generate random
Seed: SB7W36WTOVQCHH5Q5DMUF7HESWGOUIXMPIGXWWCIGS4GR43MJF5YJI4H
Address: GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO
> export ACCOUNT_ONE_SEED=SB7W36WTOVQCHH5Q5DMUF7HESWGOUIXMPIGXWWCIGS4GR43MJF5YJI4H
> export ACCOUNT_ONE_ID=GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO
> stellarctl account testfill --id ${ACCOUNT_ONE_ID} --testnet
> stellarctl account info --id ${ACCOUNT_ONE_ID} --testnet
balances:
- balance: "10000.0000000"
  limit: ""
  asset:
    type: native
    code: ""
    issuer: ""
data: {}
flags:
  authrequired: false
  authrevocable: false
home_domain: ""
id: GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO
inflations_destination: ""
sequence: "30428520442232832"
signers:
- publickey: GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO
  weight: 1
  key: GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO
  type: ed25519_public_key
subentry_count: 0
thresholds:
  lowthreshold: 0
  medthreshold: 0
  highthreshold: 0
```

### Create an account and fund it using an existing account:
```bash
> stellarctl account generate random
Seed: SB6ZPJIKHAM4EFWEWHC6GA76T4EC7NISGBXH52KMDV4NYKVZUFASFREU
Address: GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ
> export ACCOUNT_TWO_SEED=SB6ZPJIKHAM4EFWEWHC6GA76T4EC7NISGBXH52KMDV4NYKVZUFASFREU
> export ACCOUNT_TWO_ID=GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ
> stellarctl account create \
    --id ${ACCOUNT_TWO_ID} \
    --seed ${ACCOUNT_ONE_SEED} \
    --amount 10 \
    --testnet
> stellarctl account info --id ${ACCOUNT_TWO_ID} --testnet
balances:
- balance: "10.0000000"
  limit: ""
  asset:
    type: native
    code: ""
    issuer: ""
data: {}
flags:
  authrequired: false
  authrevocable: false
home_domain: ""
id: GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ
inflations_destination: ""
sequence: "30429014363471872"
signers:
- publickey: GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ
  weight: 1
  key: GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ
  type: ed25519_public_key
subentry_count: 0
thresholds:
  lowthreshold: 0
  medthreshold: 0
  highthreshold: 0
```

### Send XLM
```bash
> stellarctl send --from ${ACCOUNT_ONE_SEED} --to ${ACCOUNT_TWO_ID} --amount 100 --testnet
> stellarctl account info --id ${ACCOUNT_TWO_ID} --testnet --format json | jq ".balances[0].balance"
"110.0000000"
```

### Create a trustline and receive given asset
```bash
> stellarctl trust \
    --code YLM --issuer ${ACCOUNT_ONE_ID} \
    --seed ${ACCOUNT_TWO_SEED} --testnet
> stellarctl send --from ${ACCOUNT_ONE_SEED} --to ${ACCOUNT_TWO_ID} \
    --asset-code YLM --asset-issuer ${ACCOUNT_ONE_ID} --amount 100 --testnet
> stellarctl account info --id ${ACCOUNT_TWO_ID} --testnet --format json | jq '.balances[] | select(.asset_code == "YLM")'
{
  "balance": "100.0000000",
  "limit": "922337203685.4775807",
  "asset_type": "credit_alphanum4",
  "asset_code": "YLM",
  "asset_issuer": "GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO"
}
```

### Get effects of the last transaction of an account
```bash
> tid=$(stellarctl account info --id ${ACCOUNT_ONE_ID} transactions --testnet --limit 1 --format json | jq '.[0].id' -r)
>  stellarctl transaction info --id ${tid} effects --testnet --format json | jq '.[] | {type: .type, account: .account, amount: .amount}'
{
  "type": "account_credited",
  "account": "GB3AS47AUAHGOQYSJTS3KNFH5XLIKKESVBGM4BZJU55YFSARVFLOYPHZ",
  "amount": "100.0000000"
}
{
  "type": "account_debited",
  "account": "GCNZFKUUCSCAYP5JIMUJZRFQEMCCJN2YAZGQR6I6NJMDTXWCCEETIWIO",
  "amount": "100.0000000"
}
```

### Send a payment with a text memo
```bash
> stellarctl send --from ${ACCOUNT_ONE_SEED} --to ${ACCOUNT_TWO_ID} --amount 10 --memo "some tips for you" --testnet
> stellarctl account info --id ${ACCOUNT_TWO_ID} transactions \
    --limit 1 --testnet --format json | jq '.[0] | {id: .id, memo: .memo}'
{
  "id": "11849dcc909c12432d931645a629963c189891d8e050eba9a6e0b789ba6bbfc9",
  "memo": "some tips for you"
}
> # just curious what effect on my account this transaction had...
> stellarctl transaction info --id 11849dcc909c12432d931645a629963c189891d8e050eba9a6e0b789ba6bbfc9 effects \
    --testnet --format json | jq ".[] | select(.account == \"${ACCOUNT_TWO_ID}\") | {type: .type, amount: .amount}"
{
  "type": "account_credited",
  "amount": "10.0000000"
}
```

### Generate a new mnemonic passphrase and accounts according to SEP-0005
```bash
> stellarctl account generate mnemonic --password "some secret salting" --count 3
Mnemonic: clarify ritual pony physical one juice coconut nurse oval enhance also shrug sentence speed until climb camera setup engine trigger loop town match around
Account 0:
Seed: SCPJXFDG4YKE5UPRR3USWUBPS3OHU6TFKWL2UABM55RRD4YODFXZX6NF
Addr: GAQRBOI7G3FHNRXYJ63WQICLERXPBRUEZV3336QQ457BXDCHH2MPEMCS
---
Account 1:
Seed: SAWA6VVFH3UJND34AA6PSJXD7DITUXOLE5K6ZU6XKV7W2WNPPK3XG2U6
Addr: GBZOFASYWAQ57G3B4IUZKH3PP7N3H5KZJGSOUG5CGFCLMWHBYP7EMZVU
---
Account 2:
Seed: SDGJDPSES7IW3Z7FJ3363NASQRII5VLCBPPEZJJZLJNTVNQOMOZOFXA7
Addr: GAM3CNPKKRHRAAEE3TNPEZKQWZPLJ44RI247VZGKLAMNM5PQXKHM5SMC
---
```

### Restore accounts from a mnemonic passphrase according to SEP-0005
```bash
> stellarctl account generate mnemonic --mnemonic "clarify ritual pony physical one juice coconut nurse oval enhance also shrug sentence speed until climb camera setup engine trigger loop town match around" --password "some secret salting" --count 3
Mnemonic: clarify ritual pony physical one juice coconut nurse oval enhance also shrug sentence speed until climb camera setup engine trigger loop town match around
Account 0:
Seed: SCPJXFDG4YKE5UPRR3USWUBPS3OHU6TFKWL2UABM55RRD4YODFXZX6NF
Addr: GAQRBOI7G3FHNRXYJ63WQICLERXPBRUEZV3336QQ457BXDCHH2MPEMCS
---
Account 1:
Seed: SAWA6VVFH3UJND34AA6PSJXD7DITUXOLE5K6ZU6XKV7W2WNPPK3XG2U6
Addr: GBZOFASYWAQ57G3B4IUZKH3PP7N3H5KZJGSOUG5CGFCLMWHBYP7EMZVU
---
Account 2:
Seed: SDGJDPSES7IW3Z7FJ3363NASQRII5VLCBPPEZJJZLJNTVNQOMOZOFXA7
Addr: GAM3CNPKKRHRAAEE3TNPEZKQWZPLJ44RI247VZGKLAMNM5PQXKHM5SMC
---
```

### Manage an SDEX offer
```bash
> stellarctl trust \
    --seed ${ACCOUNT_ONE_SEED} \
    --code BTC \
    --issuer ${ACCOUNT_TWO_ID} \
    --testnet
> stellarctl offer create \
    --seed ${ACCOUNT_ONE_SEED} \
    --buying-asset-code BTC \
    --buying-asset-issuer ${ACCOUNT_TWO_ID} \
    --selling-asset-code XLM \
    --amount 0.005 \
    --price 0.0000465 \
    --testnet
> stellarctl offer list --id ${ACCOUNT_ONE_ID} --testnet
offers:
- links:
    self:
      href: https://horizon-testnet.stellar.org/offers/101741
      templated: false
    offermaker:
      href: https://horizon-testnet.stellar.org/accounts/GAKJROQ2PUBACRYL5ZOODBV5HDGIZ7CKU3RTEQ4NCLFXRDMZWHBSULRG
      templated: false
  id: 101741
  pt: "101741"
  seller: GAKJROQ2PUBACRYL5ZOODBV5HDGIZ7CKU3RTEQ4NCLFXRDMZWHBSULRG
  selling:
    type: native
    code: ""
    issuer: ""
  buying:
    type: credit_alphanum4
    code: BTC
    issuer: GAK3IXVT3CFUXKW2NRVZIB4ADH2NR3WZJ2VUFQFQQPD2NWOKFAG7CO6N
  amount: "0.0500000"
  pricer:
    "n": 57
    d: 1250000
  price: "0.0000456"
```

### Create a payment transaction from scratch

first get the current sequence number for your account and add one

```bash
> seq=$(stellarctl account info --id ${ACCOUNT_ONE_ID} --format json --testnet | jq -r .sequence)
> expr $seq + 1
31535724356435971
```

now create a transaction in a file called tx.yaml containing the sequence number from above

```yaml
source_account: GCBMYFXSK3SYC2WTFAMJ2ERO6DY6AQR67DPUIOBXLZOSB543QAMRRZBS
sequence_id: 31535724356435971
network: test
memo: "payment for you!"
operations:
- type: payment
  destination: GBDT3K42LOPSHNAEHEJ6AVPADIJ4MAR64QEKKW2LQPBSKLYD22KUEH4P
  asset:
    code: XLM
  amount: "10"
```

Now you can sign and commit it!

```bash
> stellarctl transaction sign --input tx.yaml --output tx.signed.yaml --seed ${ACCOUNT_ONE_SEED}
> stellarctl transaction commit --input tx.signed.yaml --testnet
```

see [test-transaction.yaml](./transaction/test-transaction.yaml) for a full list of supported operations (should be everything what's currently possible in stellar 😅)
