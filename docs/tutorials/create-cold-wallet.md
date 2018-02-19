Create Cold Wallet
==================

This tutorial will explain how you create a safe, cold wallet using `stellarctl`

## 1. Create Account Keys

First you will need to generate a new mnemonic passphrase and choose a password to protect your account. Remember: stronger passwords mean more security!

```bash
> stellarctl account generate mnemonic --password "super secure password"
Mnemonic: afford glow term mom have leave liquid electric leopard arctic outer extend perfect silly attract vacant chef cover noise dinosaur glide valid source frame
Account 0:
Seed: SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT
Addr: GDA4ZFNYLG73L5TEOXDGWBPCNKZEHX355DS7EOADKVAMZ6K7NIRLLCU7
---
```

## 2. Write down your mnemonic words on paper.

Do this very carefully. If you just make a single mistake your account will be lost. Also please don't write down your password. This is the only thing which will prevent an attacker which knows your mnemonic words from accessing your account. After writing them down, put the paper in a safe or bring it to your bank and let them keep it safe for you.

## 3. Remember your password!

Once again, you NEED your password to recover your account from the mnemonic words!

## 4. Create your account on the network

Next step is to create your account in the stellar blockchain. To do this, you need to create a transaction with a create-account operation. Unfortunately you need a working, funded account for this. You can beg on reddit, or you can buy some lumen on an exchange of your choice and withdraw to your account address. Most exchanges will create the account for you if it doesn't exist on the network. If you already own a funded account you can use `stellarctl` to create your new account:

```bash
> stellarctl account create --id ${YOUR_NEW_ACCOUNT_ADDRESS} --seed ${YOUR_FUNDING_ACCOUNT_SEED} --amount 10
```

Another good way to do this is to spend some tips you earned on reddit to fund your account. The stellar tip bot creates accounts when withdrawing from it. This also has the advantage that a reddit account is quiet anonymous, so you get basically an anonymous account in contrast to when funding from an exchange.

## 5. (Optionally) setup plausable deniability

This will setup another account from the same mnemonic words just by using a different password. The reason to do so is that you can then tell the second password to an attacker which already stole your mnemonic words and threats you to give him the password. Basically you just generate an account from the same mnemonic words, but with another password and fund it with a little amount of XLM.

```bash
> stellarctl account generate mnemonic \
    --password "a secondary password" \
    --mnemonic "afford glow term mom have leave liquid electric leopard arctic outer extend perfect silly attract vacant chef cover noise dinosaur glide valid source frame"
Mnemonic: afford glow term mom have leave liquid electric leopard arctic outer extend perfect silly attract vacant chef cover noise dinosaur glide valid source frame
Account 0:
Seed: SDPIUGHFFBL64EQLHSU3GEXPOBATCPYFLX3A7ZDPXGQRVUX6TDPZDTEO
Addr: GBAFIE4L5KOWXBY47VY3D2AZNTNG2LX27RLXEFA2YM56QWX56LRMBART

> stellarctl account create \
    --id GBAFIE4L5KOWXBY47VY3D2AZNTNG2LX27RLXEFA2YM56QWX56LRMBART \
    --seed SAXK2WMD3T7FNZRIOJOAMV2OMY3WRX5JZHVUT7IJHLUD7ZFXTDJWBJNT \
    --amount 10
```

No you have a secondary account and an attacker could not say that this is not the only account you own.
