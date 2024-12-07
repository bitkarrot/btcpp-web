module github.com/base58btc/btcpp-web

go 1.21.4

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/alexedwards/scs/v2 v2.7.0
	github.com/base58btc/cln-checkout v0.0.0-00010101000000-000000000000
	github.com/base58btc/clnsocket v0.0.0-00010101000000-000000000000
	github.com/base58btc/mailer v0.0.0-20230510213939-04e8f45514b5
	github.com/chromedp/cdproto v0.0.0-20231126232103-8e31ff06e23b
	github.com/chromedp/chromedp v0.9.3
	github.com/google/go-cmp v0.5.9
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/schema v1.2.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/sorcererxw/go-notion v0.2.4
	github.com/stripe/stripe-go/v76 v76.7.0
	github.com/yassun/go-bip21 v0.0.0-20190813103626-13a183e13826
)

require (
	github.com/aead/siphash v1.0.1 // indirect
	github.com/btcsuite/btcd v0.23.5-0.20230905170901-80f5a0ffdf36 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.4-0.20230904040416-d4f519f5dc05 // indirect
	github.com/btcsuite/btcd/btcutil/psbt v1.1.8 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.3 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/btcwallet v0.16.10-0.20231129183218-5df09dd43358 // indirect
	github.com/btcsuite/btcwallet/wallet/txauthor v1.3.3 // indirect
	github.com/btcsuite/btcwallet/wallet/txrules v1.2.0 // indirect
	github.com/btcsuite/btcwallet/wallet/txsizes v1.2.3 // indirect
	github.com/btcsuite/btcwallet/walletdb v1.4.0 // indirect
	github.com/btcsuite/btcwallet/wtxmgr v1.5.0 // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/decred/dcrd/lru v1.1.2 // indirect
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.1 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kkdai/bstream v1.0.0 // indirect
	github.com/lightninglabs/gozmq v0.0.0-20191113021534-d20a764486bf // indirect
	github.com/lightninglabs/neutrino v0.16.0 // indirect
	github.com/lightninglabs/neutrino/cache v1.1.2 // indirect
	github.com/lightningnetwork/lnd v0.17.3-beta.rc1 // indirect
	github.com/lightningnetwork/lnd/clock v1.1.1 // indirect
	github.com/lightningnetwork/lnd/queue v1.1.1 // indirect
	github.com/lightningnetwork/lnd/ticker v1.1.1 // indirect
	github.com/lightningnetwork/lnd/tlv v1.1.2 // indirect
	github.com/lightningnetwork/lnd/tor v1.1.3 // indirect
	github.com/mailgun/mailgun-go/v4 v4.12.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.18 // indirect
	github.com/miekg/dns v1.1.57 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niftynei/lnsocket/go v0.0.0-20231126210829-f7651ea81661 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sendgrid/rest v2.6.9+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.14.0+incompatible // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/exp v0.0.0-20231127185646-65229373498e // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
)

replace github.com/base58btc/cln-checkout v0.0.0-00010101000000-000000000000 => ../cln-checkout/checkout

replace github.com/base58btc/clnsocket v0.0.0-00010101000000-000000000000 => ../cln-checkout/cln

replace github.com/niftynei/lnsocket/go v0.0.0-20231126210829-f7651ea81661 => ../lnsocket/go

replace github.com/sorcererxw/go-notion v0.2.4 => github.com/niftynei/go-notion v0.0.0-20230323155332-a2c93bab119e
