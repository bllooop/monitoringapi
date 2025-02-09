package configs

const (
	DockerAPIURL = "http://localhost:2375/containers/json"
	BackendURL   = "http://backend:8000/api/data/create"
	PingTimeout  = 2
	AmqpURL      = "amqp://guest:guest@message-broker:5672/"
)
