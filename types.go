package routeros

type SystemRouterboard struct {
	BoardName       string `json:"board-name,omitempty,omitempty"`
	CurrentFirmware string `json:"current-firmwareomitempty,omitempty"`
	FactoryFirmware string `json:"factory-firmwareomitempty,omitempty"`
	FirmwareType    string `json:"firmware-type,omitempty"`
	Model           string `json:"model,omitempty"`
	Routerboard     string `json:"routerboard,omitempty"`
	SerialNumber    string `json:"serial-number,omitempty"`
	UpgradeFirmware string `json:"upgrade-firmware,omitempty"`
}

type IPFirewallAddresslists []IPFirewallAddresslist

type IPFirewallAddresslist struct {
	ID           string `json:".id,omitempty"`
	Address      string `json:"address,omitempty"`
	CreationTime string `json:"creation-time,omitempty"`
	Disabled     string `json:"disabled,omitempty"`
	Dynamic      string `json:"dynamic,omitempty"`
	List         string `json:"list,omitempty"`
}
