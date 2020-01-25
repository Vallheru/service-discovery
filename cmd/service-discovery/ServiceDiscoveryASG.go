package main

import (
    "errors"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/autoscaling"
    "github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// ServiceDiscoveryASG returns all instances in the given autoscaling group
type ServiceDiscoveryASG struct {
    ServiceDiscovery
    AsgClient     autoscalingiface.AutoScalingAPI
    EC2Client     ec2iface.EC2API
    Input         *ServiceDiscoveryInput
}

// Discover returns all instances for autoscaling group
func (svc ServiceDiscoveryASG) Discover() ([]DiscoveredInstance, error) {
    input := &autoscaling.DescribeAutoScalingGroupsInput{
        AutoScalingGroupNames: []*string{
            aws.String(svc.Input.AsgName),
        },
    }

    result, err := svc.AsgClient.DescribeAutoScalingGroups(input)

    if err != nil {
        return nil, err
    }

    if len(result.AutoScalingGroups) < 1 {
        return nil, errors.New("Not found any ASG")
    }

    ids := []*string{}
    
    for _, asg := range result.AutoScalingGroups {
        for _, instance := range asg.Instances {
            ids = append(ids, instance.InstanceId)
        }
    }
    
    return svc.getInstancesIPs(ids)
}

func (svc ServiceDiscoveryASG) getInstancesIPs(instancesIDs []*string) ([]DiscoveredInstance, error) {
    input := &ec2.DescribeInstancesInput{
        InstanceIds: instancesIDs,
    }
    
    result, err := svc.EC2Client.DescribeInstances(input)
    if err != nil {
        return nil, err
    }

    output := []DiscoveredInstance{}
    for _, instanceGroup := range result.Reservations {
        for _, instance := range instanceGroup.Instances {
            output = append(output, DiscoveredInstance{
                ID:        *instance.InstanceId,   
                PublicIP:  *instance.PublicIpAddress,
                PrivateIP: *instance.PrivateIpAddress,
            })
        }
    }
    
    return output, nil
}