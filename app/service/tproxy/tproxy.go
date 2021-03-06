package tproxy

type ITProxy interface {
	AddRoute() error
	DelRoute() error
	EnableProxy() error
	DisableProxy() error
}

var Tproxy = makeITProxy()

type IProxy interface {
	ApplyProxy()
}

var Proxy = makeIProxy()

const (
	tableId   int = 100
	proxyPort int = 65535
	markId    int = 1
)

var (
	intranet  []string = []string{"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16", "172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.88.99.0/24", "192.168.0.0/16", "198.18.0.0/15", "198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/4", "240.0.0.0/4", "255.255.255.255/32"}
	intranet6 []string = []string{"::/128", "::1/128", "::ffff:0:0/96", "100::/64", "64:ff9b::/96", "2001::/32", "2001:10::/28", "2001:20::/28", "2001:db8::/32", "2002::/16", "fc00::/7", "fe80::/10", "ff00::/8"}
)
