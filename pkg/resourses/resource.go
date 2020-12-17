package resourses

import "github.com/Kaiser925/devctl/pkg/common"

// Resource represents resource can be operated.
type Resource interface {
	Kind() string
	Create() error
	Delete() error
}

type ResourceConstructor func(config *common.ResourceConfig) Resource
