package nymapi

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

type Health struct {
	Status      string `json:"status"`
	ChainStatus string `json:"chain_status"`
	Uptime      uint64 `json:"uptime"`
}

type NodeAnnotation struct {
	NodeID     uint32 `json:"node_id"`
	Annotation struct {
		Last24hPerformance  float32 `json:"last_24h_performance,string"`
		DetailedPerformance struct {
			PerformanceScore float32 `json:"performance_score"`
			RoutingScore     struct {
				Score float32 `json:"score"`
			} `json:"routing_score"`
			ConfigScore struct {
				Score                      float32 `json:"score"`
				VersionsBehind             uint16  `json:"versions_behind"`
				SelfDescribedAPIAvailable  bool    `json:"self_described_api_available"`
				AcceptedTermsAndConditions bool    `json:"accepted_terms_and_conditions"`
				RunsNymNodeBinary          bool    `json:"runs_nym_node_binary"`
			} `json:"config_score"`
		} `json:"detailed_performance"`
	} `json:"annotation"`
}

type NodePerformance struct {
	Date        string  `json:"date"`
	Performance float32 `json:"performance"`
}
