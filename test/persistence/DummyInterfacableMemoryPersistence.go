package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type DummyInterfacableMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence[DummyInterfacable, string]
}

func NewDummyInterfacableMemoryPersistence() *DummyInterfacableMemoryPersistence {
	return &DummyInterfacableMemoryPersistence{
		*cpersist.NewIdentifiableMemoryPersistence[DummyInterfacable, string](),
	}
}

func (c *DummyInterfacableMemoryPersistence) GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (cdata.DataPage[DummyInterfacable], error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetPageByFilter(ctx, correlationId,
			func(item DummyInterfacable) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			},
			paging,
			func(a, b DummyInterfacable) bool {
				return len(a.Key) < len(b.Key)
			},
			nil,
		)
}

func (c *DummyInterfacableMemoryPersistence) GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetCountByFilter(ctx, correlationId,
			func(item DummyInterfacable) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			})
}
