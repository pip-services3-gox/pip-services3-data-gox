package test_persistence

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type Dummy struct {
	Id      string `json:"id"`
	Key     string `json:"key"`
	Content string `json:"content"`
}

func (d Dummy) Clone() Dummy {
	return Dummy{
		Id:      d.Id,
		Key:     d.Key,
		Content: d.Content,
	}
}

func (d Dummy) IsEqualId(id string) bool {
	return d.Id == id
}

func (d Dummy) GetId() string {
	return d.Id
}

func (d Dummy) IsZeroId() bool {
	return d.Id == ""
}

func (d Dummy) WithId(id string) Dummy {
	d.Id = id
	return d
}

func (d Dummy) WithGeneratedId() Dummy {
	d.Id = data.IdGenerator.NextLong()
	return d
}

//func (d *Dummy) Merge(t *Dummy) *Dummy {
//	if t.Key != "" {
//		d.Key = t.Key
//	}
//	if t.Content != "" {
//		d.Content = t.Content
//	}
//	return d
//}
