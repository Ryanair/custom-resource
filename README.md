# Cloudformation custom resource template
## Usage example

### Cloudformation resource definition
Define service token (lambda arn) and input properties
```
MyCustomResource:
 Type: Custom::MyCustomResource
 Version: "1.0"
 Properties:
   ServiceToken: <lambda_arn>
   Property1: "test1"
   Property2: "test2"
```

### Handler implementation
Pass struct that implements model.ResourceProperties interface as a second parameter
```
func Handle(evt *cloudformationevt.Event, runtimeCtx *runtime.Context) (interface{}, error) {
	return customresource.Execute(evt, MyResourceProperties{})
}
```

### Resource properties interface implementation
#### Properties
Define properties as in cloudformation

```
type MyResourceProperties struct {
	Property1 string
	Property2 string
}
```

#### Validate()
Validate required fields and other data if necessary

```
func (data MyResourceProperties) Validate() error {
	if r.Property1 == "" {
		return errors.New("Property1 must be set")
	}
	return nil
}
```

#### Create()
* Action on resource create
* It must always return non-null physical resource ID
* If error is not nil it will return as FAILED, otherwise SUCCESS
```
func (data MyResourceProperties) Create() (string, error) {
  //create logic
	return "<physicalResourceId>", nil
}
```

#### Update()
* Action on resource update
* It must always return non-null physical resource ID
* If ID is changed, cloudformation will invoke recreate action
* If error is not nil it will return as FAILED, otherwise SUCCESS
```
func (data MyResourceProperties) Update() (string, error) {
  //update logic
	return "<physicalResourceId>", nil
}
```

#### Delete()
* Action on resource delete
* If error is not nil it will return as FAILED, otherwise SUCCESS
```
func (data MyResourceProperties) Delete() error {

	return nil
}
```

#### GetInstance()
Helper method to allow json unmarshalling
```
func (data MyResourceProperties) GetInstance() interface{} {
	return &MyResourceProperties{}
}```
