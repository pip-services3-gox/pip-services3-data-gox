package test_persistence

import (
	"testing"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
)

type DummyMapPersistenceFixture struct {
	dummy1      map[string]interface{}
	dummy2      map[string]interface{}
	persistence IDummyMapPersistence
}

func NewDummyMapPersistenceFixture(persistence IDummyMapPersistence) *DummyMapPersistenceFixture {
	c := DummyMapPersistenceFixture{}
	c.dummy1 = map[string]interface{}{"Id": "", "Key": "Key 11", "Content": "Content 1"}
	c.dummy2 = map[string]interface{}{"Id": "", "Key": "Key 2", "Content": "Content 2"}
	c.persistence = persistence
	return &c
}

func (c *DummyMapPersistenceFixture) TestCrudOperations(t *testing.T) {
	var dummy1 map[string]interface{}
	var dummy2 map[string]interface{}

	result, err := c.persistence.Create("", c.dummy1)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy1 = result
	assert.NotNil(t, dummy1)
	assert.NotNil(t, dummy1["Id"])
	assert.Equal(t, c.dummy1["Key"], dummy1["Key"])
	assert.Equal(t, c.dummy1["Content"], dummy1["Content"])

	// Create another dummy by set pointer
	result, err = c.persistence.Create("", c.dummy2)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy2 = result
	assert.NotNil(t, dummy2)
	assert.NotNil(t, dummy2["Id"])
	assert.Equal(t, c.dummy2["Key"], dummy2["Key"])
	assert.Equal(t, c.dummy2["Content"], dummy2["Content"])

	page, errp := c.persistence.GetPageByFilter("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	if errp != nil {
		t.Errorf("GetPageByFilter method error %v", err)
	}
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)
	//Testing default sorting by Key field len

	item1 := page.Data[0]
	assert.Equal(t, item1["Key"], dummy2["Key"])
	item2 := page.Data[1]
	assert.Equal(t, item2["Key"], dummy1["Key"])

	// Get count
	count, errc := c.persistence.GetCountByFilter("", cdata.NewEmptyFilterParams())
	assert.Nil(t, errc)
	assert.Equal(t, count, int64(2))

	// Update the dummy
	dummy1["Content"] = "Updated Content 1"
	result, err = c.persistence.Update("", dummy1)
	if err != nil {
		t.Errorf("GetPageByFilter method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1["Id"], result["Id"])
	assert.Equal(t, dummy1["Key"], result["Key"])
	assert.Equal(t, dummy1["Content"], result["Content"])

	// Partially update the dummy
	updateMap := cdata.NewAnyValueMapFromTuples("Content", "Partially Updated Content 1")
	result, err = c.persistence.UpdatePartially("", dummy1["Id"].(string), updateMap)
	if err != nil {
		t.Errorf("UpdatePartially method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1["Id"], result["Id"])
	assert.Equal(t, dummy1["Key"], result["Key"])
	assert.Equal(t, "Partially Updated Content 1", result["Content"])

	// Get the dummy by Id
	result, err = c.persistence.GetOneById("", dummy1["Id"].(string))
	if err != nil {
		t.Errorf("GetOneById method error %v", err)
	}
	// Try to get item
	assert.NotNil(t, result)
	assert.Equal(t, dummy1["Id"], result["Id"])
	assert.Equal(t, dummy1["Key"], result["Key"])
	assert.Equal(t, "Partially Updated Content 1", result["Content"])

	// Delete the dummy
	result, err = c.persistence.DeleteById("", dummy1["Id"].(string))
	if err != nil {
		t.Errorf("DeleteById method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1["Id"], result["Id"])
	assert.Equal(t, dummy1["Key"], result["Key"])
	assert.Equal(t, "Partially Updated Content 1", result["Content"])

	// Get the deleted dummy
	result, err = c.persistence.GetOneById("", dummy1["Id"].(string))
	if err != nil {
		t.Errorf("GetOneById method error %v", err)
	}
	// Try to get item
	assert.Nil(t, result)
}

func (c *DummyMapPersistenceFixture) TestBatchOperations(t *testing.T) {
	var dummy1 map[string]interface{}
	var dummy2 map[string]interface{}

	// Create one dummy
	result, err := c.persistence.Create("", c.dummy1)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy1 = result
	assert.NotNil(t, dummy1)
	assert.NotNil(t, dummy1["Id"])
	assert.Equal(t, c.dummy1["Key"], dummy1["Key"])
	assert.Equal(t, c.dummy1["Content"], dummy1["Content"])

	// Create another dummy
	result, err = c.persistence.Create("", c.dummy2)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy2 = result
	assert.NotNil(t, dummy2)
	assert.NotNil(t, dummy2["Id"])
	assert.Equal(t, c.dummy2["Key"], dummy2["Key"])
	assert.Equal(t, c.dummy2["Content"], dummy2["Content"])

	// Read batch
	items, err := c.persistence.GetListByIds("", []string{dummy1["Id"].(string), dummy2["Id"].(string)})
	if err != nil {
		t.Errorf("GetListByIds method error %v", err)
	}
	//assert.isArray(t,items)
	assert.NotNil(t, items)
	assert.Len(t, items, 2)

	// Delete batch
	err = c.persistence.DeleteByIds("", []string{dummy1["Id"].(string), dummy2["Id"].(string)})
	if err != nil {
		t.Errorf("DeleteByIds method error %v", err)
	}
	assert.Nil(t, err)

	// Read empty batch
	items, err = c.persistence.GetListByIds("", []string{dummy1["Id"].(string), dummy2["Id"].(string)})
	if err != nil {
		t.Errorf("GetListByIds method error %v", err)
	}
	assert.Nil(t, items)
	assert.Len(t, items, 0)

}
