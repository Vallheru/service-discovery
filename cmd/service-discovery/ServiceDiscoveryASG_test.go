package main

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)


type mockAsgClientCorrectResult struct {
    autoscalingiface.AutoScalingAPI
}

func (t mockAsgClientCorrectResult) DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    res := &autoscaling.DescribeAutoScalingGroupsOutput{
        AutoScalingGroups: []*autoscaling.Group{
            { Instances: []*autoscaling.Instance{ { InstanceId: aws.String("i-1") } } },
            { Instances: []*autoscaling.Instance{ { InstanceId: aws.String("i-2") } } },
        },
    }

    return res, nil
}

type mockEC2ClientCorrectResult struct {
    ec2iface.EC2API
}

func (t mockEC2ClientCorrectResult) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
    res := &ec2.DescribeInstancesOutput{
        Reservations: []*ec2.Reservation{
            {
                Instances: []*ec2.Instance{
                    {
                        InstanceId: 	  aws.String("i-1"),
                        PublicIpAddress:  aws.String("143.143.143.143"),
                        PrivateIpAddress: aws.String("10.0.0.1"),
                    },
                },
            },
            {
                Instances: []*ec2.Instance{
                    {
                        InstanceId: 	  aws.String("i-2"),
                        PublicIpAddress:  aws.String("143.143.143.144"),
                        PrivateIpAddress: aws.String("10.0.0.2"),
                    },
                },
            },
        },
    }

    return res, nil
}

type mockAsgClientEmptyAsg struct {
    autoscalingiface.AutoScalingAPI
}

func (t mockAsgClientEmptyAsg) DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    res := &autoscaling.DescribeAutoScalingGroupsOutput{
        AutoScalingGroups: []*autoscaling.Group{},
    }

    return res, nil
}

type mockEC2ClientEmptyInstances struct {
    ec2iface.EC2API
}

func (t mockEC2ClientEmptyInstances) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
    res := &ec2.DescribeInstancesOutput{
        Reservations: []*ec2.Reservation{},
    }

    return res, nil
}


type mockEC2ClientError struct {
    ec2iface.EC2API
}

func (t mockEC2ClientError) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
    return nil, errors.New("Example error from EC2")
}


type mockAsgClientError struct {
    autoscalingiface.AutoScalingAPI
}

func (t mockAsgClientError) DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    return nil, errors.New("Example error from ASG")
}

func TestASGDiscoverCorrectResult(t *testing.T) {
	table := []struct {
		AsgClient      autoscalingiface.AutoScalingAPI
		EC2Client      ec2iface.EC2API
		expectedError  error
		expectedOutput []DiscoveredInstance
	}{
		{ 
			&mockAsgClientCorrectResult{},
			&mockEC2ClientCorrectResult{},
			nil,
			[]DiscoveredInstance{
				{
					ID:        "i-1",
					PublicIP:  "143.143.143.143",
					PrivateIP: "10.0.0.1",
				},
				{
					ID:        "i-2",
					PublicIP:  "143.143.143.144",
					PrivateIP: "10.0.0.2",
				},
			},
		},
		{
			&mockAsgClientError{},
			&mockEC2ClientCorrectResult{},
			errors.New("Example error from ASG"),
			nil,
		},
		{
			&mockAsgClientCorrectResult{},
			&mockEC2ClientError{},
			errors.New("Example error from EC2"),
			nil,
		},
		{
			&mockAsgClientEmptyAsg{},
			&mockEC2ClientCorrectResult{},
			errors.New("Not found any ASG"),
			nil,
		},
		{
			&mockAsgClientCorrectResult{},
			&mockEC2ClientEmptyInstances{},
			nil,
			[]DiscoveredInstance{},
		},
	}

	for _, row := range table {
		service := ServiceDiscoveryASG{
			Input:	   &ServiceDiscoveryInput{ AsgName: "Test" },
			AsgClient: row.AsgClient,
			EC2Client: row.EC2Client,
		}
	
		output, err := service.Discover()
		
		assert.Equal(t, row.expectedError, err)
		assert.Equal(t, row.expectedOutput, output)
	}	
}