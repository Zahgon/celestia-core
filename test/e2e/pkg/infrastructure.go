package e2e

import (
	"net"
)

const (
	dockerIPv4CIDR = "10.186.73.0/24"
	dockerIPv6CIDR = "fd80:b10c::/48"
)

// InfrastructureData contains the relevant information for a set of existing
// infrastructure that is to be used for running a testnet.
type InfrastructureData struct {
	Path string

	// Provider is the name of infrastructure provider backing the testnet.
	// Currently, only 'docker' is supported.
	Provider string `json:"provider"`

	// Instances is a map of all of the machine instances on which to run
	// processes for a testnet.
	// The key of the map is the name of the instance, which each must correspond
	// to the names of one of the testnet nodes defined in the testnet manifest.
	Instances map[string]InstanceData `json:"instances"`

	// Network is the CIDR notation range of IP addresses that all of the instances'
	// IP addresses are expected to be within.
	Network string `json:"network"`

	// TracePushConfig is the URL of the server to push trace data to.
	TracePushConfig string `json:"trace_push_config,omitempty"`

	// TracePullAddress is the address to listen on for pulling trace data.
	TracePullAddress string `json:"trace_pull_address,omitempty"`

	// PyroscopeURL is the URL of the pyroscope instance to use for continuous
	// profiling. If not specified, data will not be collected.
	PyroscopeURL string `json:"pyroscope_url,omitempty"`

	// PyroscopeTrace enables adding trace data to pyroscope profiling.
	PyroscopeTrace bool `json:"pyroscope_trace,omitempty"`

	// PyroscopeProfileTypes is the list of profile types to collect.
	PyroscopeProfileTypes []string `json:"pyroscope_profile_types,omitempty"`
}

// InstanceData contains the relevant information for a machine instance backing
// one of the nodes in the testnet.
type InstanceData struct {
	IPAddress    net.IP `json:"ip_address"`
	ExtIPAddress net.IP `json:"ext_ip_address"`
	Port         uint32 `json:"port"`
}

func sortNodeNames(m Manifest) []string {
	_ = "STUB: not implemented"
	// Set up nodes, in alphabetical order (IPs and ports get same order).
	return nil
}

//nolint:prealloc

func NewDockerInfrastructureData(m Manifest) (InfrastructureData, error) {
	_ = "STUB: not implemented"
	return *new(InfrastructureData), nil
}
