# csv-ynab-go

Convert exported CSV to a format supported by YNAB4 (old YNAB)

## Disclaimer

This repository is "as is". It is following format of my bank (especially the part where real payment date is in the memo in some cases where payments take time to clear). I'm adding some configuration options without any warranty it works, but if you'd like to fork and create a PR - be my guest.

This is also my first attempt at Go with mostly googling around and asking ChatGPT, and I'm sure the code could be more optimal, but it serves my purpose.

## Why?

Some of us old-timers have a "lifetime" license of YNAB4 from back in the day. YNAB has since gambled and changed their pricing model into subscription based, causing dissatisfaction with existing users while causing the old version obsolete. The new YNAB, a.k.a. NYNAB is a beautiful product with many cool features, but if you're like me and don't like 1000 subscriptions draining money.

While there are limitations on syncing, I'm usually doing family budget on a scheduled basis and I don't need it on the go or other devices.

I also don't want my transaction logs and balances on a 3rd party company's servers, when I can have my budget on an encrypted Cryptomator volume.

If you can get your hands on YNAB4 install, and have a license key somewhere in your email, then YNAB4 is perfectly capable of handling all the basics that one needs out of a budgeting software.

For this reason, the only cumbersome thing that's left for me is getting transactions from my bank. Luckily it supports exports to CSV, which is where this tool comes in.

## Prerequisites

- [Go](https://go.dev/doc/)

## How to use

1. clone this repo
2. export semicolon separated CSV from your bank
3. place the CSV in the `working-dir` of this repo
4. in the terminal, run `go run .`
5. Open YNAB4, select target account, then click on "Import" button at the top
6. Select converted csv file from the `working-dir` and hit open
7. In a popup window, configure correct date format and other options
8. Click "Import"

The transactions should now be in the list.

## Useful links

### Running YNAB4 on MacOS M1

[YNAB4-64bit script](https://github.com/banesto/YNAB4-64bit)

### Encrypting directories

[Cryptomator](https://github.com/cryptomator/cryptomator)
