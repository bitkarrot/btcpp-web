package helpers

import (
	"fmt"
)

func pullTypeAndLen(data []int) (uint, uint, []int, error) {
	if len(data) < 3 {
		return 0, 0, nil, fmt.Errorf("Data too short. len:%d, %v", len(data), data)
	}

	lenVal := data[1]*32 + data[2]
	if len(data) < 3+lenVal {
		return 0, 0, nil, fmt.Errorf("Data too short. exp: %d, len:%d %v", lenVal+3, len(data), data)
	}

	/* Return a copy of the slice :/ */
	var dst []int
	dst = append(dst, data[3:3+lenVal]...)
	return uint(data[0]), uint(lenVal), dst, nil
}

/*
Given an invoice, pull out just the fallback addresses.

	Note that assumes invoice is valid.
	FIXME: actually finish invoice validation etc
*/
func GetFallbackAddrs(invstring string) ([]string, error) {
	/* Split out HRP and data, for now ignore HRP */
	_, data, err := Decode(invstring)
	if err != nil {
		return nil, err
	}

	/* first 7 parts are the timestamp + last 104 are the signature */
	if len(data) < 7+104 {
		return nil, fmt.Errorf("invoice data too small. %s", invstring)
	}
	typeDatas := data[7 : len(data)-104]

	addrs := make([]string, 0)
	for len(typeDatas) > 0 {
		typeVal, lenVal, typeData, err := pullTypeAndLen(typeDatas)
		if err != nil {
			return nil, err
		}
		typeDatas = typeDatas[3+lenVal:]

		/* We're just looking for fallback onchain addresses */
		if typeVal == 9 {
			/* Convert to a bitcoin address */
			if typeData[0] >= 19 {
				/* ignored */
				continue
			}
			/* this is a witness version + data */
			if typeData[0] <= 16 {
				/* FIXME: what network are we on? */
				version := BECH32M
				if typeData[0] == 0 {
					version = BECH32
				}
				addr, err := Encode(version, "bcrt", typeData)
				if err != nil {
					return nil, err
				}
				addrs = append(addrs, addr)
			}
			/* P2PKH */
			if typeData[0] == 17 {
				/* Encode the rest as Base58 addr */
				/* FIXME: do this */
				return nil, fmt.Errorf("Unimplemented P2PKH")
			}
			/* P2SH */
			if typeData[0] == 18 {
				/* Encode the rest as Base58 addr */
				/* FIXME: do this */
				return nil, fmt.Errorf("Unimplemented P2SH")
			}
		}
	}

	return addrs, nil
}
