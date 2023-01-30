// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an access policy associated with an existing destination. An
// access policy is an IAM policy document
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) that
// is used to authorize claims to register a subscription filter against a given
// destination. If multiple Amazon Web Services accounts are sending logs to this
// destination, each sender account must be listed separately in the policy. The
// policy does not support specifying * as the Principal or the use of the
// aws:PrincipalOrgId global key.
func (c *Client) PutDestinationPolicy(ctx context.Context, params *PutDestinationPolicyInput, optFns ...func(*Options)) (*PutDestinationPolicyOutput, error) {
	if params == nil {
		params = &PutDestinationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDestinationPolicy", params, optFns, c.addOperationPutDestinationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDestinationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDestinationPolicyInput struct {

	// An IAM policy document that authorizes cross-account users to deliver their log
	// events to the associated destination. This can be up to 5120 bytes.
	//
	// This member is required.
	AccessPolicy *string

	// A name for an existing destination.
	//
	// This member is required.
	DestinationName *string

	// Specify true if you are updating an existing destination policy to grant
	// permission to an organization ID instead of granting permission to individual
	// AWS accounts. Before you update a destination policy this way, you must first
	// update the subscription filters in the accounts that send logs to this
	// destination. If you do not, the subscription filters might stop working. By
	// specifying true for forceUpdate, you are affirming that you have already updated
	// the subscription filters. For more information, see  Updating an existing
	// cross-account subscription
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Cross-Account-Log_Subscription-Update.html)
	// If you omit this parameter, the default of false is used.
	ForceUpdate *bool

	noSmithyDocumentSerde
}

type PutDestinationPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDestinationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutDestinationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDestinationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutDestinationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDestinationPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutDestinationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutDestinationPolicy",
	}
}