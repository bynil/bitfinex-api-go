package user

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/convert"
)

type Info struct {
	ID                int64
	Email             string
	Username          string
	Verified          bool
	VerificationLevel int
}

func FromRaw(raw []interface{}) (info *Info, err error) {
	if len(raw) < 55 {
		return info, fmt.Errorf("data slice too short for info: %#v", raw)
	}

	info = &Info{
		ID:                convert.I64ValOrZero(raw[0]),
		Email:             convert.SValOrEmpty(raw[1]),
		Username:          convert.SValOrEmpty(raw[2]),
		Verified:          convert.BValOrFalse(raw[4]),
		VerificationLevel: convert.IValOrZero(raw[5]),
	}

	return
}
