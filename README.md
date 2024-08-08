# Gometric

![build](https://github.com/davidjosearaujo/gometric/actions/workflows/go.yml/badge.svg)

## Overview

This project is a metrics server built with Go that collects system metrics and exposes them via a GraphQL API. It includes support for CPU, memory, disk, network, and process metrics. Deployment can be automated using Terraform and Ansible.

![image](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![image](https://img.shields.io/badge/GraphQl-E10098?style=for-the-badge&logo=graphql&logoColor=white)
![image](https://img.shields.io/badge/Terraform-7B42BC?style=for-the-badge&logo=terraform&logoColor=white)

## Features

- System Metrics: Collects detailed CPU, memory, disk, network, and process metrics.
- GraphQL API: Exposes collected metrics via a GraphQL API.
- Deployment Automation: Supports deployment automation using Terraform and Ansible.

## Installation

```bash
git clone https://github.com/davidjosearaujo/gometric
cd gometric
go build -o gometric
```

## Usage

Running the server

```bash
./gometric
```

Query the server

```bash
$ curl -g 'http://localhost:7000/gometric?query={disk(device:"/dev/nvme0n1p1"){fstype}}'
{"data":{"disk":{"fstype":"ext2/ext3"}}}
```

## API Documentation

Bellow you can find the specification of the GraphQL schema, query and types.

```graphql
enum timeEnum {
    ONE
    FIVE
    FIFTEEN
}

enum cpuTimeEnum {
    USER
    SYSTEM
    IDLE
    IOWAIT
    IRQ
    NICE
    SOFTIRQ
    STEAL
}

enum protocolNetstatEnum {
    TCP
    IP
}

enum protocolSNMPEnum {
    IP
    ICMP
    ICMPMsg
    TCP
    UDP
    UDPLite
}

enum modeEnum {
    BYTES
    PERCENT
}

type cpuType {
    load(time: timeEnum):       String!
    times(stat: cpuTimeEnum):   String!
    cores:                      Int!
    info:                       String!
}

type hostType {
    architecture:       String!
    nativeArchitecture: String!
    bootTime:           Date!
    uptime:             String!
    containerized:      Boolean!
    hostname:           String!
    ips:                [String]!
    kernelVersion:      String!
    macs:               [String]!
    os:                 String!
    timezone:           String!
    timezoneOffsetSec:  Int!
    uniqueID:           String!
}

type memoryType {
    total:          String!
    used:           String!
    available:      String!
    free:           String!
    virtualTotal:   String!
    virtualUsed:    String!
    virtualFree:    String!
}

type networkType {
    netstat(protocol: protocolNetstatEnum, counter: String):    String!
    snmp(protocol: protocolSNMPEnum, counter: String):          String!
}

type osType {
    type:       String!
    family:     String!
    platform:   String!
    name:       String!
    version:    String!
    major:      String!
    minor:      String!
    patch:      String!
    build:      String!
    codename:   String!
}

type diskType {
    devices:                    [String]!
    fstype:                     String!
    mountpoint:                 String!
    opts:                       String!
    total:                      String!
    free:                       String!
    used(mode: modeEnum):       String!
    inodestotal:                String!
    inodesused(mode: modeEnum): String!
    inodesfree:                 String!            
}

type Query {
    cpu:                    cpuType
    host:                   hostType
    memory:                 memoryType
    network:                networkType
    os:                     osType
    disk(device: String):   diskType
}

schema {
    query: Query
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the Apache License. See the [LICENSE](./LICENSE) file for details.
