// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_elb",
			Name:     "Classic Load Balancer",
		},
		{
			Factory:  dataSourceHostedZoneID,
			TypeName: "aws_elb_hosted_zone_id",
			Name:     "Hosted Zone ID",
		},
		{
			Factory:  dataSourceServiceAccount,
			TypeName: "aws_elb_service_account",
			Name:     "Service Account",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAppCookieStickinessPolicy,
			TypeName: "aws_app_cookie_stickiness_policy",
			Name:     "App Cookie Stickiness Policy",
		},
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_elb",
			Name:     "Classic Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceAttachment,
			TypeName: "aws_elb_attachment",
			Name:     "Attachment",
		},
		{
			Factory:  resourceCookieStickinessPolicy,
			TypeName: "aws_lb_cookie_stickiness_policy",
			Name:     "LB Cookie Stickiness Policy",
		},
		{
			Factory:  resourceSSLNegotiationPolicy,
			TypeName: "aws_lb_ssl_negotiation_policy",
			Name:     "SSL Negotiation Policy",
		},
		{
			Factory:  resourceBackendServerPolicy,
			TypeName: "aws_load_balancer_backend_server_policy",
			Name:     "Backend Server Policy",
		},
		{
			Factory:  resourceListenerPolicy,
			TypeName: "aws_load_balancer_listener_policy",
			Name:     "Listener Policy",
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_load_balancer_policy",
			Name:     "Load Balancer Policy",
		},
		{
			Factory:  resourceProxyProtocolPolicy,
			TypeName: "aws_proxy_protocol_policy",
			Name:     "Proxy Protocol Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELB
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticloadbalancing.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return elasticloadbalancing.NewFromConfig(cfg,
		elasticloadbalancing.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
