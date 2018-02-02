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

## ToDo

* manage SDEX offers
* implement watch commands
* declarative transaction build system

## Help commands (some of them)
```
➜  stellarctl git:(master) ./stellarctl help
a tool to interact with the stellar network

Usage:
  stellarctl [command]

Available Commands:
  account     interact with accounts
  help        Help about any command
  send        send assets
  transaction interact with transactions
  trust       upsert a trustline for an asset

Flags:
      --config string   config file (default is $HOME/.stellarctl.yaml)
      --format string   output format (default "yaml")
  -h, --help            help for stellarctl
      --testnet         use the testnet

Use "stellarctl [command] --help" for more information about a command.
➜  stellarctl git:(master) ./stellarctl account help
interact with accounts

Usage:
  stellarctl account [command]

Available Commands:
  address     get account address for a given seed
  create      create a new account
  generate    generate a new key pair
  info        infos about a account
  set-options set account options
  testfill    testfill a newly created account

Flags:
  -h, --help   help for account

Global Flags:
      --config string   config file (default is $HOME/.stellarctl.yaml)
      --format string   output format (default "yaml")
      --testnet         use the testnet

Use "stellarctl account [command] --help" for more information about a command.
➜  stellarctl git:(master) ./stellarctl send help   
Error: encoded value is 0 bytes; minimum valid length is 3
Usage:
  stellarctl send [flags]

Flags:
      --amount string         amount (default "0")
      --asset-code string     asset code
      --asset-issuer string   asset issuer
      --from string           source account seed
  -h, --help                  help for send
      --to string             destination account address

Global Flags:
      --config string   config file (default is $HOME/.stellarctl.yaml)
      --format string   output format (default "yaml")
      --testnet         use the testnet

```
