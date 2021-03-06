// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpnTunnelOptionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The tunnel options to modify.
	//
	// TunnelOptions is a required field
	TunnelOptions *ModifyVpnTunnelOptionsSpecification `type:"structure" required:"true"`

	// The ID of the AWS Site-to-Site VPN connection.
	//
	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`

	// The external IP address of the VPN tunnel.
	//
	// VpnTunnelOutsideIpAddress is a required field
	VpnTunnelOutsideIpAddress *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpnTunnelOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpnTunnelOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpnTunnelOptionsInput"}

	if s.TunnelOptions == nil {
		invalidParams.Add(aws.NewErrParamRequired("TunnelOptions"))
	}

	if s.VpnConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnConnectionId"))
	}

	if s.VpnTunnelOutsideIpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnTunnelOutsideIpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpnTunnelOptionsOutput struct {
	_ struct{} `type:"structure"`

	// Describes a VPN connection.
	VpnConnection *VpnConnection `locationName:"vpnConnection" type:"structure"`
}

// String returns the string representation
func (s ModifyVpnTunnelOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpnTunnelOptions = "ModifyVpnTunnelOptions"

// ModifyVpnTunnelOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the options for a VPN tunnel in an AWS Site-to-Site VPN connection.
// You can modify multiple options for a tunnel in a single request, but you
// can only modify one tunnel at a time. For more information, see Site-to-Site
// VPN Tunnel Options for Your Site-to-Site VPN Connection (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using ModifyVpnTunnelOptionsRequest.
//    req := client.ModifyVpnTunnelOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpnTunnelOptions
func (c *Client) ModifyVpnTunnelOptionsRequest(input *ModifyVpnTunnelOptionsInput) ModifyVpnTunnelOptionsRequest {
	op := &aws.Operation{
		Name:       opModifyVpnTunnelOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpnTunnelOptionsInput{}
	}

	req := c.newRequest(op, input, &ModifyVpnTunnelOptionsOutput{})
	return ModifyVpnTunnelOptionsRequest{Request: req, Input: input, Copy: c.ModifyVpnTunnelOptionsRequest}
}

// ModifyVpnTunnelOptionsRequest is the request type for the
// ModifyVpnTunnelOptions API operation.
type ModifyVpnTunnelOptionsRequest struct {
	*aws.Request
	Input *ModifyVpnTunnelOptionsInput
	Copy  func(*ModifyVpnTunnelOptionsInput) ModifyVpnTunnelOptionsRequest
}

// Send marshals and sends the ModifyVpnTunnelOptions API request.
func (r ModifyVpnTunnelOptionsRequest) Send(ctx context.Context) (*ModifyVpnTunnelOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpnTunnelOptionsResponse{
		ModifyVpnTunnelOptionsOutput: r.Request.Data.(*ModifyVpnTunnelOptionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpnTunnelOptionsResponse is the response type for the
// ModifyVpnTunnelOptions API operation.
type ModifyVpnTunnelOptionsResponse struct {
	*ModifyVpnTunnelOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpnTunnelOptions request.
func (r *ModifyVpnTunnelOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
