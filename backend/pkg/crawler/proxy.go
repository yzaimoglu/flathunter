package crawler

import (
	"strconv"

	"github.com/yzaimoglu/flathunter/pkg/models"
)

func ProxyString(proxy *models.Proxy) (proxy_string string) {
	return ("socks5://" + proxy.Username + ":" + proxy.Password + "@" + proxy.IP + ":" + strconv.Itoa(proxy.Port))
}
