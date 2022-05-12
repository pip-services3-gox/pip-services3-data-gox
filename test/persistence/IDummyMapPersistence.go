package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IDummyMapPersistence interface {
	GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[DummyMap], err error)
	GetListByIds(ctx context.Context, correlationId string, ids []string) (items []DummyMap, err error)
	GetOneById(ctx context.Context, correlationId string, id string) (item DummyMap, err error)
	Create(ctx context.Context, correlationId string, item DummyMap) (result DummyMap, err error)
	Update(ctx context.Context, correlationId string, item DummyMap) (result DummyMap, err error)
	UpdatePartially(ctx context.Context, correlationId string, id string, data cdata.AnyValueMap) (item DummyMap, err error)
	DeleteById(ctx context.Context, correlationId string, id string) (item DummyMap, err error)
	DeleteByIds(ctx context.Context, correlationId string, ids []string) (err error)
	GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error)
}
