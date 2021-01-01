package freebsd

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// INET ...
type INET struct {
	IP        string `json:"ip,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Broadcast string `json:"broadcast,omitempty"`
}

// INET6 ...
type INET6 struct {
	IP        string `json:"ip,omitempty"`
	Prefixlen string `json:"prefixlen,omitempty"`
	ScopeID   string `json:"scopeid,omitempty"`
}

// NetworkInterface ...
type NetworkInterface struct {
	Name    string  `json:"name"`
	Flags   string  `json:"flags"`
	Options string  `json:"options,omitempty"`
	Ether   string  `json:"ether,omitempty"`
	INET    []INET  `json:"inet,omitempty"`
	INET6   []INET6 `json:"inet6,omitempty"`
	Groups  string  `json:"groups,omitempty"`
	Media   string  `json:"media,omitempty"`
	Status  string  `json:"status,omitempty"`
	ND6     string  `json:"nd6,omitempty"`
}

// options=80009<RXCSUM,VLAN_MTU,LINKSTATE>
// ether b8:27:eb:89:fa:7a
// inet 10.23.0.103 netmask 0xffffff00 broadcast 10.23.0.255
// media: Ethernet autoselect (1000baseT <full-duplex,master>)
// status: active
// nd6 options=29<PERFORMNUD,IFDISABLED,AUTO_LINKLOCAL>

// GetNetworkInterfaces returns the freebsd-version command output
func GetNetworkInterfaces() ([]NetworkInterface, error) {
	output, err := getCommandOutput("ifconfig")
	if err != nil {
		return nil, err
	}

	interfaces := make([]NetworkInterface, 0)
	current := NetworkInterface{}
	lines := strings.Split(output, "\n")
	for n, line := range lines {
		r, _ := utf8.DecodeRuneInString(line[:1])
		if !unicode.IsSpace(r) {
			if n != 0 {
				interfaces = append(interfaces, current)
			}
			splits := strings.Split(line, ":")
			flags := strings.TrimPrefix(splits[1], " flags=")
			current = NetworkInterface{
				Name:  splits[0],
				Flags: flags,
			}

			continue
		}

		trimmed := trim(line, "")
		if strings.HasPrefix(trimmed, "options") {
			current.Options = trim(trimmed, "options=")
		}
		if strings.HasPrefix(trimmed, "ether") {
			current.Ether = trim(trimmed, "ether ")
		}
		if strings.HasPrefix(trimmed, "inet ") {
			parseINET(&current, trimmed)
		}
		if strings.HasPrefix(trimmed, "inet6 ") {
			parseINET6(&current, trimmed)
		}
		if strings.HasPrefix(trimmed, "media") {
			current.Media = trim(trimmed, "media:")
		}
		if strings.HasPrefix(trimmed, "status") {
			current.Status = trim(trimmed, "status:")
		}
		if strings.HasPrefix(trimmed, "groups") {
			current.Groups = trim(trimmed, "groups:")
		}
	}

	interfaces = append(interfaces, current)
	return interfaces, err
}

func trim(in, prefix string) string {
	spaces := " \n\r\t"
	return strings.Trim(strings.TrimPrefix(in, prefix), spaces)
}

func parseINET(ni *NetworkInterface, str string) {
	str = trim(str, "inet ")
	splits := strings.Split(str, " ")
	inet := INET{}
	for n, s := range splits {
		if n == 0 {
			inet.IP = trim(splits[0], "")
		}
		if s == "netmask" {
			inet.Netmask = trim(splits[n+1], "")
		}
		if s == "broadcast" {
			inet.Broadcast = trim(splits[n+1], "")
		}
	}
	if ni.INET != nil {
		ni.INET = append(ni.INET, inet)
	} else {
		ni.INET = []INET{inet}
	}
}

func parseINET6(ni *NetworkInterface, str string) {
	str = trim(str, "inet6 ")
	splits := strings.Split(str, " ")
	inet := INET6{}
	for n, s := range splits {
		if n == 0 {
			inet.IP = trim(splits[0], "")
		}
		if s == "prefixlen" {
			inet.Prefixlen = trim(splits[n+1], "")
		}
		if s == "scopeid" {
			inet.ScopeID = trim(splits[n+1], "")
		}
	}
	if ni.INET6 != nil {
		ni.INET6 = append(ni.INET6, inet)
	} else {
		ni.INET6 = []INET6{inet}
	}
}
