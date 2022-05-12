package test_persistence

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

type DummyMap map[string]any

func (d DummyMap) Clone() DummyMap {
	buf := make(DummyMap)
	for k, v := range d {
		buf[k] = v
	}
	return buf
}

func (d DummyMap) IsEqualId(id string) bool {
	if val, ok := d["Id"]; ok {
		if _id, ok := val.(string); ok {
			return _id == id
		}
	}
	return false
}

func (d DummyMap) GetId() string {
	if val, ok := d["Id"]; ok {
		if _id, ok := val.(string); ok {
			return _id
		}
	}
	return ""
}

func (d DummyMap) IsZeroId() bool {
	if val, ok := d["Id"]; ok {
		if _id, ok := val.(string); ok {
			return _id == ""
		}
	}
	return false
}

func (d DummyMap) WithId(id string) DummyMap {
	d["Id"] = id
	return d
}

func (d DummyMap) WithGeneratedId() DummyMap {
	d["Id"] = data.IdGenerator.NextLong()
	return d
}
