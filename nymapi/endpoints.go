package nymapi

const (
	EndpointTemplate = "https://%s/api/v1/%s"

	// Api Status
	EndpointBuildInformation = "api-status/build-information"
	EndpointHealth           = "api-status/health"

	// Nodes
	EndpointNodeAnnotation  = "nym-nodes/annotation/{node_id}"
	EndpointNodePerformance = "nym-nodes/performance/{node_id}"
)

var (
	Endpoints = []string{
		EndpointBuildInformation,
		EndpointHealth,
		EndpointNodeAnnotation,
		EndpointNodePerformance,
	}
)
