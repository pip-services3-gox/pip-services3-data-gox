package test_persistence

import (
	"context"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type DummyMapFilePersistence struct {
	DummyMapMemoryPersistence
	persister *cpersist.JsonFilePersister[DummyMap]
}

func NewDummyMapFilePersistence(path string) *DummyMapFilePersistence {
	c := &DummyMapFilePersistence{
		DummyMapMemoryPersistence: *NewDummyMapMemoryPersistence(),
	}

	persister := cpersist.NewJsonFilePersister[DummyMap](path)
	c.persister = persister
	c.IdentifiableMemoryPersistence.Loader = persister
	c.IdentifiableMemoryPersistence.Saver = persister

	return c
}

func (c *DummyMapFilePersistence) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.DummyMapMemoryPersistence.Configure(ctx, config)
	c.persister.Configure(ctx, config)
}
