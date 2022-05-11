package test_persistence

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type DummyRef struct {
	Id      string `json:"id"`
	Key     string `json:"key"`
	Content string `json:"content"`
}

func (d *DummyRef) Clone() *DummyRef {
	return &DummyRef{
		Id:      d.Id,
		Key:     d.Key,
		Content: d.Content,
	}
}

func (d *DummyRef) IsEqualId(id string) bool {
	return d.Id == id
}

func (d *DummyRef) GetId() string {
	return d.Id
}

func (d *DummyRef) IsZeroId() bool {
	return d.Id == ""
}

func (d *DummyRef) WithId(id string) *DummyRef {
	d.Id = id
	return d
}

func (d *DummyRef) WithGeneratedId() *DummyRef {
	d.Id = data.IdGenerator.NextLong()
	return d
}
