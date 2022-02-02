package types

type Interface struct {
	ID             string `json:".id,omitempty"`
	ActualMtu      string `json:"actual-mtu,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Disabled       string `json:"disabled,omitempty"`
	FpRxByte       string `json:"fp-rx-byte,omitempty"`
	FpRxPacket     string `json:"fp-rx-packet,omitempty"`
	FpTxByte       string `json:"fp-tx-byte,omitempty"`
	FpTxPacket     string `json:"fp-tx-packet,omitempty"`
	LastLinkUpTime string `json:"last-link-up-time,omitempty"`
	LinkDowns      string `json:"link-downs,omitempty"`
	Mtu            string `json:"mtu,omitempty"`
	Name           string `json:"name,omitempty"`
	Running        string `json:"running,omitempty"`
	RxByte         string `json:"rx-byte,omitempty"`
	RxDrop         string `json:"rx-drop,omitempty"`
	RxError        string `json:"rx-error,omitempty"`
	RxPacket       string `json:"rx-packet,omitempty"`
	TxByte         string `json:"tx-byte,omitempty"`
	TxDrop         string `json:"tx-drop,omitempty"`
	TxError        string `json:"tx-error,omitempty"`
	TxPacket       string `json:"tx-packet,omitempty"`
	TxQueueDrop    string `json:"tx-queue-drop,omitempty"`
	Type           string `json:"type,omitempty"`
}

func (t *Interface) GetRestPath() string {
	return "/interface"
}

type Interfaces []Interface

func (t *Interfaces) GetRestPath() string {
	return "/interface"
}

type InterfaceEoIP struct {
	ID                      string `json:".id,omitempty"`
	ActualMtu               string `json:"actual-mtu,omitempty"`
	AllowFastPath           string `json:"allow-fast-path,omitempty"`
	Arp                     string `json:"arp,omitempty"`
	ArpTimeout              string `json:"arp-timeout,omitempty"`
	ClampTCPMss             string `json:"clamp-tcp-mss,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	CurrentRemoteAddress    string `json:"current-remote-address,omitempty"`
	Disabled                string `json:"disabled,omitempty"`
	DontFragment            string `json:"dont-fragment,omitempty"`
	Dscp                    string `json:"dscp,omitempty"`
	Keepalive               string `json:"keepalive,omitempty"`
	L2Mtu                   string `json:"l2mtu,omitempty"`
	LocalAddress            string `json:"local-address,omitempty"`
	LoopProtect             string `json:"loop-protect,omitempty"`
	LoopProtectDisableTime  string `json:"loop-protect-disable-time,omitempty"`
	LoopProtectSendInterval string `json:"loop-protect-send-interval,omitempty"`
	LoopProtectStatus       string `json:"loop-protect-status,omitempty"`
	MacAddress              string `json:"mac-address,omitempty"`
	Mtu                     string `json:"mtu,omitempty"`
	Name                    string `json:"name,omitempty"`
	RemoteAddress           string `json:"remote-address,omitempty"`
	Running                 string `json:"running,omitempty"`
	TunnelID                string `json:"tunnel-id,omitempty"`
}

func (t *InterfaceEoIP) GetRestPath() string {
	return "/interface/eoip"
}

type InterfaceEoIPs []InterfaceEoIP

func (t *InterfaceEoIPs) GetRestPath() string {
	return "/interface/eoip"
}

