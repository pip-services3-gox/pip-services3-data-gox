package test_persistence

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyInterfacableMemoryPersistence(t *testing.T) {
	persistence := NewDummyInterfacableMemoryPersistence()
	persistence.Configure(context.Background(), cconf.NewEmptyConfigParams())

	fixture := NewDummyInterfacablePersistenceFixture(persistence)

	t.Run("DummyInterfacableMemoryPersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyInterfacableMemoryPersistence:Batch", fixture.TestBatchOperations)

}
