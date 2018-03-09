package customresource

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudformationevt"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/Ryanair/custom-resource/model"
	"strings"
	"log"
	"encoding/json"
	"errors"
	"bytes"
	"strconv"
	"io/ioutil"
	"net/http"
)



func Execute(evt *cloudformationevt.Event, properties model.ResourceProperties) (interface{}, error) {

	log.Printf("INPUT %v", evt)
	properties, err := unmarshal(evt, properties)
	physicalResourceId := strings.Join([]string{evt.ServiceToken, evt.RequestID}, "/")
	if err != nil {
		return sendResponse(evt, physicalResourceId, err)
	}

	if evt.RequestType == "Create" {
		err := properties.Validate()
		if err != nil {
			return sendResponse(evt, physicalResourceId, err)
		}
		physicalResourceId, err = properties.Create()
		return sendResponse(evt, physicalResourceId, err)
	} else if evt.RequestType == "Update" {
		err := properties.Validate()
		if err != nil {
			return sendResponse(evt, physicalResourceId, err)
		}
		physicalResourceId, err = properties.Update()
		return sendResponse(evt, physicalResourceId, err)
	} else if evt.RequestType == "Delete" {
		err := properties.Delete()
		return sendResponse(evt, physicalResourceId, err)
	} else {
		err := errors.New("unknown request type")
		return sendResponse(evt, physicalResourceId, err)
	}
}

func unmarshal(evt *cloudformationevt.Event, properties model.ResourceProperties) (model.ResourceProperties, error) {
	instance := properties.GetInstance()
	err := json.Unmarshal(evt.ResourceProperties, &instance)
	properties = instance.(model.ResourceProperties)
	log.Printf("resource properties: %v", awsutil.Prettify(properties))
	return properties, err
}

func sendResponse(evt *cloudformationevt.Event, physicalResourceId string, err error) (string, error) {

	status := model.StatusSuccess
	var reason string

	if err != nil {
		status = model.StatusFailed
		reason = err.Error()
	}

	requestBody := model.CustomResourceResponse{
		Status:             status,
		Reason:             reason,
		StackId:            evt.StackID,
		LogicalResourceId:  evt.LogicalResourceID,
		PhysicalResourceId: physicalResourceId,
		RequestId:          evt.RequestID,
	}
	log.Printf("execution result: %v", requestBody)
	requestBytes, error := json.Marshal(requestBody)
	if error != nil {
		panic(error)
	}

	request, error := http.NewRequest("PUT", evt.ResponseURL, bytes.NewBuffer(requestBytes))
	if error != nil {
		panic(error)
	}
	request.Header.Set("Content-Length", strconv.Itoa(len(requestBytes)))
	request.Header.Set("Content-Type", "")

	client := &http.Client{}
	resp, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	log.Printf("S3 PUT response %v", resp)
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}
	log.Println("S3 PUT response Body:", string(body))

	defer resp.Body.Close()
	return physicalResourceId, err
}
