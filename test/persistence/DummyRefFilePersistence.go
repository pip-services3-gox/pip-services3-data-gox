package test_persistence

import (
	"context"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

//  extends DummyMemoryPersistence
type DummyRefFilePersistence struct {
	DummyRefMemoryPersistence
	persister *cpersist.JsonFilePersister[*DummyRef]
}

func NewDummyRefFilePersistence(path string) *DummyRefFilePersistence {
	c := &DummyRefFilePersistence{
		DummyRefMemoryPersistence: *NewDummyRefMemoryPersistence(),
	}
	persister := cpersist.NewJsonFilePersister[*DummyRef](path)
	c.persister = persister
	c.IdentifiableMemoryPersistence.Loader = persister
	c.IdentifiableMemoryPersistence.Saver = persister
	return c
}

func (c *DummyRefFilePersistence) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.DummyRefMemoryPersistence.Configure(ctx, config)
	c.persister.Configure(ctx, config)
}
