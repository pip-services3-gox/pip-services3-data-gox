package test_persistence

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyMemoryPersistence(t *testing.T) {
	persistence := NewDummyMemoryPersistence()
	persistence.Configure(context.Background(), cconf.NewEmptyConfigParams())

	fixture := NewDummyPersistenceFixture(persistence)

	t.Run("DummyMemoryPersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyMemoryPersistence:Batch", fixture.TestBatchOperations)

}
