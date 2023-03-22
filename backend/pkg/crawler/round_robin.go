package crawler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// ProxyString returns the formatted proxy string.
func ProxyString(proxy *models.Proxy) (proxy_string string) {
	return ("socks5://" + proxy.Username + ":" + proxy.Password + "@" + proxy.IP + ":" + strconv.Itoa(proxy.Port))
}

// GetUserAgentRA returns a round robin object for the user agent.
func GetUserAgentRA() (RoundRobinUA, error) {
	var user_agents []*models.UserAgent
	readFile, err := os.Open("./assets/user_agents.crawl")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		user_agents = append(user_agents, &models.UserAgent{UserAgent: line})
	}

	return NewUserAgent(user_agents...)
}

// GetProxyRA returns a round robin object for the proxy.
func GetProxyRA() (RoundRobinProxy, error) {
	var proxies []*models.Proxy
	readFile, err := os.Open("./assets/proxies.crawl")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ":")
		ip := line[0]
		port, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			slog.Fatal(err)
		}
		username := line[2]
		password := line[3]
		proxies = append(proxies, &models.Proxy{IP: ip, Port: int(port), Username: username, Password: password})
	}
	readFile.Close()
	return NewProxy(proxies...)
}

// RoundRobinUA is the interface for the round robin user agent.
type RoundRobinUA interface {
	Next() *models.UserAgent
}

// RoundRobinProxy is the interface for the round robin proxy.
type RoundRobinProxy interface {
	Next() *models.Proxy
}

// roundrobin_proxy is the implementation of the round robin proxy.
type roundrobin_proxy struct {
	proxies []*models.Proxy
	next    uint32
}

// roundrobin_ua is the implementation of the round robin user agent.
type roundrobin_ua struct {
	user_agents []*models.UserAgent
	next        uint32
}

// NewProxy returns a new round robin proxy.
func NewProxy(proxies ...*models.Proxy) (RoundRobinProxy, error) {
	if len(proxies) == 0 {
		return nil, errors.New("no proxy found")
	}

	return &roundrobin_proxy{
		proxies: proxies,
	}, nil
}

// NewUserAgent returns a new round robin user agent.
func NewUserAgent(user_agents ...*models.UserAgent) (RoundRobinUA, error) {
	if len(user_agents) == 0 {
		return nil, errors.New("no user agent found")
	}

	return &roundrobin_ua{
		user_agents: user_agents,
	}, nil
}

// Next returns the next proxy.
func (r *roundrobin_proxy) Next() *models.Proxy {
	n := atomic.AddUint32(&r.next, 1)
	return r.proxies[(int(n)-1)%len(r.proxies)]
}

// Next returns the next user agent.
func (r *roundrobin_ua) Next() *models.UserAgent {
	n := atomic.AddUint32(&r.next, 1)
	return r.user_agents[(int(n)-1)%len(r.user_agents)]
}
