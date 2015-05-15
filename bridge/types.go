//go:generate go-extpoints . AdapterFactory
package bridge

import (
	"net/url"

	dockerapi "github.com/fsouza/go-dockerclient"
)

type AdapterFactory interface {
	New(uri *url.URL) RegistryAdapter
}

type RegistryAdapter interface {
	Ping() error
	Register(service *Service) error
	Deregister(service *Service) error
	Refresh(service *Service) error
}

type Config struct {
	HostIp          string
	Internal        bool
	ForceTags       string
	RefreshTtl      int
	RefreshInterval int
	DeregisterCheck string
}

type Service struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	Port   int               `json:"port"`
	IP     string            `json:"ip"`
	Tags   []string          `json:"tags"`
	Attrs  map[string]string `json:"attrs"`
	TTL    int               `json:"ttl"`
	Origin ServicePort
}

type DeadContainer struct {
	TTL      int
	Services []*Service
}

type ServicePort struct {
	HostPort          string
	HostIP            string
	ExposedPort       string
	ExposedIP         string
	PortType          string
	ContainerHostname string
	ContainerID       string
	container         *dockerapi.Container
}

func (r ServicePort) GetContainer() *dockerapi.Container {
	return r.container
}
