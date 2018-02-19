How to make an simple ICO
=========================

This guide will show you how easily you can setup your own ICO (initial coin offering) on stellar using only `stellarctl`.
Here we will create an ICO for 1.000.000 of the new and famous SMPL token (SaMPLe token) which we will sell for 5XLM each.

## 1. Create issuer account

```bash
> stellarctl account generate mnemonic --password "super secret"
Mnemonic: afford glow term mom have leave liquid electric leopard arctic outer extend perfect silly attract vacant chef cover noise dinosaur glide valid source frame
Account 0:
Seed: SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT
Addr: GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
---
> export ISSUER_SEED=SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT
> export ISSUER_ID=GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
```

## 2. Create an offer for your tokens on SDEX

```bash
> stellarctl offer create \
    --seed ${ISSUER_SEED} \
    --selling-asset-code SMPL \
    --selling-asset-issuer ${ISSUER_ID} \
    --amount 1000000 \
    --buying-asset-code XLM
    --price 5
```

"price" means that you want 5 XLM for one SMPL token.
"amount" means that you want to sell one million of SMPL tokens.

## 3. Marketing time!

Now your ICO is live and the only thing left is to make people buy your tokens. So now its time to create a webpage, make social media advertising and so on.
You want as much people to do step 4 as you can find.

## 4. Investors buy tokens

You investors need to trust you and place an buy order against your sell order.

```bash
> stellarctl trust \
    --seed ${INVESTOR_SEED}
    --code SMPL \
    --issuer ${ISSUER_ID}
> stellarctl offer create \
    --seed ${INVESTOR_SEED} \
    --buying-asset-code SMPL \
    --buying-asset-issuer ${ISSUER_ID} \
    --amount 1000 \
    --selling-asset-code XLM
    --price 0.2
```

"price" means that you want to want to get 0.2 SMPL token for one XLM.
"amount" means that you want to spend 1000 XLM.
