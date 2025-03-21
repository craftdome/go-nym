package models

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
