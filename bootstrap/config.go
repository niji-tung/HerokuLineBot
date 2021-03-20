package bootstrap

import (
	"strconv"
)

type Config struct {
	Server       `yaml:"server"`
	LineBot      `yaml:"line_bot"`
	GoogleScript `yaml:"google_script"`
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

type LineBot struct {
	AdminID            string `yaml:"admin_id"`
	RoomID             string `yaml:"room_id"`
	ChannelAccessToken string `yaml:"channel_access_token"`
}

type GoogleScript struct {
	Url string `yaml:"url"`
}
