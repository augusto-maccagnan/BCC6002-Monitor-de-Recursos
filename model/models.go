package model

// Representa Recursos
type Resources struct {
	Cpu []CpuParams `json:"cpu"`
	CpuCores []CpuCores `json:"cpu_cores"`
	Memory []Memory `json:"memory"`
	Disk []Disk `json:"disk"`
}

type CpuParams struct {
	MaxFrequency int `json:"max_frequency"`
	MinFrequency int `json:"min_frequency"`
	CoreNumber int `json:"core_number"`
	TotalUse float64 `json:"total_use"`
	Percentage float64 `json:"percentage"`
}

type CpuCores struct {
	Name string `json:"name"`
	Frequency int `json:"frequency"`
	Percentage float64 `json:"percentage"`
}

type Memory struct {
	Total float64 `json:"total"`
	Available float64 `json:"available"`
	Free float64 `json:"free"`
	Used float64 `json:"used"`
	Percentage float64 `json:"percentage"`
}

type Disk struct {
	Name string `json:"name"`
	Total int `json:"total"`
	Used int `json:"used"`
	Free int `json:"free"`
	Percentage float64 `json:"percentage"`
}