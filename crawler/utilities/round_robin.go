package utilities

import (
	"errors"
	"sync/atomic"

	models "github.com/yzaimoglu/flathunter/models/crawler"
)

type RoundRobinUA interface {
  Next() *models.UserAgent
}

type RoundRobinProxy interface {
  Next() *models.Proxy
}

type roundrobin_proxy struct {
  proxies []*models.Proxy
  next uint32
}

type roundrobin_ua struct {
  user_agents []*models.UserAgent
  next uint32
}

func NewProxy(proxies ...*models.Proxy) (RoundRobinProxy, error) {
  if len(proxies) == 0 {
    return nil, errors.New("No Proxy found")
  }

  return &roundrobin_proxy{
    proxies: proxies,
  }, nil
}

func NewUserAgent(user_agents ...*models.UserAgent) (RoundRobinUA, error) {
  if len(user_agents) == 0 {
    return nil, errors.New("No User Agent found")
  }

  return &roundrobin_ua{
    user_agents: user_agents,
  }, nil
}

func (r *roundrobin_proxy) Next() *models.Proxy {
  n := atomic.AddUint32(&r.next, 1)
  return r.proxies[(int(n)-1)%len(r.proxies)]
}

func (r *roundrobin_ua) Next() *models.UserAgent {
  n := atomic.AddUint32(&r.next, 1)
  return r.user_agents[(int(n)-1)%len(r.user_agents)]
}
