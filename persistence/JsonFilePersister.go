package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	"io/ioutil"
	"os"
)

/*
Persistence component that loads and saves data from/to flat file.

It is used by FilePersistence, but can be useful on its own.

 Configuration parameters

  - path:          path to the file where data is stored

 Example

  persister := NewJsonFilePersister(reflect.TypeOf(MyData{}), "./data/data.json");

  err_sav := persister.Save("123", ["A", "B", "C"])
  if err_sav == nil {
  	items, err_lod := persister.Load("123")
  	if err_lod == nil {
  		fmt.Println(items);// Result: ["A", "B", "C"]
  	}
*/
// implements ILoader, ISaver, IConfigurable
type JsonFilePersister[T any] struct {
	path      string
	convertor convert.IJSONEngine[[]T]
}

const JsonFilePersistenceConfigParamPath = "path"

// Creates a new instance of the persistence.
// Parameters:
//  - path  string
//  (optional) a path to the file where data is stored.
func NewJsonFilePersister[T any](path string) *JsonFilePersister[T] {
	return &JsonFilePersister[T]{
		path:      path,
		convertor: convert.NewDefaultCustomTypeJsonConvertor[[]T](),
	}
}

// Gets the file path where data is stored.
// Returns the file path where data is stored.
func (c *JsonFilePersister[T]) Path() string {
	return c.path
}

// Sets the file path where data is stored.
// Parameters:
//  - value  string
//  the file path where data is stored.
func (c *JsonFilePersister[T]) SetPath(value string) {
	c.path = value
}

// Configures component by passing configuration parameters.
// Parameters:
//  - config  config.ConfigParams
//  parameters to be set.
func (c *JsonFilePersister[T]) Configure(ctx context.Context, config *config.ConfigParams) {
	c.path = config.GetAsStringWithDefault(JsonFilePersistenceConfigParamPath, c.path)
}

// Loads data items from external JSON file.
// Parameters:
//  - correlationId  string
//  transaction id to trace execution through call chain.
// Returns []interface{}, error
// loaded items or error.
func (c *JsonFilePersister[T]) Load(ctx context.Context, correlationId string) ([]T, error) {
	if c.path == "" {
		return nil, errors.NewConfigError("", "NO_PATH", "Data file path is not set")
	}

	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		return nil, err
	}

	jsonStr, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, errors.NewFileError(
			correlationId,
			"READ_FAILED",
			"Failed to read data file: "+c.path).
			WithCause(err)
	}

	if len(jsonStr) == 0 {
		return nil, nil
	}

	if list, err := c.convertor.FromJson(string(jsonStr)); err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

// Saves given data items to external JSON file.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - items []interface[]
//   list of data items to save
//  Retruns error
//  error or nil for success.
func (c *JsonFilePersister[T]) Save(ctx context.Context, correlationId string, items []T) error {
	json, err := c.convertor.ToJson(items)
	if err != nil {
		err := errors.NewInternalError(correlationId, "CAN'T_CONVERT", "Failed convert to JSON")
		return err
	}
	if err := ioutil.WriteFile(c.path, ([]byte)(json), 0777); err != nil {
		return errors.NewFileError(
			correlationId,
			"WRITE_FAILED",
			"Failed to write data file: "+c.path).
			WithCause(err)
	}
	return nil
}
