package customresource

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudformationevt"
	"github.com/Ryanair/custom-resource/testdata"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Ryanair/custom-resource/model"
	"encoding/json"
)

func TestCreateExecute(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Create",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusSuccess, response.Status)
	assert.Equal(t, "testPhysicalResourceId1", response.PhysicalResourceId)
}

func TestUpdateExecute(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Update",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusSuccess, response.Status)
	assert.Equal(t, "testPhysicalResourceId2", response.PhysicalResourceId)
}

func TestDeleteExecute(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Delete",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusSuccess, response.Status)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func TestValidationErrorOnCreate(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Create",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, "Property1 must be set", response.Reason)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func TestValidationErrorOnUpdate(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Update",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, "Property1 must be set", response.Reason)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func TestNoValidationErrorOnDelete(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		//Property1 required but not checked on delete
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Delete",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusSuccess, response.Status)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func TestInvalidRequestType(t *testing.T) {
	//given
	properties := testdata.ExampleProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Recreate", //invalid value
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, "unknown request type", response.Reason)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func TestCreateExecuteWithError(t *testing.T) {
	//given
	properties := testdata.ExampleErrorProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Create",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, "testPhysicalResourceId1", response.PhysicalResourceId)
}

func TestUpdateExecuteWithError(t *testing.T) {
	//given
	properties := testdata.ExampleErrorProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Update",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, "testPhysicalResourceId2", response.PhysicalResourceId)
}

func TestDeleteExecuteWithError(t *testing.T) {
	//given
	properties := testdata.ExampleErrorProperties{
		Property1: "test1",
		Property2: "test2",
	}
	bytes, _ := json.Marshal(properties)
	request := cloudformationevt.Event{
		ServiceToken:       "lambdaArn",
		RequestID:          "id-123",
		RequestType:        "Delete",
		LogicalResourceID:  "LogicalId",
		PhysicalResourceID: "PhysicalId",
		ResourceProperties: bytes,
		ResponseURL:        "http://mock",
	}
	//when
	response, error := Execute(&request, properties)
	//then
	assertExpectedHttpError(t, error)
	assert.Equal(t, model.StatusFailed, response.Status)
	assert.Equal(t, request.PhysicalResourceID, response.PhysicalResourceId)
}

func assertExpectedHttpError(t *testing.T, error error) {
	assert.Equal(t, "Put http://mock: dial tcp: lookup mock: no such host", error.Error())
}
