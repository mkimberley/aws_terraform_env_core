package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// An example of how to test the Terraform module in examples/terraform-aws-network-example using Terratest.
func TestTerraformAwsNetworkExample(t *testing.T) {
	t.Parallel()

	// Give the VPC and the subnets correct CIDRs
	environment := "test"
	vpcName := "Test VPC"
	vpcEnableDnsHostnames := true
	vpcEnableDnsSupport := true
	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	//awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	awsRegion := "eu-west-2"
	primarySubnetName := "Test Primary"
	secondarySubnetName := "Test Secondary"
	backendSubnetName := "Test Backend"
	vpcInstanceTenancy := "default"
	vpcPrimaryPublicIpOnLaunch := true
	vpcSecondaryPublicIpOnLaunch := false
	vpcCidr := "10.40.0.0/16"
	subAzs := `["eu-west-2a","eu-west-2b","eu-west-2c"]`
	publicSubnets := `["10.40.1.0/24","10.40.2.0/24","10.40.3.0/24"]`
	privateSubnets := `["10.40.4.0/24","10.40.5.0/24","10.40.6.0/24"]`
	backendSubnets := `["10.40.7.0/24","10.40.8.0/24","10.40.9.0/24"]`

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"environment":                       environment,
			"vpc_name":                          vpcName,
			"vpc_enable_dns_hostnames":          vpcEnableDnsHostnames,
			"vpc_enable_dns_support":            vpcEnableDnsSupport,
			"primary_subnet_name":               primarySubnetName,
			"secondary_subnet_name":             secondarySubnetName,
			"backend_subnet_name":               backendSubnetName,
			"vpc_instance_tenancy":              vpcInstanceTenancy,
			"vpc_primary_public_ip_on_launch":   vpcPrimaryPublicIpOnLaunch,
			"vpc_secondary_public_ip_on_launch": vpcSecondaryPublicIpOnLaunch,
			"vpc_cidr":                          vpcCidr,
			"sub_azs":                           subAzs,
			"public_subnets":                    publicSubnets,
			"private_subnets":                   privateSubnets,
			"backend_subnets":                   backendSubnets,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	publicSubnetId := terraform.Output(t, terraformOptions, "public_subnets")
	privateSubnetId := terraform.Output(t, terraformOptions, "private_subnets")
	vpcId := terraform.Output(t, terraformOptions, "core_vpc_id")

	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

	require.Equal(t, 9, len(subnets))
	// Verify if the network that is supposed to be public is really public
	assert.True(t, aws.IsPublicSubnet(t, publicSubnetId, awsRegion))
	// Verify if the network that is supposed to be private is really private
	assert.False(t, aws.IsPublicSubnet(t, privateSubnetId, awsRegion))
}
