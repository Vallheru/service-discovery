package main

import (
	"bytes"
	"fmt"
	"encoding/json"
)

// OutputFormatter ...
type OutputFormatter interface {
	Format(out *[]DiscoveredInstance) (string, error)
}

// TextFormatter ...
type TextFormatter struct {
	OutputFormatter
	Field			string
}

// JSONFormatter ...
type JSONFormatter struct {
	OutputFormatter
}

// Format ...
func (fmtO TextFormatter) Format(instances *[]DiscoveredInstance) (string, error) {
	var out bytes.Buffer

	for _, item := range *instances {
		switch fmtO.Field{
		case "private-ip":
			out.WriteString(fmt.Sprintf("%s\n", item.PrivateIP))
		case "public-ip":
			out.WriteString(fmt.Sprintf("%s\n", item.PublicIP))
		default:
			out.WriteString(fmt.Sprintf("%s\t%s\t%s\n", item.ID, item.PrivateIP, item.PublicIP))
		}
	}

	return out.String(), nil
}



// Format ...
func (fmtO JSONFormatter) Format(instances *[]DiscoveredInstance) (string, error) {
	res, err := json.Marshal(instances)
	if err != nil {
		return "", nil
	}

	return string(res), nil
}

