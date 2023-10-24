package rest

import (
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/permission"
	"path"
)

type KeyInfoService struct {
	requestFactory
	Synchronous
}

func (ks *KeyInfoService) Permissions() (result *permission.Snapshot, err error) {
	req, err := ks.requestFactory.NewAuthenticatedRequest(common.PermissionRead, path.Join("permissions"))
	if err != nil {
		return nil, err
	}
	raw, err := ks.Request(req)
	if err != nil {
		return nil, err
	}
	result, err = permission.SnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return result, nil
}
