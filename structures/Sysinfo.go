package structures

type Sysinfo struct {
	Name string `json:"name"`
	Model string `json:"model"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	MBSerial string `json:"MBSerial"`
	Arch string `json:"arch"`
	Core string `json:"core"`
	Os string `json:"os"`
	RDMS_Version string `json:"rdms_version"`
	Anydesk string `json:"anydesk"`
	Packages []Package `json:"packages"`
}