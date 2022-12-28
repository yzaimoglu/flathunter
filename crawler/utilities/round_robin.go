package utilities

import (
	"errors"
	"sync/atomic"

	models "github.com/yzaimoglu/flathunter/models/crawler"
)

type RoundRobin interface {
  Next() *models.UserAgent
}

type roundrobin struct {
  user_agents []*models.UserAgent
  next uint32
}

func New(user_agents ...*models.UserAgent) (RoundRobin, error) {
  if len(user_agents) == 0 {
    return nil, errors.New("No User Agent found")
  }

  return &roundrobin{
    user_agents: user_agents,
  }, nil
}

func (r *roundrobin) Next() *models.UserAgent {
  n := atomic.AddUint32(&r.next, 1)
  return r.user_agents[(int(n)-1)%len(r.user_agents)]
}
