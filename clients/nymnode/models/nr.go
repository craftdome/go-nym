package models

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
