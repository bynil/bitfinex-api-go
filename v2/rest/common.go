package rest

type PageQuery struct {
	Start int64
	End   int64
	Limit int
}

func (pq *PageQuery) ToPayload() map[string]interface{} {
	payload := map[string]interface{}{}
	if pq == nil {
		return payload
	}
	if pq.Start > 0 {
		payload["start"] = pq.Start
	}
	if pq.End > 0 {
		payload["end"] = pq.End
	}
	if pq.Limit > 0 {
		payload["limit"] = pq.Limit
	}
	return payload
}
