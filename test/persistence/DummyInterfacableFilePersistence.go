package test_persistence

import (
	"context"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type DummyInterfacableFilePersistence struct {
	DummyInterfacableMemoryPersistence
	persister *cpersist.JsonFilePersister[DummyInterfacable]
}

func NewDummyInterfacableFilePersistence(path string) *DummyInterfacableFilePersistence {
	c := &DummyInterfacableFilePersistence{
		DummyInterfacableMemoryPersistence: *NewDummyInterfacableMemoryPersistence(),
	}
	persister := cpersist.NewJsonFilePersister[DummyInterfacable](path)
	c.persister = persister
	c.Loader = persister
	c.Saver = persister
	return c
}

func (c *DummyInterfacableFilePersistence) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.DummyInterfacableMemoryPersistence.Configure(ctx, config)
	c.persister.Configure(ctx, config)
}
