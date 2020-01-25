package main

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTextFormatter(t *testing.T) {
    table := []struct{
        Field            string
        ExpectedResult    string
        Instances         []DiscoveredInstance
    }{
        {
            "all",
            fmt.Sprintf("%s\n%s\n",
                "i-12345678901234567\t10.10.10.10\t3.3.3.3",
                "i-99999999999999999\t10.10.10.11\t3.3.3.4",
            ),
            []DiscoveredInstance{
                {"i-12345678901234567", "3.3.3.3", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", "10.10.10.11"},
            },
        },
        {
            "all",
            "",
            []DiscoveredInstance{},
        },
        {
            "private-ip",
            "",
            []DiscoveredInstance{},
        },
        {
            "private-ip",
            "10.10.10.10\n10.10.10.11\n",
            []DiscoveredInstance{
                {"i-12345678901234567", "3.3.3.3", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", "10.10.10.11"},
            },
        },
        {
            "public-ip",
            "3.3.3.3\n3.3.3.4\n",
            []DiscoveredInstance{
                {"i-12345678901234567", "3.3.3.3", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", "10.10.10.11"},
            },
        },
        {
            "invalid",
            fmt.Sprintf("%s\n%s\n",
                "i-12345678901234567\t10.10.10.10\t3.3.3.3",
                "i-99999999999999999\t10.10.10.11\t3.3.3.4",
            ),
            []DiscoveredInstance{
                {"i-12345678901234567", "3.3.3.3", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", "10.10.10.11"},
            },
        },
    }

    var formatter OutputFormatter
    for _, item := range table {
        formatter = TextFormatter{
            Field: item.Field,
        }
        res, err := formatter.Format(&item.Instances)

        assert.Equal(t, item.ExpectedResult, res)
        assert.Nil(t, err)
    }
}

func TestJSONFormatter(t *testing.T) {
    table := []struct{
        ExpectedResult    string
        Instances         []DiscoveredInstance
    }{
        {
            fmt.Sprintf("[%s,%s]",
                "{\"ID\":\"i-12345678901234567\",\"PublicIP\":\"3.3.3.3\",\"PrivateIP\":\"10.10.10.10\"}",
                "{\"ID\":\"i-99999999999999999\",\"PublicIP\":\"3.3.3.4\",\"PrivateIP\":\"10.10.10.11\"}",
            ),
            []DiscoveredInstance{
                {"i-12345678901234567", "3.3.3.3", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", "10.10.10.11"},
            },
        },
        {
            "[]",
            []DiscoveredInstance{},
        },
        {
            fmt.Sprintf("[%s,%s,%s]",
                "{\"ID\":\"i-12345678901234567\",\"PublicIP\":\"\",\"PrivateIP\":\"10.10.10.10\"}",
                "{\"ID\":\"i-99999999999999999\",\"PublicIP\":\"3.3.3.4\",\"PrivateIP\":\"\"}",
                "{\"ID\":\"\",\"PublicIP\":\"3.3.3.5\",\"PrivateIP\":\"10.10.10.12\"}",
            ),
            []DiscoveredInstance{
                {"i-12345678901234567", "", "10.10.10.10"},
                {"i-99999999999999999", "3.3.3.4", ""},
                {"", "3.3.3.5", "10.10.10.12"},
            },
        },
    }

    var formatter OutputFormatter
    for _, item := range table {
        formatter = JSONFormatter{}
        res, err := formatter.Format(&item.Instances)

        assert.Equal(t, item.ExpectedResult, res)
        assert.Nil(t, err)
    }
}