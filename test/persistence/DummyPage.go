package test_persistence

type DummyPage struct {
	Total *int64  `json:"total"`
	Data  []Dummy `json:"data"`
}

func NewEmptyDummyPage() *DummyPage {
	return &DummyPage{}
}

func NewDummyPage(total *int64, data []Dummy) *DummyPage {
	return &DummyPage{Total: total, Data: data}
}