type InterfaceEthernet struct {
	ID                      string `json:".id,omitempty"`
	Advertise               string `json:"advertise,omitempty"`
	Arp                     string `json:"arp,omitempty"`
	ArpTimeout              string `json:"arp-timeout,omitempty"`
	AutoNegotiation         string `json:"auto-negotiation,omitempty"`
	Bandwidth               string `json:"bandwidth,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	DefaultName             string `json:"default-name,omitempty"`
	Disabled                string `json:"disabled,omitempty"`
	DriverRxByte            string `json:"driver-rx-byte,omitempty"`
	DriverRxPacket          string `json:"driver-rx-packet,omitempty"`
	DriverTxByte            string `json:"driver-tx-byte,omitempty"`
	DriverTxPacket          string `json:"driver-tx-packet,omitempty"`
	FullDuplex              string `json:"full-duplex,omitempty"`
	L2Mtu                   string `json:"l2mtu,omitempty"`
	LoopProtect             string `json:"loop-protect,omitempty"`
	LoopProtectDisableTime  string `json:"loop-protect-disable-time,omitempty"`
	LoopProtectSendInterval string `json:"loop-protect-send-interval,omitempty"`
	LoopProtectStatus       string `json:"loop-protect-status,omitempty"`
	MacAddress              string `json:"mac-address,omitempty"`
	Mtu                     string `json:"mtu,omitempty"`
	Name                    string `json:"name,omitempty"`
	OrigMacAddress          string `json:"orig-mac-address,omitempty"`
	Running                 string `json:"running,omitempty"`
	RxBytes                 string `json:"rx-bytes,omitempty"`
	RxFcsError              string `json:"rx-fcs-error,omitempty"`
	RxFlowControl           string `json:"rx-flow-control,omitempty"`
	RxOverflow              string `json:"rx-overflow,omitempty"`
	RxPacket                string `json:"rx-packet,omitempty"`
	RxPause                 string `json:"rx-pause,omitempty"`
	RxTooLong               string `json:"rx-too-long,omitempty"`
	RxTooShort              string `json:"rx-too-short,omitempty"`
	SfpShutdownTemperature  string `json:"sfp-shutdown-temperature,omitempty"`
	Speed                   string `json:"speed,omitempty"`
	TxBytes                 string `json:"tx-bytes,omitempty"`
	TxFlowControl           string `json:"tx-flow-control,omitempty"`
	TxPacket                string `json:"tx-packet,omitempty"`
	TxTotalCollision        string `json:"tx-total-collision,omitempty"`
}

func (t *InterfaceEthernet) GetRestPath() string {
	return "/interface/ethernet"
}

type InterfaceEthernets []InterfaceEthernet

func (t *InterfaceEthernets) GetRestPath() string {
	return "/interface/ethernet"
}

type InterfaceGRE struct {
	ID            string `json:".id,omitempty"`
	ActualMtu     string `json:"actual-mtu,omitempty"`
	AllowFastPath string `json:"allow-fast-path,omitempty"`
	ClampTCPMss   string `json:"clamp-tcp-mss,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Disabled      string `json:"disabled,omitempty"`
	DontFragment  string `json:"dont-fragment,omitempty"`
	Dscp          string `json:"dscp,omitempty"`
	LocalAddress  string `json:"local-address,omitempty"`
	Mtu           string `json:"mtu,omitempty"`
	Name          string `json:"name,omitempty"`
	RemoteAddress string `json:"remote-address,omitempty"`
	Running       string `json:"running,omitempty"`
}

func (t *InterfaceGRE) GetRestPath() string {
	return "/interface/gre"
}

type InterfaceGREs []InterfaceGRE

func (t *InterfaceGREs) GetRestPath() string {
	return "/interface/gre"
}

type InterfacePPPoEClient struct {
	ID                   string `json:".id,omitempty"`
	AcName               string `json:"ac-name,omitempty"`
	AddDefaultRoute      string `json:"add-default-route,omitempty"`
	Allow                string `json:"allow,omitempty"`
	Comment              string `json:"comment,omitempty"`
	DefaultRouteDistance string `json:"default-route-distance,omitempty"`
	DialOnDemand         string `json:"dial-on-demand,omitempty"`
	Disabled             string `json:"disabled,omitempty"`
	Interface            string `json:"interface,omitempty"`
	Invalid              string `json:"invalid,omitempty"`
	KeepaliveTimeout     string `json:"keepalive-timeout,omitempty"`
	MaxMru               string `json:"max-mru,omitempty"`
	MaxMtu               string `json:"max-mtu,omitempty"`
	Mrru                 string `json:"mrru,omitempty"`
	Name                 string `json:"name,omitempty"`
	Password             string `json:"password,omitempty"`
	Profile              string `json:"profile,omitempty"`
	Running              string `json:"running,omitempty"`
	ServiceName          string `json:"service-name,omitempty"`
	UsePeerDNS           string `json:"use-peer-dns,omitempty"`
	User                 string `json:"user,omitempty"`
}

