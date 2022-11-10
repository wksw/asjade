package config

// Config .
type Config struct {
	App           string
	Service       *Service
	ConfigCenter  *ConfigCenter
	ServiceCenter *ServiceCenter
}

// Service service config
type Service struct {
	// service name
	Name string
	// service version
	Version string
}

// ConfigCenter config center config
type ConfigCenter struct{}

// ServiceCenter service center config
type ServiceCenter struct{}

// Init config init
func Init() error {
	return nil
}
