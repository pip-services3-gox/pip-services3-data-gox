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

//func (c *DummyRefMemoryPersistence) Create(correlationId string, item *Dummy) (result *Dummy, err error) {
//	value, err := c.IdentifiableMemoryPersistence.Create(correlationId, item)
//
//	if value != nil {
//		val, _ := value.(*Dummy)
//		result = val
//	}
//	return result, err
//}
//
//func (c *DummyRefMemoryPersistence) GetListByIds(correlationId string, ids []string) (items []*Dummy, err error) {
//	convIds := make([]interface{}, len(ids))
//	for i, v := range ids {
//		convIds[i] = v
//	}
//	result, err := c.IdentifiableMemoryPersistence.GetListByIds(correlationId, convIds)
//	items = make([]*Dummy, len(result))
//	for i, v := range result {
//		val, _ := v.(*Dummy)
//		items[i] = val
//	}
//	return items, err
//}
//
//func (c *DummyRefMemoryPersistence) GetOneById(correlationId string, id string) (item *Dummy, err error) {
//	result, err := c.IdentifiableMemoryPersistence.GetOneById(correlationId, id)
//	if result != nil {
//		val, _ := result.(*Dummy)
//		item = val
//	}
//	return item, err
//}
//
//func (c *DummyRefMemoryPersistence) Update(correlationId string, item *Dummy) (result *Dummy, err error) {
//	value, err := c.IdentifiableMemoryPersistence.Update(correlationId, item)
//	if value != nil {
//		val, _ := value.(*Dummy)
//		result = val
//	}
//	return result, err
//}
//
//func (c *DummyRefMemoryPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item *Dummy, err error) {
//	result, err := c.IdentifiableMemoryPersistence.UpdatePartially(correlationId, id, data)
//
//	if result != nil {
//		val, _ := result.(*Dummy)
//		item = val
//	}
//	return item, err
//}
//
//func (c *DummyRefMemoryPersistence) DeleteById(correlationId string, id string) (item *Dummy, err error) {
//	result, err := c.IdentifiableMemoryPersistence.DeleteById(correlationId, id)
//	if result != nil {
//		val, _ := result.(*Dummy)
//		item = val
//	}
//	return item, err
//}
//
//func (c *DummyRefMemoryPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
//	convIds := make([]interface{}, len(ids))
//	for i, v := range ids {
//		convIds[i] = v
//	}
//	return c.IdentifiableMemoryPersistence.DeleteByIds(correlationId, convIds)
//}

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
