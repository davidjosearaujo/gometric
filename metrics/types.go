package metrics

import (
	"reflect"

	"github.com/elastic/go-sysinfo/types"
	"github.com/graphql-go/graphql"
)

var (
	hostType    *graphql.Object
	osType      *graphql.Object
	cpuType     *graphql.Object
	memoryType  *graphql.Object
	diskType    *graphql.Object // TODO
	networkType *graphql.Object
	processType *graphql.Object
)

func initTypes() {
	timeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Time",
		Description: "One of the time windows for CPU usage",
		Values: graphql.EnumValueConfigMap{
			"ONE": &graphql.EnumValueConfig{
				Value: "One",
			},
			"FIVE": &graphql.EnumValueConfig{
				Value: "Five",
			},
			"FIFTEEN": &graphql.EnumValueConfig{
				Value: "Fifteen",
			},
		},
	})

	cpuTimeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Time",
		Description: "One of the time windows for CPU usage",
		Values: graphql.EnumValueConfigMap{
			"USER": &graphql.EnumValueConfig{
				Value: "User",
			},
			"SYSTEM": &graphql.EnumValueConfig{
				Value: "System",
			},
			"IDLE": &graphql.EnumValueConfig{
				Value: "Idle",
			},
			"IOWAIT": &graphql.EnumValueConfig{
				Value: "IOWait",
			},
			"IRQ": &graphql.EnumValueConfig{
				Value: "IRQ",
			},
			"NICE": &graphql.EnumValueConfig{
				Value: "Nice",
			},
			"SOFTIRQ": &graphql.EnumValueConfig{
				Value: "SoftIRQ",
			},
			"STEAL": &graphql.EnumValueConfig{
				Value: "Steal",
			},
		},
	})

	protocolNetstatEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Protocol",
		Description: "Netstat protocol",
		Values: graphql.EnumValueConfigMap{
			"TCP": &graphql.EnumValueConfig{
				Value: "TCP",
			},
			"IP": &graphql.EnumValueConfig{
				Value: "IP",
			},
		},
	})

	protocolSNMPEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Protocol",
		Description: "Netstat protocol",
		Values: graphql.EnumValueConfigMap{
			"IP": &graphql.EnumValueConfig{
				Value: "IP",
			},
			"ICMP": &graphql.EnumValueConfig{
				Value: "ICMP",
			},
			"ICMPMsg": &graphql.EnumValueConfig{
				Value: "ICMPMsg",
			},
			"TCP": &graphql.EnumValueConfig{
				Value: "TCP",
			},
			"UDP": &graphql.EnumValueConfig{
				Value: "UDP",
			},
			"UDPLite": &graphql.EnumValueConfig{
				Value: "UDPLite",
			},
		},
	})

	modeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Mode",
		Description: "Either in bytes or percentage",
		Values: graphql.EnumValueConfigMap{
			"BYTES": &graphql.EnumValueConfig{
				Value: false,
			},
			"PERCENT": &graphql.EnumValueConfig{
				Value: true,
			},
		},
	})

	cpuType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CPU",
		Description: "CPU info",
		Fields: graphql.Fields{
			"load": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Process hardware architecture",
				Args: graphql.FieldConfigArgument{
					"time": &graphql.ArgumentConfig{
						Description: "There are three available time windows, 1, 5 and 15 minutes",
						Type:        timeEnum,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						if p.Args["time"] != nil {
							return reflect.ValueOf(*cpu.Load).FieldByName(p.Args["time"].(string)), nil
						}
						return *cpu.Load, nil
					}
					return nil, nil
				},
			},
			"times": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Timing stats for a process",
				Args: graphql.FieldConfigArgument{
					"stat": &graphql.ArgumentConfig{
						Description: "There are three available time windows, 1, 5 and 15 minutes",
						Type:        cpuTimeEnum,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						if p.Args["stat"] != nil {
							return reflect.ValueOf(cpu.Time).FieldByName(p.Args["stat"].(string)), nil
						}
						return cpu.Time.Total().String(), nil
					}
					return nil, nil
				},
			},
			"cores": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Number of cores in the CPU",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						return cpu.CoreCount, nil
					}
					return nil, nil
				},
			},
			"info": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Overall CPU info",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						return cpu.Info, nil
					}
					return nil, nil
				},
			},
		},
	})

	hostType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Host",
		Description: "Host info",
		Fields: graphql.Fields{
			"architecture": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Process hardware architecture",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Architecture, nil
					}
					return nil, nil
				},
			},
			"nativeArchitecture": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Native OS hardware architecture",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.NativeArchitecture, nil
					}
					return nil, nil
				},
			},
			"bootTime": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "Host boot time",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.BootTime, nil
					}
					return nil, nil
				},
			},
			"uptime": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Host uptime",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Uptime().String(), nil
					}
					return nil, nil
				},
			},
			"containerized": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Boolean),
				Description: "Is the process containerized",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok && host.Containerized != nil {
						return *host.Containerized, nil
					}
					return nil, nil
				},
			},
			"hostname": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
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
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Kernel version",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.KernelVersion, nil
					}
					return nil, nil
				},
			},
			"macs": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(graphql.String)),
				Description: "List of MAC addresses",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.MACs, nil
					}
					return nil, nil
				},
			},
			"os": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS information",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.OS.Name + " " + host.OS.Version, nil
					}
					return nil, nil
				},
			},
			"timezone": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "System timezone",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.Timezone, nil
					}
					return nil, nil
				},
			},
			"timezoneOffsetSec": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Timezone offset (seconds from UTC)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if host, ok := p.Source.(types.HostInfo); ok {
						return host.TimezoneOffsetSec, nil
					}
					return nil, nil
				},
			},
			"uniqueID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
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

	memoryType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Memory",
		Description: "Host memory info",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total physical memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.Total, nil
					}
					return nil, nil
				},
			},
			"used": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total used memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.Used, nil
					}
					return nil, nil
				},
			},
			"available": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Amount of memory available without swapping in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.Available, nil
					}
					return nil, nil
				},
			},
			"free": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Amount of memory not used by the system in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.Free, nil
					}
					return nil, nil
				},
			},
			"virtualTotal": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total virtual memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.VirtualTotal, nil
					}
					return nil, nil
				},
			},
			"virtualUsed": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total used virtual memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.VirtualUsed, nil
					}
					return nil, nil
				},
			},
			"virtualFree": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Virtual memory that is not used in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(types.HostMemoryInfo); ok {
						return mem.VirtualFree, nil
					}
					return nil, nil
				},
			},
		},
	})

	networkType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Network",
		Description: "Host network info",
		Fields: graphql.Fields{
			"netstat": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Netstat",
				Args: graphql.FieldConfigArgument{
					"protocol": &graphql.ArgumentConfig{
						Description: "Either IP or TCP",
						Type:        protocolNetstatEnum,
					},
					"counter": &graphql.ArgumentConfig{
						Description: "Netstat counter",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						var data map[string]uint64
						switch p.Args["protocol"] {
						case "IP":
							data = net.Network.Netstat.IPExt
						case "TCP":
							data = net.Network.Netstat.TCPExt
						default:
							return nil, nil
						}

						if p.Args["counter"] == nil {
							return reflect.ValueOf(data).MapKeys(), nil
						} else {
							return data[p.Args["counter"].(string)], nil
						}
					}
					return nil, nil
				},
			},
			"snmp": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "SNMP",
				Args: graphql.FieldConfigArgument{
					"protocol": &graphql.ArgumentConfig{
						Description: "Either IP, ICMP, ICMPMsg, TCP, UDP, UDPLite",
						Type:        protocolSNMPEnum,
					},
					"counter": &graphql.ArgumentConfig{
						Description: "Netstat counter",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						var data map[string]uint64
						switch p.Args["protocol"] {
						case "IP":
							data = net.Network.SNMP.IP
						case "ICMP":
							data = net.Network.SNMP.ICMP
						case "ICMPMsg":
							data = net.Network.SNMP.ICMPMsg
						case "TCP":
							data = net.Network.SNMP.TCP
						case "UDP":
							data = net.Network.SNMP.UDP
						case "UDPLite":
							data = net.Network.SNMP.UDPLite
						default:
							return nil, nil
						}

						if p.Args["counter"] == nil {
							return reflect.ValueOf(data).MapKeys(), nil
						} else {
							return data[p.Args["counter"].(string)], nil
						}
					}
					return nil, nil
				},
			},
		},
	})

	osType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "OS",
		Description: "Host OS info",
		Fields: graphql.Fields{
			"type": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS Type (one of linux, macos, unix, windows)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Type, nil
					}
					return nil, nil
				},
			},
			"family": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS Family (e.g. redhat, debian, freebsd, windows)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Family, nil
					}
					return nil, nil
				},
			},
			"platform": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS platform (e.g. centos, ubuntu, windows)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Platform, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS Name (e.g. Mac OS X, CentOS)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Name, nil
					}
					return nil, nil
				},
			},
			"version": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS version (e.g. 10.12.6)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Version, nil
					}
					return nil, nil
				},
			},
			"major": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Major release version",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Major, nil
					}
					return nil, nil
				},
			},
			"minor": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Minor release version",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Minor, nil
					}
					return nil, nil
				},
			},
			"patch": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Patch release version",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Patch, nil
					}
					return nil, nil
				},
			},
			"build": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Build (e.g. 16G1114)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Build, nil
					}
					return nil, nil
				},
			},
			"codename": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "OS codename (e.g. jessie)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if os, ok := p.Source.(types.OSInfo); ok {
						return os.Codename, nil
					}
					return nil, nil
				},
			},
		},
	})

	diskType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Disk",
		Description: "Disk info",
		Fields: graphql.Fields{
			"devices": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.NewList(graphql.String)),
				Description: "List of devices",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						var devices []string
						for _, device := range disk.Partitions {
							devices = append(devices, device.Device)
						}
						return devices, nil
					}
					return nil, nil
				},
			},
			"fstype": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Filesystem type",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.UsageStat.Fstype, nil
					}
					return nil, nil
				},
			},
			"mountpoint": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Filesystem type",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.Partitions[0].Mountpoint, nil
					}
					return nil, nil
				},
			},
			"opts": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Filesystem type",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.Partitions[0].Opts, nil
					}
					return nil, nil
				},
			},
			"total": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total storage space",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.UsageStat.Total, nil
					}
					return nil, nil
				},
			},
			"free": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Free storage space",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.UsageStat.Free, nil
					}
					return nil, nil
				},
			},
			"used": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Used storage space",
				Args: graphql.FieldConfigArgument{
					"mode": &graphql.ArgumentConfig{
						Type:        modeEnum,
						Description: "Either in BYTES or PERCENT",
						DefaultValue: false,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						if p.Args["mode"].(bool) {
							return disk.UsageStat.UsedPercent, nil
						}
						return disk.UsageStat.Used, nil
					}
					return nil, nil
				},
			},
			"inodestotal": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Total inodes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.UsageStat.InodesTotal, nil
					}
					return nil, nil
				},
			},
			"inodesused": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Used inodes",
				Args: graphql.FieldConfigArgument{
					"mode": &graphql.ArgumentConfig{
						Type:        modeEnum,
						Description: "Either in BYTES or PERCENT",
						DefaultValue: false,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						if p.Args["mode"].(bool) {
							return disk.UsageStat.InodesUsedPercent, nil
						}
						return disk.UsageStat.InodesUsed, nil
					}
					return nil, nil
				},
			},
			"inodesfree": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Free inodes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.UsageStat.InodesFree, nil
					}
					return nil, nil
				},
			},
		},
	})
}
