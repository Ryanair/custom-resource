package model

type ResourceProperties interface {
	Validate() error
	Create() (string, error)
	Update() (string, error)
	Delete() error
	GetInstance() interface{}
}
