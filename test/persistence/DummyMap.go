package test_persistence

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type DummyMap struct {
	data.AnyValueMap
}

func (d DummyMap) Clone() DummyMap {
	return DummyMap{AnyValueMap: *data.NewAnyValueMap(d.Value())}
}

func (d DummyMap) IsEqualId(id string) bool {
	if val, ok := d.GetAsNullableString("Id"); ok {
		return val == id
	}
	return false
}

func (d DummyMap) GetId() string {
	return d.GetAsString("Id")
}

func (d DummyMap) IsZeroId() bool {
	return d.GetAsString("Id") == ""
}

func (d DummyMap) WithId(id string) DummyMap {
	d.Append(map[string]any{"Id": id})
	return d
}

func (d DummyMap) WithGeneratedId() DummyMap {
	d.Append(map[string]any{"Id": data.IdGenerator.NextLong()})
	return d
}
