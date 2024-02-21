// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	apprunner_sdkv2 "github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceDeployment,
			Name:    "Deployment",
		},
		{
			Factory: newResourceIndex,
			Name:    "Default AutoScaling Configuration Version",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAutoScalingConfigurationVersion,
			TypeName: "aws_apprunner_auto_scaling_configuration_version",
			Name:     "AutoScaling Configuration Version",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_apprunner_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceCustomDomainAssociation,
			TypeName: "aws_apprunner_custom_domain_association",
			Name:     "Custom Domain Association",
		},
		{
			Factory:  resourceObservabilityConfiguration,
			TypeName: "aws_apprunner_observability_configuration",
			Name:     "Observability Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceService,
			TypeName: "aws_apprunner_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceVPCConnector,
			TypeName: "aws_apprunner_vpc_connector",
			Name:     "VPC Connector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceVPCIngressConnection,
			TypeName: "aws_apprunner_vpc_ingress_connection",
			Name:     "VPC Ingress Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppRunner
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*apprunner_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return apprunner_sdkv2.NewFromConfig(cfg, func(o *apprunner_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
