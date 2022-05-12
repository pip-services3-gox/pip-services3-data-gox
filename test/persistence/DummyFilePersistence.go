package test_persistence

import (
	"context"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type DummyFilePersistence struct {
	DummyMemoryPersistence
	persister *cpersist.JsonFilePersister[Dummy]
}

func NewDummyFilePersistence(path string) *DummyFilePersistence {
	c := &DummyFilePersistence{
		DummyMemoryPersistence: *NewDummyMemoryPersistence(),
	}
	persister := cpersist.NewJsonFilePersister[Dummy](path)
	c.persister = persister
	c.Loader = persister
	c.Saver = persister
	return c
}

func (c *DummyFilePersistence) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.DummyMemoryPersistence.Configure(ctx, config)
	c.persister.Configure(ctx, config)
}
