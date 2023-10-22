package rest

import (
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/user"
	"path"
)

type UserService struct {
	requestFactory
	Synchronous
}

func (us *UserService) Info() (info *user.Info, err error) {
	req, err := us.requestFactory.NewAuthenticatedRequest(common.PermissionRead, path.Join("info/user"))
	if err != nil {
		return nil, err
	}
	raw, err := us.Request(req)
	if err != nil {
		return nil, err
	}
	info, err = user.FromRaw(raw)
	if err != nil {
		return nil, err
	}
	return info, nil
}
