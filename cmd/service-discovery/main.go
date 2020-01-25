package main

import (
    "fmt"
    "os"
    "flag"
)

func main() {
    region      := flag.String("region", "us-east-1", "Region of AWS service")
    asgName     := flag.String("asg-name", "", "Autoscalling group name")
    outputValue := flag.String("output-value", "private-ip", "One of the: private-ip, public-ip")
    flag.Parse()

    input := &ServiceDiscoveryInput{
        Region:  *region,
        AsgName: *asgName,
    }

    service, err := GetService(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    result, err1 := service.Discover()
    if err1 != nil {
        fmt.Println(err1.Error())
        os.Exit(2)
    }

    switch *outputValue {
    case "private-ip":
        for _, item := range result {
            fmt.Println(item.PrivateIP)
        }

    default:
        for _, item := range result {
            fmt.Println(item.PublicIP)
        }
    }
    
}