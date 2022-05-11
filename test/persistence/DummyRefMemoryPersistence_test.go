package test_persistence

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyRefMemoryPersistence(t *testing.T) {
	persister := NewDummyRefMemoryPersistence()
	persister.Configure(context.Background(), cconf.NewEmptyConfigParams())

	fixture := NewDummyRefPersistenceFixture(persister)

	t.Run("DummyRefMemoryPersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyRefMemoryPersistence:Batch", fixture.TestBatchOperations)

}
