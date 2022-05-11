package test_persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// extends IGetter<Dummy, String>, IWriter<Dummy, String>, IPartialUpdater<Dummy, String> {
type IDummyRefPersistence interface {
	GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[*DummyRef], err error)
	GetListByIds(ctx context.Context, correlationId string, ids []string) (items []*DummyRef, err error)
	GetOneById(ctx context.Context, correlationId string, id string) (item *DummyRef, err error)
	Create(ctx context.Context, correlationId string, item *DummyRef) (result *DummyRef, err error)
	Update(ctx context.Context, correlationId string, item *DummyRef) (result *DummyRef, err error)
	UpdatePartially(ctx context.Context, correlationId string, id string, data cdata.AnyValueMap) (item *DummyRef, err error)
	DeleteById(ctx context.Context, correlationId string, id string) (item *DummyRef, err error)
	DeleteByIds(ctx context.Context, correlationId string, ids []string) (err error)
	GetCountByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams) (count int64, err error)
}
