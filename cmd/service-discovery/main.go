package main

import (
    "fmt"
    "os"
    "flag"
)

func main() {
    region       := flag.String("region", "us-east-1", "Region of AWS service")
    asgName      := flag.String("asg-name", "", "Autoscalling group name")
    outputField  := flag.String("field", "all", "One of the: all, private-ip, public-ip")
    outputFormat := flag.String("format", "text", "Text format")
    flag.Parse()

    if *asgName == "" {
        flag.Usage()
        os.Exit(255)
    }

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

    var formatter OutputFormatter
    if (*outputFormat == "json") {
        formatter = JSONFormatter{}
    } else {
        formatter = TextFormatter{Field: *outputField}
    }

    res, err2 := formatter.Format(&result)
    if err2 != nil {
        fmt.Println(err1.Error())
        os.Exit(3)
    }

    fmt.Println(res)
}