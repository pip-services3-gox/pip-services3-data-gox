package test_persistence

import (
	"context"
	"testing"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type DummyMapPersistenceFixture struct {
	dummy1      DummyMap
	dummy2      DummyMap
	persistence IDummyMapPersistence
}

func NewDummyMapPersistenceFixture(persistence IDummyMapPersistence) *DummyMapPersistenceFixture {
	c := DummyMapPersistenceFixture{}
	c.dummy1 = DummyMap{
		AnyValueMap: *cdata.NewAnyValueMap(map[string]any{"Id": "", "Key": "Key 11", "Content": "Content 1"}),
	}
	c.dummy2 = DummyMap{
		AnyValueMap: *cdata.NewAnyValueMap(map[string]any{"Id": "", "Key": "Key 2", "Content": "Content 2"}),
	}
	c.persistence = persistence
	return &c
}

func (c *DummyMapPersistenceFixture) TestCrudOperations(t *testing.T) {
	var dummy1 DummyMap
	var dummy2 DummyMap

	result, err := c.persistence.Create(context.Background(), "", c.dummy1)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy1 = result
	assert.NotNil(t, dummy1)
	assert.True(t, dummy1.GetAsString("Id") != "")
	assert.Equal(t, c.dummy1.GetAsString("Key"), dummy1.GetAsString("Key"))
	assert.Equal(t, c.dummy1.GetAsString("Content"), dummy1.GetAsString("Content"))

	// Create another dummy by set pointer
	result, err = c.persistence.Create(context.Background(), "", c.dummy2)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy2 = result
	assert.NotNil(t, dummy2)
	assert.True(t, dummy2.GetAsString("Id") != "")
	assert.Equal(t, c.dummy2.GetAsString("Key"), dummy2.GetAsString("Key"))
	assert.Equal(t, c.dummy2.GetAsString("Content"), dummy2.GetAsString("Content"))

	page, errp := c.persistence.GetPageByFilter(context.Background(), "", *cdata.NewEmptyFilterParams(), *cdata.NewEmptyPagingParams())
	if errp != nil {
		t.Errorf("GetPageByFilter method error %v", err)
	}
	assert.NotNil(t, page)
	_data, ok := page.Data()
	assert.True(t, ok)
	assert.Len(t, _data, 2)
	//Testing default sorting by Key field len

	assert.Equal(t, _data[0].GetAsString("Key"), dummy2.GetAsString("Key"))
	assert.Equal(t, _data[1].GetAsString("Key"), dummy1.GetAsString("Key"))

	// Get count
	count, errc := c.persistence.GetCountByFilter(context.Background(), "", *cdata.NewEmptyFilterParams())
	assert.Nil(t, errc)
	assert.Equal(t, count, int64(2))

	// Update the dummy
	dummy1.Append(map[string]any{"Content": "Updated Content 1"})
	result, err = c.persistence.Update(context.Background(), "", dummy1)
	if err != nil {
		t.Errorf("GetPageByFilter method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1.GetAsString("Id"), result.GetAsString("Id"))
	assert.Equal(t, dummy1.GetAsString("Key"), result.GetAsString("Key"))
	assert.Equal(t, dummy1.GetAsString("Content"), result.GetAsString("Content"))

	// Partially update the dummy
	updateMap := *cdata.NewAnyValueMapFromTuples("Content", "Partially Updated Content 1")
	result, err = c.persistence.UpdatePartially(context.Background(), "", dummy1.GetAsString("Id"), updateMap)
	if err != nil {
		t.Errorf("UpdatePartially method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1.GetAsString("Id"), result.GetAsString("Id"))
	assert.Equal(t, dummy1.GetAsString("Key"), result.GetAsString("Key"))
	assert.Equal(t, "Partially Updated Content 1", result.GetAsString("Content"))

	// Get the dummy by Id
	result, err = c.persistence.GetOneById(context.Background(), "", dummy1.GetAsString("Id"))
	if err != nil {
		t.Errorf("GetOneById method error %v", err)
	}
	// Try to get item
	assert.NotNil(t, result)
	assert.Equal(t, dummy1.GetAsString("Id"), result.GetAsString("Id"))
	assert.Equal(t, dummy1.GetAsString("Key"), result.GetAsString("Key"))
	assert.Equal(t, "Partially Updated Content 1", result.GetAsString("Content"))

	// Delete the dummy
	result, err = c.persistence.DeleteById(context.Background(), "", dummy1.GetAsString("Id"))
	if err != nil {
		t.Errorf("DeleteById method error %v", err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, dummy1.GetAsString("Id"), result.GetAsString("Id"))
	assert.Equal(t, dummy1.GetAsString("Key"), result.GetAsString("Key"))
	assert.Equal(t, "Partially Updated Content 1", result.GetAsString("Content"))

	// Get the deleted dummy
	result, err = c.persistence.GetOneById(context.Background(), "", dummy1.GetAsString("Id"))
	if err != nil {
		t.Errorf("GetOneById method error %v", err)
	}
	// Try to get item
	//assert.Nil(t, result.AnyValueMap)
}

func (c *DummyMapPersistenceFixture) TestBatchOperations(t *testing.T) {
	var dummy1 DummyMap
	var dummy2 DummyMap

	// Create one dummy
	result, err := c.persistence.Create(context.Background(), "", c.dummy1)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy1 = result
	assert.NotNil(t, dummy1)
	assert.True(t, dummy1.GetAsString("Id") != "")
	assert.Equal(t, c.dummy1.GetAsString("Key"), dummy1.GetAsString("Key"))
	assert.Equal(t, c.dummy1.GetAsString("Content"), dummy1.GetAsString("Content"))

	// Create another dummy
	result, err = c.persistence.Create(context.Background(), "", c.dummy2)
	if err != nil {
		t.Errorf("Create method error %v", err)
	}
	dummy2 = result
	assert.NotNil(t, dummy2)
	assert.True(t, dummy2.GetAsString("Id") != "")
	assert.Equal(t, c.dummy2.GetAsString("Key"), dummy2.GetAsString("Key"))
	assert.Equal(t, c.dummy2.GetAsString("Content"), dummy2.GetAsString("Content"))

	// Read batch
	items, err := c.persistence.GetListByIds(context.Background(), "", []string{
		dummy1.GetAsString("Id"), dummy2.GetAsString("Id")})
	if err != nil {
		t.Errorf("GetListByIds method error %v", err)
	}
	//assert.isArray(t,items)
	assert.NotNil(t, items)
	assert.Len(t, items, 2)

	// Delete batch
	err = c.persistence.DeleteByIds(context.Background(), "", []string{
		dummy1.GetAsString("Id"), dummy2.GetAsString("Id")})
	if err != nil {
		t.Errorf("DeleteByIds method error %v", err)
	}
	assert.Nil(t, err)

	// Read empty batch
	items, err = c.persistence.GetListByIds(context.Background(), "", []string{
		dummy1.GetAsString("Id"), dummy2.GetAsString("Id")})
	if err != nil {
		t.Errorf("GetListByIds method error %v", err)
	}
	assert.Len(t, items, 0)

}
