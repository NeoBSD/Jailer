package jailer

// Jail represents a single jail config
type Jail struct {
	JID                int    `json:"jid" jail:"jid"`
	Parent             int    `json:"parent" jail:"parent"`
	Name               string `json:"name,omitempty" jail:"name"`
	Hostname           string `json:"hostname,omitempty" jail:"host.hostname"`
	HostID             int    `json:"hostid" jail:"host.hostid"`
	HostUUD            string `json:"hostuuid,omitempty" jail:"host.hostuuid"`
	OSRelease          string `json:"osrelease,omitempty" jail:"osrelease"`
	OSReleaseDate      int    `json:"osreldate" jail:"osreldate"`
	DevFSRuleset       int    `json:"devfs_ruleset" jail:"devfs_ruleset"`
	EnforceStatFS      int    `json:"enforce_statfs" jail:"enforce_statfs"`
	SecureLevel        int    `json:"securelevel" jail:"securelevel"`
	ChildrenCurrent    int    `json:"children_current" jail:"children.cur"`
	ChildrenMax        int    `json:"children_max" jail:"children.max"`
	CPUSetID           int    `json:"cpuset_id" jail:"cpuset.id"`
	IP                 string `json:"ip,omitempty" jail:"ip4.addr"`
	Path               string `json:"path,omitempty" jail:"path"`
	MountDevFS         bool   `json:"mount_devfs,omitempty" jail:"mount_devfs"`
	ExecStart          string `json:"exec_start,omitempty" jail:"exec_start"`
	ExecStop           string `json:"exec_stop,omitempty" jail:"exec_stop"`
	AllowNochFlags     bool   `json:"allow_nochflags,omitempty" jail:"allow.nochflags"`
	AllowNoMLock       bool   `json:"allow_nomlock,omitempty" jail:"allow.nomlock"`
	AllowNoMount       bool   `json:"allow_nomount,omitempty" jail:"allow.nomount"`
	AllowMountNoDevFS  bool   `json:"allow_mount_nodevfs,omitempty" jail:"allow.mount.nodevfs"`
	AllowMountNoProcFS bool   `json:"allow_mount_noprocfs,omitempty" jail:"allow.mount.noprocfs"`
	AllowMountNoTmpFS  bool   `json:"allow_mount_notmpfs,omitempty" jail:"allow.mount.notmpfs"`
	AllowMountNoZFS    bool   `json:"allow_mount_nozfs,omitempty" jail:"allow.mount.nozfs"`
	AllowNoQuotas      bool   `json:"allow_noquotas,omitempty" jail:"allow.noquotas"`
	AllowNoRawSockets  bool   `json:"allow_noraw_sockets,omitempty" jail:"allow.noraw_sockets"`
	AllowNoReadMsgBuf  bool   `json:"allow_noread_msgbuf,omitempty" jail:"allow.noread_msgbuf"`
	AllowReservedPorts bool   `json:"allow_reserved_ports,omitempty" jail:"allow.reserved_ports"`
	AllowSetHostname   bool   `json:"allow_set_hostname,omitempty" jail:"allow.set_hostname"`
	AllowNoSocketAF    bool   `json:"allow_nosocket_af,omitempty" jail:"allow.nosocket_af"`
	AllowNoSysVIPC     bool   `json:"allow_nosysvipc,omitempty" jail:"allow.nosysvipc"`
}

// =1202000
// =0
// =2
// =-1
// =0
// children.max=0
// =2
// =0

// nodying
// host=new
// ip4=inherit
// ip6=inherit
//  nopersist
// sysvmsg=disable
// sysvsem=disable
// sysvshm=disable
// vnet=inherit
// host.domainname=/""}""
// =test_01.tobante.local
// =
// ip4.saddrsel
// ip6.addr=
// ip6.saddrsel
