package metrics

import (
	"github.com/graphql-go/graphql"
	"github.com/elastic/go-sysinfo/types"
)

var (
	hostType			*graphql.Object
	cpuType           	*graphql.Object
	memoryType        	*graphql.Object
	diskType          	*graphql.Object
	networkType       	*graphql.Object
	processType       	*graphql.Object
	systemMetricsType 	*graphql.Object
)

func initTypes() {
	hostType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Host",
		Description: "Host info",
		Fields: graphql.Fields{
			"architecture": &graphql.Field{
				Type:        graphql.String,
				Description: "Process hardware architecture",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Architecture, nil
					}
					return nil, nil
				},
			},
			"nativeArchitecture": &graphql.Field{
				Type:        graphql.String,
				Description: "Native OS hardware architecture",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.NativeArchitecture, nil
					}
					return nil, nil
				},
			},
			"bootTime": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "Host boot time",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.BootTime, nil
					}
					return nil, nil
				},
			},
			"containerized": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Is the process containerized",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok && host.Containerized != nil {
						return *host.Containerized, nil
					}
					return nil, nil
				},
			},
			"hostname": &graphql.Field{
				Type:        graphql.String,
				Description: "Hostname",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Hostname, nil
					}
					return nil, nil
				},
			},
			"ips": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "List of all IPs",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.IPs, nil
					}
					return nil, nil
				},
			},
			"kernelVersion": &graphql.Field{
				Type:        graphql.String,
				Description: "Kernel version",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.KernelVersion, nil
					}
					return nil, nil
				},
			},
			"macs": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "List of MAC addresses",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.MACs, nil
					}
					return nil, nil
				},
			},
			"os": &graphql.Field{
				Type:        graphql.String, // Assume osType is defined elsewhere
				Description: "OS information",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.OS.Name + " " + host.OS.Version, nil
					}
					return nil, nil
				},
			},
			"timezone": &graphql.Field{
				Type:        graphql.String,
				Description: "System timezone",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Timezone, nil
					}
					return nil, nil
				},
			},
			"timezoneOffsetSec": &graphql.Field{
				Type:        graphql.Int,
				Description: "Timezone offset (seconds from UTC)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.TimezoneOffsetSec, nil
					}
					return nil, nil
				},
			},
			"uniqueID": &graphql.Field{
				Type:        graphql.String,
				Description: "Unique ID of the host (optional)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.UniqueID, nil
					}
					return nil, nil
				},
			},
		},
	})

	

}
