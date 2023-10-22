package fundingoffer

import (
	"encoding/json"
	"fmt"
)

type CancelRequest struct {
	ID int64
}

func (cr *CancelRequest) ToJSON() ([]byte, error) {
	resp := struct {
		ID int64 `json:"id"`
	}{
		ID: cr.ID,
	}
	return json.Marshal(resp)
}

// MarshalJSON converts the offer cancel object into the format required by the
// bitfinex websocket service.
func (cr *CancelRequest) MarshalJSON() ([]byte, error) {
	b, err := cr.ToJSON()
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("[0, \"foc\", null, %s]", string(b))), nil
}

type CancelAllRequest struct {
	Symbol string
}

func (cr *CancelAllRequest) ToJSON() ([]byte, error) {
	resp := struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: cr.Symbol,
	}
	return json.Marshal(resp)
}

func (cr *CancelAllRequest) MarshalJSON() ([]byte, error) {
	b, err := cr.ToJSON()
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("[0, \"foc\", null, %s]", string(b))), nil
}
