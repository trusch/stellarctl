stellarctl
==========

The long missing stellar command line utility!

## Features

### functional

* create accounts
* send lumens and other assets
* manage trust
* issue assets
* generate keys
* set account options
  * inflation address
  * thresholds
  * signers
  * ...
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

* manage SDEX offers
* implement watch commands
* declarative transaction build system


## Examples

### Create a test account and fund it using friendbot:
```bash
> stellarctl account generate
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
> stellarctl account generate
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
