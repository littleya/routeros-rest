package types

type IPAddress struct {
	ID              string `json:".id,omitempty"`
	ActualInterface string `json:"actual-interface,omitempty"`
	Address         string `json:"address,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Disabled        string `json:"disabled,omitempty"`
	Dynamic         string `json:"dynamic,omitempty"`
	Interface       string `json:"interface,omitempty"`
	Invalid         string `json:"invalid,omitempty"`
	Network         string `json:"network,omitempty"`
}

func (t *IPAddress) GetRestPath() string {
	return "/ip/address"
}

type IPAddresses []IPAddress

func (t *IPAddresses) GetRestPath() string {
	return "/ip/address"
}

type IPARP struct {
	ID         string `json:".id,omitempty"`
	Dhcp       string `json:"DHCP,omitempty"`
	Address    string `json:"address,omitempty"`
	Comment    string `json:"comment,omitempty"`
	Complete   string `json:"complete,omitempty"`
	Disabled   string `json:"disabled,omitempty"`
	Dynamic    string `json:"dynamic,omitempty"`
	Interface  string `json:"interface,omitempty"`
	Invalid    string `json:"invalid,omitempty"`
	MacAddress string `json:"mac-address,omitempty"`
	Published  string `json:"published,omitempty"`
}

func (t *IPARP) GetRestPath() string {
	return "/ip/arp"
}

type IPARPs []IPARP

func (t *IPARPs) GetRestPath() string {
	return "/ip/arp"
}

type IPDHCPServer struct {
	ID            string `json:".id,omitempty"`
	AddressPool   string `json:"address-pool,omitempty"`
	Authoritative string `json:"authoritative,omitempty"`
	Disabled      string `json:"disabled,omitempty"`
	Dynamic       string `json:"dynamic,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Interface     string `json:"interface,omitempty"`
	Invalid       string `json:"invalid,omitempty"`
	LeaseScript   string `json:"lease-script,omitempty"`
	LeaseTime     string `json:"lease-time,omitempty"`
	Name          string `json:"name,omitempty"`
	UseRadius     string `json:"use-radius,omitempty"`
}

func (t *IPDHCPServer) GetRestPath() string {
	return "/ip/dhcp-server"
}

type IPDHCPServers []IPDHCPServer

func (t *IPDHCPServers) GetRestPath() string {
	return "/ip/dhcp-server"
}

type IPDHCPServerLease struct {
	ID               string `json:".id,omitempty"`
	ActiveAddress    string `json:"active-address,omitempty"`
	ActiveMacAddress string `json:"active-mac-address,omitempty"`
	ActiveServer     string `json:"active-server,omitempty"`
	Address          string `json:"address,omitempty"`
	AddressLists     string `json:"address-lists,omitempty"`
	Blocked          string `json:"blocked,omitempty"`
	DhcpOption       string `json:"dhcp-option,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Disabled         string `json:"disabled,omitempty"`
	Dynamic          string `json:"dynamic,omitempty"`
	ExpiresAfter     string `json:"expires-after,omitempty"`
	HostName         string `json:"host-name,omitempty"`
	LastSeen         string `json:"last-seen,omitempty"`
	MacAddress       string `json:"mac-address,omitempty"`
	Radius           string `json:"radius,omitempty"`
	Server           string `json:"server,omitempty"`
	Status           string `json:"status,omitempty"`
}

func (t *IPDHCPServerLease) GetRestPath() string {
	return "/ip/dhcp-server/lease"
}

type IPDHCPServerLeases []IPDHCPServerLease

func (t *IPDHCPServerLeases) GetRestPath() string {
	return "/ip/dhcp-server/lease"
}

type IPDHCPServerNetwork struct {
	ID          string `json:".id,omitempty"`
	Address     string `json:"address,omitempty"`
	CapsManager string `json:"caps-manager,omitempty"`
	Comment     string `json:"comment,omitempty"`
	DhcpOption  string `json:"dhcp-option,omitempty"`
	DNSServer   string `json:"dns-server,omitempty"`
	Dynamic     string `json:"dynamic,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	NtpServer   string `json:"ntp-server,omitempty"`
	WinsServer  string `json:"wins-server,omitempty"`
}

func (t *IPDHCPServerNetwork) GetRestPath() string {
	return "/ip/dhcp-server/network"
}

type IPDHCPServerNetworks []IPDHCPServerNetwork

func (t *IPDHCPServerNetworks) GetRestPath() string {
	return "/ip/dhcp-server/network"
}

type IPFirewallAddresslist struct {
	ID           string `json:".id,omitempty"`
	Address      string `json:"address,omitempty"`
	CreationTime string `json:"creation-time,omitempty"`
	Disabled     string `json:"disabled,omitempty"`
	Dynamic      string `json:"dynamic,omitempty"`
	List         string `json:"list,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

func (t *IPFirewallAddresslist) GetRestPath() string {
	return "/ip/firewall/address-list"
}

type IPFirewallAddresslists []IPFirewallAddresslist

func (t *IPFirewallAddresslists) GetRestPath() string {
	return "/ip/firewall/address-list"
}

type IPRoute struct {
	ID                string `json:".id,omitempty"`
	Active            string `json:"active,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Disabled          string `json:"disabled,omitempty"`
	Distance          string `json:"distance,omitempty"`
	DstAddress        string `json:"dst-address,omitempty"`
	Dynamic           string `json:"dynamic,omitempty"`
	Ecmp              string `json:"ecmp,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	HwOffloaded       string `json:"hw-offloaded,omitempty"`
	ImmediateGw       string `json:"immediate-gw,omitempty"`
	Inactive          string `json:"inactive,omitempty"`
	PrefSrc           string `json:"pref-src,omitempty"`
	RoutingTable      string `json:"routing-table,omitempty"`
	Scope             string `json:"scope,omitempty"`
	Static            string `json:"static,omitempty"`
	SuppressHwOffload string `json:"suppress-hw-offload,omitempty"`
	TargetScope       string `json:"target-scope,omitempty"`
}

func (t *IPRoute) GetRestPath() string {
	return "/ip/route"
}

type IPRoutes []IPRoute

func (t *IPRoutes) GetRestPath() string {
	return "/ip/route"
}
