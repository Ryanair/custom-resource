package model

import "github.com/aws/aws-sdk-go/aws/awsutil"

type CustomResourceResponse struct {
	Status             string `type:"string" enum:"Status"`
	Reason             string
	PhysicalResourceId string
	StackId            string
	RequestId          string
	LogicalResourceId  string
	Data               map[string]string
}

func (s CustomResourceResponse) String() string {
	return awsutil.Prettify(s)
}

const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)
