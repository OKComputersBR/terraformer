// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSourceCredentialsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListSourceCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

type ListSourceCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// A list of SourceCredentialsInfo objects. Each SourceCredentialsInfo object
	// includes the authentication type, token ARN, and type of source provider
	// for one set of credentials.
	SourceCredentialsInfos []SourceCredentialsInfo `locationName:"sourceCredentialsInfos" type:"list"`
}

// String returns the string representation
func (s ListSourceCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSourceCredentials = "ListSourceCredentials"

// ListSourceCredentialsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Returns a list of SourceCredentialsInfo objects.
//
//    // Example sending a request using ListSourceCredentialsRequest.
//    req := client.ListSourceCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/ListSourceCredentials
func (c *Client) ListSourceCredentialsRequest(input *ListSourceCredentialsInput) ListSourceCredentialsRequest {
	op := &aws.Operation{
		Name:       opListSourceCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSourceCredentialsInput{}
	}

	req := c.newRequest(op, input, &ListSourceCredentialsOutput{})
	return ListSourceCredentialsRequest{Request: req, Input: input, Copy: c.ListSourceCredentialsRequest}
}

// ListSourceCredentialsRequest is the request type for the
// ListSourceCredentials API operation.
type ListSourceCredentialsRequest struct {
	*aws.Request
	Input *ListSourceCredentialsInput
	Copy  func(*ListSourceCredentialsInput) ListSourceCredentialsRequest
}

// Send marshals and sends the ListSourceCredentials API request.
func (r ListSourceCredentialsRequest) Send(ctx context.Context) (*ListSourceCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSourceCredentialsResponse{
		ListSourceCredentialsOutput: r.Request.Data.(*ListSourceCredentialsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSourceCredentialsResponse is the response type for the
// ListSourceCredentials API operation.
type ListSourceCredentialsResponse struct {
	*ListSourceCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSourceCredentials request.
func (r *ListSourceCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
