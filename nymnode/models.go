package nymnode

type MixnetWebsockets struct {
	WsPort  uint16 `json:"ws_port"`
	WssPort uint16 `json:"wss_port,omitempty"`
}

type Wireguard struct {
	Port      uint16 `json:"port"`
	PublicKey string `json:"public_key"`
}

type ClientInterfaces struct {
	MixnetWebsockets `json:"mixnet_websockets,omitempty"`
	Wireguard        `json:"wireguard,omitempty"`
}

type Gateway struct {
	ClientInterfaces `json:"client_interfaces,omitempty"`

	EnforcesZKNyms bool `json:"enforces_zk_nyms"`
}

type Health struct {
	Status string `json:"status"`

	// Uptime in seconds
	Uptime uint64 `json:"uptime"`
}

type IPPacketRouter struct {
	EncodedIdentityKey string `json:"encoded_identity_key"`
	EncodedX25519Key   string `json:"encoded_x25519_key"`
	Address            string `json:"address"`
}

type PacketsStatsMetrics struct {
	IngressMixing struct {
		ForwardHopPacketsReceived uint64 `json:"forward_hop_packets_received"`
		FinalHopPacketsReceived   uint64 `json:"final_hop_packets_received"`
		MalformedPacketsReceived  uint64 `json:"malformed_packets_received"`
		ExcessiveDelayPackets     uint64 `json:"excessive_delay_packets"`
		ForwardHopPacketsDropped  uint64 `json:"forward_hop_packets_dropped"`
		FinalHopPacketsDropped    uint64 `json:"final_hop_packets_dropped"`
	} `json:"ingress_mixing"`
	EgressMixing struct {
		ForwardHopPacketsSent    uint64 `json:"forward_hop_packets_sent"`
		ForwardHopPacketsDropped uint64 `json:"forward_hop_packets_dropped"`
		AckPacketsSent           uint64 `json:"ack_packets_sent"`
	} `json:"egress_mixing"`
}

type PrometheusMetrics string

type AuxiliaryDetails struct {
	Location      string `json:"location"`
	AnnouncePorts struct {
		VerlocPort uint16 `json:"verloc_port,omitempty"`
		MixPort    uint16 `json:"mix_port,omitempty"`
	} `json:"announce_ports"`
	AcceptedOperatorTermsAndConditions bool `json:"accepted_operator_terms_and_conditions"`
}

type BuildInformation struct {
	BinaryName      string `json:"binary_name"`
	BuildTimestamp  string `json:"build_timestamp"`
	BuildVersion    string `json:"build_version"`
	CommitSHA       string `json:"commit_sha"`
	CommitTimestamp string `json:"commit_timestamp"`
	CommitBranch    string `json:"commit_branch"`
	RustcVersion    string `json:"rustc_version"`
	RustcChannel    string `json:"rustc_channel"`
	CargoProfile    string `json:"cargo_profile"`
	CargoTriple     string `json:"cargo_triple"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type HostInformation struct {
	Data struct {
		IPs      []string `json:"ip_address"`
		Hostname string   `json:"hostname,omitempty"`
		Keys     struct {
			ED25519Identity string `json:"ed25519_identity"`
			X25519Sphinx    string `json:"x25519_sphinx"`
			X25519Noise     string `json:"x25519_noise,omitempty"`
		} `json:"keys"`
	} `json:"data"`
	Signature string `json:"signature"`
}

type Roles struct {
	MixnodeEnabled          bool `json:"mixnode_enabled"`
	GatewayEnabled          bool `json:"gateway_enabled"`
	NetworkRequesterEnabled bool `json:"network_requester_enabled"`
	IpPacketRouterEnabled   bool `json:"ip_packet_router_enabled"`
}

type SystemInformation struct {
	SystemName    string `json:"system_name"`
	KernelVersion string `json:"kernel_version"`
	OSVersion     string `json:"os_version"`
	Hardware      struct {
		CPU []struct {
			Brand     string `json:"brand"`
			Frequency uint32 `json:"frequency"`
			Name      string `json:"name"`
			VendorID  string `json:"vendor_id"`
		} `json:"cpu"`
		Crypto struct {
			AESNI                    bool     `json:"aesni"`
			AVX2                     bool     `json:"avx2"`
			OSXSAVE                  bool     `json:"osxsave"`
			SMTLogicalProcessorCount []uint32 `json:"smt_logical_processor_count"`
			SGX                      bool     `json:"sgx"`
			XSAVE                    bool     `json:"xsave"`
		} `json:"crypto"`
		TotalMemory uint64 `json:"total_memory"`
	} `json:"hardware"`
}

type PolicyAction string

const (
	Accept  PolicyAction = "accept"
	Reject  PolicyAction = "reject"
	Accept6 PolicyAction = "accept6"
	Reject6 PolicyAction = "reject6"
)

type NetworkRequester struct {
	EncodedIdentityKey string `json:"encoded_identity_key"`
	EncodedX25519Key   string `json:"encoded_x25519_key"`
	Address            string `json:"address"`
}

type NetworkRequesterExitPolicy struct {
	Enabled        bool   `json:"enabled"`
	UpstreamSource string `json:"upstream_source"`
	LastUpdated    uint64 `json:"last_updated"`
	Policy         *struct {
		Rules []struct {
			Action  PolicyAction `json:"action"`
			Pattern struct {
				IpPattern string `json:"ip_pattern"`
				Ports     struct {
					Start uint16 `json:"start"`
					End   uint16 `json:"end"`
				} `json:"ports"`
			} `json:"pattern"`
		} `json:"rules"`
	} `json:"policy"`
}
