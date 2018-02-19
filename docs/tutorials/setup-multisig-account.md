Setup a multisignature account
=============================

## 1. Create your primary account

This can be done like explained in [create-cold-wallet](./create-cold-wallet.md) or in short:

```bash
> stellarctl account generate mnemonic --password "super secure"
Mnemonic: afford glow term mom have leave liquid electric leopard arctic outer extend perfect silly attract vacant chef cover noise dinosaur glide valid source frame
Account 0:
Seed: SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT
Addr: GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
---
> export PRIMARY_ID=GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
> export PRIMARY_SEED=SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT
> stellarctl account create --seed ${EXISTING_SEED} --id ${PRIMARY_ID}
```

This will be the account holding your assets.

## 2. create a secondary signer account

```bash
> stellarctl account generate mnemonic --password "another super secure passphrase"
Mnemonic: ice image cute birth airport pull lamp labor exact ginger stay build tell bonus surge display ice biology any muscle mobile genre word tunnel
Account 0:
Seed: SBQ5N2CR7P77S2S5Q62KYTYM4WVVPKXKLQRQRREGXL5YDWO223KTJFGJ
Addr: GDH2KP3RXGM5NYBYU7EX2JROWUCEJTS6QEBLGRED6XHU6SIZRXA6LF7M
---
> export SECONDARY_ID=GDH2KP3RXGM5NYBYU7EX2JROWUCEJTS6QEBLGRED6XHU6SIZRXA6LF7M
> export SECONDARY_SEED=SBQ5N2CR7P77S2S5Q62KYTYM4WVVPKXKLQRQRREGXL5YDWO223KTJFGJ
> stellarctl account create --seed ${PRIMARY_SEED} --id ${SECONDARY_ID}
```

## 3. add your secondary account as a signer to your primary account and adjust thresholds

```bash
> stellarctl account set-options \
    --seed ${PRIMARY_SEED} \
    --add-signer ${SECONDARY_ID} \
    --thresholds 0,2,2
```

## 4. use both seeds to perform transactions from the primary account

first get the current sequence number for your account and add one

```bash
> seq=$(stellarctl account info --id ${PRIMARY_ID} --format json | jq -r .sequence)
> expr $seq + 1
31535724356435971
```

create a transaction:

```yaml
source_account: GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
sequence_id: 31535724356435971
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
> stellarctl transaction sign --input tx.yaml --seed ${PRIMARY_SEED}
> stellarctl transaction sign --input tx.yaml --seed ${SECONDARY_SEED}
> stellarctl transaction commit --input tx.yaml
```
