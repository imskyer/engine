package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/plugin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"path/filepath"
)

func init() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetDefault("network", "tcp")
	viper.SetDefault("address", ":3000")
	if err := viper.ReadInConfig(); err != nil {
		grpclog.Warningln(err.Error())
	} else {
		grpclog.Infof("using config file: %s\n", viper.ConfigFileUsed())
	}
}

//
//
/*
Defaults:
network: "tcp"
address: ":3000"
paths: ""

Config contains configurations of gRPC and Gateway server. A new instance of Config is created from your config.yaml|config.json file in your current working directory
Network, Address, and Paths can be set in your config file to set the Config instance. Otherwise, defaults are set.
*/

type Config struct {
	Network            string   `mapstructure:"network" json:"network"`
	Address            string   `mapstructure:"address" json:"address"`
	Paths              []string `mapstructure:"paths" json:"paths"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Option             []grpc.ServerOption
}

// New creates a config from your config file. If no config file is present, the resulting Config will have the following defaults: netowork: "tcp" address: ":3000"
// use the With method to continue to modify the resulting Config object
func New() *Config {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		grpclog.Fatal(err.Error())
	}
	cfg.Plugins = plugin.LoadPlugins()
	return cfg
}

//CreateListener creates a network listener for the grpc server from the netowork address
func (c *Config) CreateListener() (net.Listener, error) {
	if c.Network == "unix" {
		dir := filepath.Dir(c.Address)
		f, err := os.Stat(dir)
		if err != nil {
			if err = os.MkdirAll(dir, 0755); err != nil {
				return nil, errors.Wrap(err, "failed to create the directory")
			}
		} else if !f.IsDir() {
			return nil, errors.Errorf("file %q already exists", dir)
		}
	}
	lis, err := net.Listen(c.Network, c.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s %s", c.Network, c.Address)
	}
	return lis, nil
}
