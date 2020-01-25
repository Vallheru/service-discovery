package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetServiceASG(t *testing.T) {
	var input 	ServiceDiscoveryInput
	var service ServiceDiscovery

	input = ServiceDiscoveryInput{ AsgName: "test" }
	service, err := GetService(&input)

	_, ok := service.(ServiceDiscoveryASG)
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestGetServiceError(t *testing.T) {
	var input 	ServiceDiscoveryInput
	var service ServiceDiscovery

	input = ServiceDiscoveryInput{}
	service, err := GetService(&input)

	assert.Nil(t, service)
	assert.Error(t, err)
}