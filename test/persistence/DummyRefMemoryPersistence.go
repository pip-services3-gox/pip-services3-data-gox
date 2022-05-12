package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

// extends IdentifiableMemoryPersistence<Dummy, string>
// implements IDummyPersistence {
type DummyRefMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence[*DummyRef, string]
}

func NewDummyRefMemoryPersistence() *DummyRefMemoryPersistence {
	return &DummyRefMemoryPersistence{
		*cpersist.NewIdentifiableMemoryPersistence[*DummyRef, string](),
	}
}

func (c *DummyRefMemoryPersistence) GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[*DummyRef], err error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetPageByFilter(ctx, correlationId,
			func(item *DummyRef) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			},
			paging,
			func(a, b *DummyRef) bool {
				return len(a.Key) < len(b.Key)
			},
			nil,
		)
}

func (c *DummyRefMemoryPersistence) GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetCountByFilter(ctx, correlationId,
			func(item *DummyRef) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			})
}
