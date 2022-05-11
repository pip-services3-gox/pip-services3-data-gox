package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

// extends IdentifiableMemoryPersistence
// implements IDummyPersistence
type DummyMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence[Dummy, string]
}

func NewDummyMemoryPersistence() *DummyMemoryPersistence {
	return &DummyMemoryPersistence{
		*cpersist.NewIdentifiableMemoryPersistence[Dummy, string](),
	}
}

// TODO:: remove after complete
//func (c *DummyMemoryPersistence) GetListByIds(ctx context.Context, correlationId string, ids []string) (items []Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.GetListByIds(ctx, correlationId, ids)
//}
//
//func (c *DummyMemoryPersistence) GetOneById(ctx context.Context, correlationId string, id string) (item Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.GetOneById(ctx, correlationId, id)
//}
//
//func (c *DummyMemoryPersistence) Create(ctx context.Context, correlationId string, item Dummy) (result Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.Create(ctx, correlationId, item)
//}
//
//func (c *DummyMemoryPersistence) Update(ctx context.Context, correlationId string, item Dummy) (result Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.Update(ctx, correlationId, item)
//}
//
//func (c *DummyMemoryPersistence) UpdatePartially(ctx context.Context, correlationId string, id string, data Dummy) (item Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.UpdatePartially(ctx, correlationId, id, data)
//}
//
//func (c *DummyMemoryPersistence) DeleteById(ctx context.Context, correlationId string, id string) (item Dummy, err error) {
//	return c.IdentifiableMemoryPersistence.DeleteById(ctx, correlationId, id)
//}
//
//func (c *DummyMemoryPersistence) DeleteByIds(ctx context.Context, correlationId string, ids []string) (err error) {
//	return c.IdentifiableMemoryPersistence.DeleteByIds(ctx, correlationId, ids)
//}

func (c *DummyMemoryPersistence) GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (cdata.DataPage[Dummy], error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetPageByFilter(ctx, correlationId,
			func(item Dummy) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			},
			paging,
			func(a, b Dummy) bool {
				return len(a.Key) < len(b.Key)
			},
			nil,
		)
}

func (c *DummyMemoryPersistence) GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error) {

	var key string

	if _key, ok := filter.GetAsNullableString("Key"); ok {
		key = _key
	}

	return c.IdentifiableMemoryPersistence.
		GetCountByFilter(ctx, correlationId,
			func(item Dummy) bool {
				if key != "" && item.Key != key {
					return false
				}
				return true
			})
}
