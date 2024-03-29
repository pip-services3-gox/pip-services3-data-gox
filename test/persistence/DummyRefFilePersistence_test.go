package test_persistence

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyRefFilePersistence(t *testing.T) {
	filename := "../../data/dummies_ref.json"

	//cleaning file before testing
	f, err := os.Create(filename)
	if err != nil {
		t.Error("Can't clean file: ", filename)
	}
	_ = f.Close()

	persistence := NewDummyRefFilePersistence(filename)
	persistence.Configure(context.Background(), cconf.NewEmptyConfigParams())

	defer persistence.Close(context.Background(), "")

	fixture := NewDummyRefPersistenceFixture(persistence)
	_ = persistence.Open(context.Background(), "")

	t.Run("DummyFilePersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyFilePersistence:Batch", fixture.TestBatchOperations)

}
