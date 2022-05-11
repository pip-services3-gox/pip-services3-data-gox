package test_persistence

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyMapMemoryPersistence(t *testing.T) {
	persister := NewDummyMapMemoryPersistence()
	persister.Configure(context.Background(), cconf.NewEmptyConfigParams())

	fixture := NewDummyMapPersistenceFixture(persister)

	t.Run("DummyMapMemoryPersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyMapMemoryPersistence:Batch", fixture.TestBatchOperations)

}
