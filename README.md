Service Discovery
=================

Software is meant to discover running servers in your infrastructure. Currently, only the AWS infrastructure is supported. 

The script allows finding all running servers by:

- AutoScaling Group name


#### Usage of service-discovery:

  - **asg-name** string Autoscalling group name
  - **field** string One of the: all, private-ip, public-ip (default "all")
  - **format** string Text format (default "text")
  - **region** string	Region of AWS service (default "us-east-1")


### Examples

#### Search services by auto-scalling group 

```
$ service-discovery --asg-name=prod-cluster-asg
i-1234567XXXXXXXXXX	10.2.4.221	3.21.21.33
i-3333333XXXXXXXXXX	10.2.4.111	35.21.21.33
i-4444444XXXXXXXXXX	10.2.1.123	54.21.21.33
i-5555555XXXXXXXXXX	10.2.1.142	34.21.21.33
```

#### Get only private ip of running services by auto-scalling group 

```
$ service-discovery --asg-name=prod-cluster-asg --field=private-ip
10.2.4.221
10.2.4.111
10.2.1.123
10.2.1.142
```

#### Get output in JSON

```
$ service-discovery --asg-name=prod-cluster-asg --format=json
[
  {
    "ID": "i-1234567XXXXXXXXXX",
    "PublicIP": "3.21.21.33",
    "PrivateIP": "10.2.4.221"
  },
  {
    "ID": "i-3333333XXXXXXXXXX",
    "PublicIP": "35.21.21.33",
    "PrivateIP": "10.2.4.111"
  },
  {
    "ID": "i-4444444XXXXXXXXXX",
    "PublicIP": "54.21.21.33",
    "PrivateIP": "10.2.1.123"
  },
  {
    "ID": "i-5555555XXXXXXXXXX",
    "PublicIP": "34.21.21.33",
    "PrivateIP": "10.2.1.142"
  }
]

```
