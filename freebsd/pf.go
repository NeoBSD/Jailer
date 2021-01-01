package freebsd

import (
	"strconv"
	"strings"
	"time"
)

// States ...
type States int

// Timeouts ...
type Timeouts struct {
	TCPFirst       time.Duration `json:"tcp_first,omitempty"`
	TCPOpening     time.Duration `json:"tcp_opening,omitempty"`
	TCPEstablished time.Duration `json:"tcp_established,omitempty"`
	TCPClosing     time.Duration `json:"tcp_closing,omitempty"`
	TCPFinwait     time.Duration `json:"tcp_finwait,omitempty"`
	TCPClosed      time.Duration `json:"tcp_closed,omitempty"`
	TCPTSDiff      time.Duration `json:"tcp_tsdiff,omitempty"`
	UDPFirst       time.Duration `json:"udp_first,omitempty"`
	UDPSingle      time.Duration `json:"udp_single,omitempty"`
	UDPMultiple    time.Duration `json:"udp_multiple,omitempty"`
	ICMPFirst      time.Duration `json:"icmp_first,omitempty"`
	ICMPError      time.Duration `json:"icmp_error,omitempty"`
	OtherFirst     time.Duration `json:"other_first,omitempty"`
	OtherSingle    time.Duration `json:"other_single,omitempty"`
	OtherMultiple  time.Duration `json:"other_multiple,omitempty"`
	Frag           time.Duration `json:"frag,omitempty"`
	Interval       time.Duration `json:"interval,omitempty"`
	AdaptiveStart  States        `json:"adaptive_start,omitempty"`
	AdaptiveEnd    States        `json:"adaptive_end,omitempty"`
	SrcTrack       time.Duration `json:"src_track,omitempty"`
}

// PFCTL ...
type PFCTL struct {
	Timeouts Timeouts `json:"timeouts,omitempty"`
}

// GetPFCTL ...
func GetPFCTL() (PFCTL, error) {
	out, err := getCommandOutput("sudo", "pfctl", "-s", "timeouts")
	if err != nil {
		return PFCTL{}, err
	}

	pfctl := PFCTL{}
	timeValues := []struct {
		ID    string
		Value *time.Duration
	}{
		{ID: "tcp.first", Value: &pfctl.Timeouts.TCPFirst},
		{ID: "tcp.opening", Value: &pfctl.Timeouts.TCPOpening},
		{ID: "tcp.established", Value: &pfctl.Timeouts.TCPEstablished},
		{ID: "tcp.closing", Value: &pfctl.Timeouts.TCPClosing},
		{ID: "tcp.finwait", Value: &pfctl.Timeouts.TCPFinwait},
		{ID: "tcp.closed", Value: &pfctl.Timeouts.TCPClosed},
		{ID: "tcp.tsdiff", Value: &pfctl.Timeouts.TCPTSDiff},
		{ID: "udp.first", Value: &pfctl.Timeouts.UDPFirst},
		{ID: "udp.single", Value: &pfctl.Timeouts.UDPSingle},
		{ID: "udp.multiple", Value: &pfctl.Timeouts.UDPMultiple},
		{ID: "icmp.first", Value: &pfctl.Timeouts.ICMPFirst},
		{ID: "icmp.error", Value: &pfctl.Timeouts.ICMPError},
		{ID: "other.first", Value: &pfctl.Timeouts.OtherFirst},
		{ID: "other.single", Value: &pfctl.Timeouts.OtherSingle},
		{ID: "other.multiple", Value: &pfctl.Timeouts.OtherMultiple},
		{ID: "frag", Value: &pfctl.Timeouts.Frag},
		{ID: "interval", Value: &pfctl.Timeouts.Interval},
		{ID: "src.track", Value: &pfctl.Timeouts.SrcTrack},
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		for _, tag := range timeValues {
			if strings.HasPrefix(line, tag.ID) {
				sections := strings.SplitAfter(line, tag.ID)
				secondsStr := trim(sections[1][:len(sections[1])-1], "")
				num, err := strconv.Atoi(secondsStr)
				if err != nil {
					return PFCTL{}, err
				}
				*tag.Value = time.Duration(num) * time.Second
			}
		}

		if strings.HasPrefix(line, "adaptive.start") {
			sections := strings.SplitAfter(line, "adaptive.start")
			secondsStr := trim(sections[1][:len(sections[1])-6], "")
			num, err := strconv.Atoi(secondsStr)
			if err != nil {
				return PFCTL{}, err
			}
			pfctl.Timeouts.AdaptiveStart = States(num)
		}

		if strings.HasPrefix(line, "adaptive.end") {
			sections := strings.SplitAfter(line, "adaptive.end")
			secondsStr := trim(sections[1][:len(sections[1])-6], "")
			num, err := strconv.Atoi(secondsStr)
			if err != nil {
				return PFCTL{}, err
			}
			pfctl.Timeouts.AdaptiveEnd = States(num)
		}
	}

	return pfctl, nil
}
