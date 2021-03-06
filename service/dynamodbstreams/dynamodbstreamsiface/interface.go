// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package dynamodbstreamsiface provides an interface for the Amazon DynamoDB Streams.
package dynamodbstreamsiface

import (
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

// DynamoDBStreamsAPI is the interface type for dynamodbstreams.DynamoDBStreams.
type DynamoDBStreamsAPI interface {
	DescribeStream(*dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error)

	GetRecords(*dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error)

	GetShardIterator(*dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error)

	ListStreams(*dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error)
}
