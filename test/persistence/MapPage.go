package test_persistence

type MapPage struct {
	Total *int64                   `json:"total"`
	Data  []map[string]interface{} `json:"data"`
}

func NewEmptyMapPage() *MapPage {
	return &MapPage{}
}

func NewMapPage(total *int64, data []map[string]interface{}) *MapPage {
	return &MapPage{Total: total, Data: data}
}
