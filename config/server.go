package config

import "github.com/spf13/viper"

type server struct {
	Name string //服务名称
	Mode string //运行模式：开发dev，测试test，生产环境prod
}

func (s *server) readConf() {
	s.Name = viper.GetString("name")
	s.Mode = viper.GetString("mode")
}

var Server = &server{}
