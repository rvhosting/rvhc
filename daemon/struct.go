package daemon

type ListenConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ConfigStruct struct {
	Listen ListenConfig `json:"listen"`
}
