package types

type SystemRouterboard struct {
	BoardName       string `json:"board-name,omitempty"`
	CurrentFirmware string `json:"current-firmware,omitempty"`
	FactoryFirmware string `json:"factory-firmware,omitempty"`
	FirmwareType    string `json:"firmware-type,omitempty"`
	Model           string `json:"model,omitempty"`
	Routerboard     string `json:"routerboard,omitempty"`
	SerialNumber    string `json:"serial-number,omitempty"`
	UpgradeFirmware string `json:"upgrade-firmware,omitempty"`
}

func (t *SystemRouterboard) GetRestPath() string {
	return "/system/routerboard"
}

type SystemResource struct {
	ArchitectureName     string `json:"architecture-name,omitempty"`
	BadBlocks            string `json:"bad-blocks,omitempty"`
	BoardName            string `json:"board-name,omitempty"`
	BuildTime            string `json:"build-time,omitempty"`
	CPU                  string `json:"cpu,omitempty"`
	CPUCount             string `json:"cpu-count,omitempty"`
	CPUFrequency         string `json:"cpu-frequency,omitempty"`
	CPULoad              string `json:"cpu-load,omitempty"`
	FactorySoftware      string `json:"factory-software,omitempty"`
	FreeHddSpace         string `json:"free-hdd-space,omitempty"`
	FreeMemory           string `json:"free-memory,omitempty"`
	Platform             string `json:"platform,omitempty"`
	TotalHddSpace        string `json:"total-hdd-space,omitempty"`
	TotalMemory          string `json:"total-memory,omitempty"`
	Uptime               string `json:"uptime,omitempty"`
	Version              string `json:"version,omitempty"`
	WriteSectSinceReboot string `json:"write-sect-since-reboot,omitempty"`
	WriteSectTotal       string `json:"write-sect-total,omitempty"`
}

func (t *SystemResource) GetRestPath() string {
	return "/system/resource"
}

type SystemResourceCPU struct {
	ID   string `json:".id,omitempty"`
	CPU  string `json:"cpu,omitempty"`
	Disk string `json:"disk,omitempty"`
	Irq  string `json:"irq,omitempty"`
	Load string `json:"load,omitempty"`
}

type SystemResourceCPUs []SystemResourceCPU

func (t *SystemResourceCPUs) GetRestPath() string {
	return "/system/resource/cpu"
}

func (t *SystemResourceCPU) GetRestPath() string {
	return "/system/resource/cpu"
}

type SystemResourceIRQ struct {
	ID          string `json:".id,omitempty"`
	ActiveCPU   string `json:"active-cpu,omitempty"`
	Count       string `json:"count,omitempty"`
	CPU         string `json:"cpu,omitempty"`
	Irq         string `json:"irq,omitempty"`
	PerCPUCount string `json:"per-cpu-count,omitempty"`
	ReadOnly    string `json:"read-only,omitempty"`
	Users       string `json:"users,omitempty"`
}

func (t *SystemResourceIRQ) GetRestPath() string {
	return "/system/resource/irq"
}

type SystemResourceIRQs []SystemResourceIRQ

func (t *SystemResourceIRQs) GetRestPath() string {
	return "/system/resource/irq"
}

type SystemResourceUSB struct {
	ID           string `json:".id,omitempty"`
	Device       string `json:"device,omitempty"`
	DeviceID     string `json:"device-id,omitempty"`
	Inactive     string `json:"inactive,omitempty"`
	Name         string `json:"name,omitempty"`
	Ports        string `json:"ports,omitempty"`
	SerialNumber string `json:"serial-number,omitempty"`
	Speed        string `json:"speed,omitempty"`
	UsbVersion   string `json:"usb-version,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	VendorID     string `json:"vendor-id,omitempty"`
}

func (t *SystemResourceUSB) GetRestPath() string {
	return "/system/resource/usb"
}

type SystemResourceUSBs []SystemResourceUSB

func (t *SystemResourceUSBs) GetRestPath() string {
	return "/system/resource/usb"
}
