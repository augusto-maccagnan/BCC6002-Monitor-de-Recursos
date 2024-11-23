package model

// Representa Recursos
type Resources struct {
	Cpu []CpuParams `json:"cpu"`
	CpuCores []CpuCores `json:"cpu_cores"`
	Memory []Memory `json:"memory"`
}

type CpuParams struct {
	Name string `json:"name"`
	Frequency string `json:"frequency"`
}

type CpuCores struct {
	Name string `json:"name"`
	Frequency string `json:"frequency"`
	Percentage string `json:"percentage"`
}

type Memory struct {
	Name string `json:"name"`
	Quantity string `json:"quantity"`
	Percentage string `json:"percentage"`
}