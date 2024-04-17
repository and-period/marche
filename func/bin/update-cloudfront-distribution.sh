#!/bin/bash

# Update the Lambda@Edge function in the CloudFront distribution
# - ORIGIN_RESPONSE_ARN: オリジンレスポンス用のLambda@Edge関数のARN
# - CLOUDFRONT_DISTRIBUTION_ID: CloudFrontディストリビューションID

# Get CloudFront distribution configuration
CLOUDFRONT_ETAG=$(aws cloudfront get-distribution-config --id ${CLOUDFRONT_DISTRIBUTION_ID} | jq -r '.ETag')

# Edit the CloudFront distribution configuration to include the new Lambda@Edge function
aws cloudfront get-distribution-config --id ${CLOUDFRONT_DISTRIBUTION_ID} | \
  jq '.DistributionConfig' | \
  jq "(.DefaultCacheBehavior.LambdaFunctionAssociations.Items[] | select(.EventType == \"origin-response\") | .LambdaFunctionARN) |= \"${ORIGIN_RESPONSE_ARN}\"" \
  > ./config.json

ls -la

# Update the CloudFront distribution with the new Lambda@Edge function
aws cloudfront update-distribution --id ${CLOUDFRONT_DISTRIBUTION_ID} --if-match ${CLOUDFRONT_ETAG} --distribution-config file://config.json
