package testdata

import "errors"

type ExampleProperties struct {
	Property1 string
	Property2 string
}

func (data ExampleProperties) Validate() error {
	if data.Property1 == "" {
		return errors.New("Property1 must be set")
	}
	return nil
}

func (data ExampleProperties) Create() (string, error) {
	return "testPhysicalResourceId1", nil
}

func (data ExampleProperties) Update() (string, error) {
	return "testPhysicalResourceId2", nil
}

func (data ExampleProperties) Delete() error {
	return nil
}

func (data ExampleProperties) GetInstance() interface{} {
	return &ExampleProperties{}
}

type ExampleErrorProperties struct {
	Property1 string
	Property2 string
}

func (data ExampleErrorProperties) Validate() error {
	if data.Property1 == "" {
		return errors.New("Property1 must be set")
	}
	return nil
}

func (data ExampleErrorProperties) Create() (string, error) {
	return "testPhysicalResourceId1", errors.New("Error during create action")
}

func (data ExampleErrorProperties) Update() (string, error) {
	return "testPhysicalResourceId2", errors.New("Error during update action")
}

func (data ExampleErrorProperties) Delete() error {
	return errors.New("Error during delete action")
}

func (data ExampleErrorProperties) GetInstance() interface{} {
	return &ExampleErrorProperties{}
}

