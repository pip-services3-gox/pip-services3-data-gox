package test_persistence

import (
	"context"
	"os"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

func TestDummyFilePersistence(t *testing.T) {
	filename := "../../data/dummies.json"

	//cleaning file before testing
	f, err := os.Create(filename)
	if err != nil {
		t.Error("Can't clean file: ", filename)
	}
	f.Close()

	persistence := NewDummyFilePersistence(filename)
	persistence.Configure(context.Background(), cconf.NewEmptyConfigParams())

	defer persistence.Close(context.Background(), "")

	fixture := NewDummyPersistenceFixture(persistence)
	persistence.Open(context.Background(), "")

	t.Run("DummyFilePersistence:CRUD", fixture.TestCrudOperations)
	t.Run("DummyFilePersistence:Batch", fixture.TestBatchOperations)

}
