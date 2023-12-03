package helpers

import (
	"encoding/base64"
	"fmt"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
	bip21 "github.com/yassun/go-bip21"
)

func MakeBIP21QR(btcAmount float64, label string, onchainAddr string) (string, error) {

	u := &bip21.URIResources{
		UrnScheme: "bitcoin",
		Address:   strings.ToUpper(onchainAddr),
		Amount:    btcAmount,
		Label:     label,
	}

	bip21uri, err := u.BuildURI()
	if err != nil {
		return "", err
	}

	bip21png, err := qrcode.Encode(bip21uri, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bip21png), nil
}

func MakeBoltQR(invoice string) (string, error) {
	bolturi := fmt.Sprintf("lightning:%s", strings.ToUpper(invoice))
	boltpng, err := qrcode.Encode(bolturi, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(boltpng), nil
}
