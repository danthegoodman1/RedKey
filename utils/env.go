package utils

import "os"

var (
	Env                    = os.Getenv("ENV")
	Env_TracingServiceName = os.Getenv("TRACING_SERVICE_NAME")
	Env_OLTPEndpoint       = os.Getenv("OLTP_ENDPOINT")

	REDIS_LISTEN = GetEnvOrDefault("REDIS_LISTEN", "0.0.0.0:6379")
)
