package rest

import (
	"fmt"
	"path"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/ledger"
)

// LedgerService manages the Ledgers endpoint.
type LedgerService struct {
	requestFactory
	Synchronous
}

var maxLimit int32 = 2500

// Ledgers - all of the past ledger entreies
// see https://docs.bitfinex.com/reference#ledgers for more info
func (s *LedgerService) Ledgers(currency string, category *int32, page *PageQuery) (*ledger.Snapshot, error) {
	if page != nil && page.Limit > int(maxLimit) {
		return nil, fmt.Errorf("Max request limit:%d, got: %d", maxLimit, page.Limit)
	}

	payload := page.ToPayload()
	if category != nil {
		payload["category"] = *category
	}
	req, err := s.requestFactory.NewAuthenticatedRequestWithData(common.PermissionRead, path.Join("ledgers", currency, "hist"), payload)
	if err != nil {
		return nil, err
	}

	raw, err := s.Request(req)
	if err != nil {
		return nil, err
	}

	lss, err := ledger.SnapshotFromRaw(raw, ledger.FromRaw)
	if err != nil {
		return nil, err
	}

	return lss, nil
}
