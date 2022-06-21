package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IDummyInterfacablePersistence interface {
	GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[DummyInterfacable], err error)
	GetListByIds(ctx context.Context, correlationId string, ids []string) (items []DummyInterfacable, err error)
	GetOneById(ctx context.Context, correlationId string, id string) (item DummyInterfacable, err error)
	Create(ctx context.Context, correlationId string, item DummyInterfacable) (result DummyInterfacable, err error)
	Update(ctx context.Context, correlationId string, item DummyInterfacable) (result DummyInterfacable, err error)
	UpdatePartially(ctx context.Context, correlationId string, id string, data cdata.AnyValueMap) (item DummyInterfacable, err error)
	DeleteById(ctx context.Context, correlationId string, id string) (item DummyInterfacable, err error)
	DeleteByIds(ctx context.Context, correlationId string, ids []string) (err error)
	GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error)
}
