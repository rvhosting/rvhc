package daemon

type ListenConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type daemonConfig struct {
	Listen ListenConfig `json:"listen"`
}
