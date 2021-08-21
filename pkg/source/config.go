package source

// Config is the configuration struct for a source
type Config struct {
	Type     string                 `yaml:"type"`
	Settings map[string]interface{} `yaml:"settings"`
}