func (t *InterfacePPPoEClient) GetRestPath() string {
	return "/interface/gre"
}

type InterfacePPPoEClients []InterfacePPPoEClient

func (t *InterfacePPPoEClients) GetRestPath() string {
	return "/interface/gre"
}

type InterfaceVLAN struct {
	ID                      string `json:".id,omitempty"`
	Arp                     string `json:"arp,omitempty"`
	ArpTimeout              string `json:"arp-timeout,omitempty"`
	Comment                 string `json:"comment,omitempty"`
	Disabled                string `json:"disabled,omitempty"`
	Interface               string `json:"interface,omitempty"`
	L2Mtu                   string `json:"l2mtu,omitempty"`
	LoopProtect             string `json:"loop-protect,omitempty"`
	LoopProtectDisableTime  string `json:"loop-protect-disable-time,omitempty"`
	LoopProtectSendInterval string `json:"loop-protect-send-interval,omitempty"`
	LoopProtectStatus       string `json:"loop-protect-status,omitempty"`
	MacAddress              string `json:"mac-address,omitempty"`
	Mtu                     string `json:"mtu,omitempty"`
	Name                    string `json:"name,omitempty"`
	Running                 string `json:"running,omitempty"`
	UseServiceTag           string `json:"use-service-tag,omitempty"`
	VlanID                  string `json:"vlan-id,omitempty"`
}

func (t *InterfaceVLAN) GetRestPath() string {
	return "/interface/vlan"
}

type InterfaceVLANs []InterfaceVLAN

func (t *InterfaceVLANs) GetRestPath() string {
	return "/interface/vlan"
}

type InterfaceWireguard struct {
	ID         string `json:".id,omitempty"`
	Disabled   string `json:"disabled,omitempty"`
	ListenPort string `json:"listen-port,omitempty"`
	Comment    string `json:"comment,omitempty"`
	Mtu        string `json:"mtu,omitempty"`
	Name       string `json:"name,omitempty"`
	PrivateKey string `json:"private-key,omitempty"`
	PublicKey  string `json:"public-key,omitempty"`
	Running    string `json:"running,omitempty"`
}

func (t *InterfaceWireguard) GetRestPath() string {
	return "/interface/wireguard"
}

type InterfaceWireguards []InterfaceWireguard

func (t *InterfaceWireguards) GetRestPath() string {
	return "/interface/wireguard"
}

type InterfaceWireguardPeer struct {
	ID                     string `json:".id,omitempty"`
	AllowedAddress         string `json:"allowed-address,omitempty"`
	Comment                string `json:"comment,omitempty"`
	CurrentEndpointAddress string `json:"current-endpoint-address,omitempty"`
	CurrentEndpointPort    string `json:"current-endpoint-port,omitempty"`
	Disabled               string `json:"disabled,omitempty"`
	EndpointAddress        string `json:"endpoint-address,omitempty"`
	EndpointPort           string `json:"endpoint-port,omitempty"`
	Interface              string `json:"interface,omitempty"`
	LastHandshake          string `json:"last-handshake,omitempty"`
	PersistentKeepalive    string `json:"persistent-keepalive,omitempty"`
	PublicKey              string `json:"public-key,omitempty"`
	Rx                     string `json:"rx,omitempty"`
	Tx                     string `json:"tx,omitempty"`
}

func (t *InterfaceWireguardPeer) GetRestPath() string {
	return "/interface/wireguard/peers"
}

type InterfaceWireguardPeers []InterfaceWireguardPeer

func (t *InterfaceWireguardPeers) GetRestPath() string {
	return "/interface/wireguard/peers"
}
