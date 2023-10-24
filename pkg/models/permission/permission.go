package permission

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/convert"
)

type Permission struct {
	Scope string
	Read  bool
	Write bool
}

type Snapshot struct {
	Snapshot []*Permission
}

func FromRaw(raw []interface{}) (p *Permission, err error) {
	if len(raw) < 3 {
		return nil, fmt.Errorf("data slice too short for permission: %#v", raw)
	}
	p = &Permission{
		Scope: convert.SValOrEmpty(raw[0]),
		Read:  convert.BValOrFalse(raw[1]),
		Write: convert.BValOrFalse(raw[2]),
	}
	return p, nil
}

func SnapshotFromRaw(raw []interface{}) (*Snapshot, error) {
	if len(raw) == 0 {
		return &Snapshot{}, nil
	}
	snapshot := make([]*Permission, 0)
	for _, v := range raw {
		if v, ok := v.([]interface{}); ok {
			s, err := FromRaw(v)
			if err != nil {
				return nil, err
			}
			snapshot = append(snapshot, s)
		} else {
			return nil, fmt.Errorf("Invalid permission snapshot")
		}
	}

	return &Snapshot{Snapshot: snapshot}, nil
}
