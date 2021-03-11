package test_persistence

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

//  extends DummyMemoryPersistence
type DummyRefFilePersistence struct {
	DummyRefMemoryPersistence
	persister *cpersist.JsonFilePersister
}

func NewDummyRefFilePersistence(path string) *DummyRefFilePersistence {
	c := &DummyRefFilePersistence{
		DummyRefMemoryPersistence: *NewDummyRefMemoryPersistence(),
	}
	persister := cpersist.NewJsonFilePersister(c.Prototype, path)
	c.persister = persister
	c.Loader = persister
	c.Saver = persister
	return c
}

func (c *DummyRefFilePersistence) Configure(config *cconf.ConfigParams) {
	c.DummyRefMemoryPersistence.Configure(config)
	c.persister.Configure(config)
}
