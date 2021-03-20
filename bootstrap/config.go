package bootstrap

import (
	"strconv"
)

type Config struct {
	Server `yaml:"server"`
}

type Server struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Router string `yaml:"router"`
	Method string `yaml:"method" json:"method"`
}

func (c *Server) Addr() string {
	return c.Host + ":" + strconv.Itoa(c.Port) + c.Router
}
