package models

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
