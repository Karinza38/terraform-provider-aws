// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*route53.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return route53.NewFromConfig(cfg,
		route53.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *route53.Options) {
			// Always override the service region
			switch config["partition"].(string) {
			case endpoints.AwsPartitionID:
				// https://docs.aws.amazon.com/general/latest/gr/r53.html Setting default to us-east-1.
				if cfg.Region != endpoints.UsEast1RegionID {
					tflog.Info(ctx, "overriding region", map[string]any{
						"original_region": cfg.Region,
						"override_region": endpoints.UsEast1RegionID,
					})
				}
				o.Region = endpoints.UsEast1RegionID
			case endpoints.AwsCnPartitionID:
				// The AWS Go SDK is missing endpoint information for Route 53 in the AWS China partition.
				// This can likely be removed in the future.
				if aws.ToString(o.BaseEndpoint) == "" {
					o.BaseEndpoint = aws.String("https://api.route53.cn")
				}
				o.Region = endpoints.CnNorthwest1RegionID
			case endpoints.AwsUsGovPartitionID:
				if cfg.Region != endpoints.UsGovWest1RegionID {
					tflog.Info(ctx, "overriding region", map[string]any{
						"original_region": cfg.Region,
						"override_region": endpoints.UsGovWest1RegionID,
					})
				}
				o.Region = endpoints.UsGovWest1RegionID
			}
		},
	), nil
}
