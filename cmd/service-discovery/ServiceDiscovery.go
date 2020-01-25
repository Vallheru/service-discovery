package main

import (
    "errors"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/autoscaling"
    "github.com/aws/aws-sdk-go/service/ec2"
)

// DiscoveredInstance is an output structure
type DiscoveredInstance struct {
    ID        string
    PublicIP  string
    PrivateIP string
}

// ServiceDiscoveryInput input structure for service discovery
type ServiceDiscoveryInput struct {
    Region            string
    AsgName           string
}

// ServiceDiscovery is main interface for all services
type ServiceDiscovery interface {
    Discover() ([]DiscoveredInstance, error)
}

// GetService returns class to discover all vms
func GetService(input *ServiceDiscoveryInput) (ServiceDiscovery, error) {
    if input.AsgName != "" {
        sess, _ := session.NewSession( &aws.Config{Region: aws.String(input.Region)} )

        return ServiceDiscoveryASG{
            Input:      input,
            AsgClient:  autoscaling.New(sess),
            EC2Client:  ec2.New(sess),
        }, nil
    }

    return nil, errors.New("Could not build DiscoveryService")
}