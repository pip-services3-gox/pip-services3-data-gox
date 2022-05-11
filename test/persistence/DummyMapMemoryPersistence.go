package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type DummyMapMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence[DummyMap, string]
}

func NewDummyMapMemoryPersistence() *DummyMapMemoryPersistence {
	return &DummyMapMemoryPersistence{
		*cpersist.NewIdentifiableMemoryPersistence[DummyMap, string](),
	}
}

func filterFunc(filter cdata.FilterParams) func(item DummyMap) bool {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return func(value DummyMap) bool {
		if _val, ok := value.GetAsNullableString("Key"); ok {
			if key != "" && _val != key {
				return false
			}
			return true
		}

		return false
	}
}

func sortFunc(a, b DummyMap) bool {
	return len(a.GetAsString("Key")) < len(b.GetAsString("Key"))
}

func (c *DummyMapMemoryPersistence) GetPageByFilter(ctx context.Context, correlationId string,
	filter cdata.FilterParams, paging cdata.PagingParams) (result cdata.DataPage[DummyMap], err error) {

	return c.IdentifiableMemoryPersistence.
		GetPageByFilter(ctx, correlationId, filterFunc(filter), paging, sortFunc, nil)
}

func (c *DummyMapMemoryPersistence) GetCountByFilter(ctx context.Context, correlationId string,
	filter cdata.FilterParams) (count int64, err error) {

	return c.IdentifiableMemoryPersistence.
		GetCountByFilter(ctx, correlationId, filterFunc(filter))
}
