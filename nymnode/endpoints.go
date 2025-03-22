package nymnode

const (
	EndpointTemplate = "http://%s/api/v1/%s"

	// Node
	EndpointAuxiliaryDetails  = "auxiliary-details"
	EndpointBuildInformation  = "build-information"
	EndpointDescription       = "description"
	EndpointHostInformation   = "host-information"
	EndpointRoles             = "roles"
	EndpointSystemInformation = "system-info"

	// Gateway
	EndpointGateway                                 = "gateway"
	EndpointGatewayClientInterfaces                 = "gateway/client-interfaces"
	EndpointGatewayClientInterfacesMixnetWebsockets = "gateway/client-interfaces/mixnet-websockets"

	// Health
	EndpointHealth = "health"

	// IP Packet Router
	EndpointIPR = "ip-packet-router"

	// Metrics
	EndpointMetricsPacketsStats = "metrics/packets-stats"
	EndpointMetricsPrometheus   = "metrics/prometheus"
	//EndpointMetricsVerloc       = "metrics/verloc"

	// Network Requester
	EndpointNR           = "network-requester"
	EndpointNRExitPolicy = "network-requester/exit-policy"
)

var (
	Endpoints = []string{
		EndpointAuxiliaryDetails,
		EndpointBuildInformation,
		EndpointDescription,
		EndpointHostInformation,
		EndpointRoles,
		EndpointSystemInformation,

		EndpointGateway,
		EndpointGatewayClientInterfaces,
		EndpointGatewayClientInterfacesMixnetWebsockets,

		EndpointHealth,

		EndpointIPR,

		EndpointMetricsPacketsStats,
		EndpointMetricsPrometheus,

		EndpointNR,
		EndpointNRExitPolicy,
	}
)
