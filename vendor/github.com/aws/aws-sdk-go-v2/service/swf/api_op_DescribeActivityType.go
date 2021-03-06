// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeActivityTypeInput struct {
	_ struct{} `type:"structure"`

	// The activity type to get information about. Activity types are identified
	// by the name and version that were supplied when the activity was registered.
	//
	// ActivityType is a required field
	ActivityType *ActivityType `locationName:"activityType" type:"structure" required:"true"`

	// The name of the domain in which the activity type is registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeActivityTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeActivityTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeActivityTypeInput"}

	if s.ActivityType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivityType"))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}
	if s.ActivityType != nil {
		if err := s.ActivityType.Validate(); err != nil {
			invalidParams.AddNested("ActivityType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Detailed information about an activity type.
type DescribeActivityTypeOutput struct {
	_ struct{} `type:"structure"`

	// The configuration settings registered with the activity type.
	//
	// Configuration is a required field
	Configuration *ActivityTypeConfiguration `locationName:"configuration" type:"structure" required:"true"`

	// General information about the activity type.
	//
	// The status of activity type (returned in the ActivityTypeInfo structure)
	// can be one of the following.
	//
	//    * REGISTERED – The type is registered and available. Workers supporting
	//    this type should be running.
	//
	//    * DEPRECATED – The type was deprecated using DeprecateActivityType,
	//    but is still in use. You should keep workers supporting this type running.
	//    You cannot create new tasks of this type.
	//
	// TypeInfo is a required field
	TypeInfo *ActivityTypeInfo `locationName:"typeInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeActivityTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeActivityType = "DescribeActivityType"

// DescribeActivityTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns information about the specified activity type. This includes configuration
// settings provided when the type was registered and other general information
// about the type.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. activityType.name: String constraint. The key is
//    swf:activityType.name. activityType.version: String constraint. The key
//    is swf:activityType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using DescribeActivityTypeRequest.
//    req := client.DescribeActivityTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeActivityTypeRequest(input *DescribeActivityTypeInput) DescribeActivityTypeRequest {
	op := &aws.Operation{
		Name:       opDescribeActivityType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeActivityTypeInput{}
	}

	req := c.newRequest(op, input, &DescribeActivityTypeOutput{})
	return DescribeActivityTypeRequest{Request: req, Input: input, Copy: c.DescribeActivityTypeRequest}
}

// DescribeActivityTypeRequest is the request type for the
// DescribeActivityType API operation.
type DescribeActivityTypeRequest struct {
	*aws.Request
	Input *DescribeActivityTypeInput
	Copy  func(*DescribeActivityTypeInput) DescribeActivityTypeRequest
}

// Send marshals and sends the DescribeActivityType API request.
func (r DescribeActivityTypeRequest) Send(ctx context.Context) (*DescribeActivityTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeActivityTypeResponse{
		DescribeActivityTypeOutput: r.Request.Data.(*DescribeActivityTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeActivityTypeResponse is the response type for the
// DescribeActivityType API operation.
type DescribeActivityTypeResponse struct {
	*DescribeActivityTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeActivityType request.
func (r *DescribeActivityTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
