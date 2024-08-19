package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func TestLogStream(t *testing.T, ctx types.TestContext) {
	cloudwatchClient := cloudwatchlogs.NewFromConfig(GetAWSConfig(t))
	streamName := terraform.Output(t, ctx.TerratestTerraformOptions(), "log_stream_name")
	streamArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "log_stream_arn")
	groupArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "log_group_arn")

	output, err := cloudwatchClient.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{
		LogStreamNamePrefix: &streamName,
		LogGroupIdentifier:  &groupArn,
	})
	if err != nil {
		t.Errorf("Unable to get log streams, %v", err)
	}

	t.Run("TestDoesstreamExist", func(t *testing.T) {
		require.Equal(t, 1, len(output.LogStreams), "Expected 1 log stream, got %d", len(output.LogStreams))
	})

	t.Run("TeststreamArn", func(t *testing.T) {
		require.Equal(t, streamArn, *output.LogStreams[0].Arn, "Expected ARN to be %s, got %s", streamArn, *output.LogStreams[0].Arn)
	})
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
