package cloudsearch

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
)

// CloudSearch is a client for Amazon CloudSearch.
type CloudSearch struct {
	*aws.Service
}

type CloudSearchConfig struct {
	*aws.Config
}

// New returns a new CloudSearch client.
func New(config *CloudSearchConfig) *CloudSearch {
	if config == nil {
		config = &CloudSearchConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "cloudsearch",
		APIVersion:  "2011-02-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)

	return &CloudSearch{service}
}