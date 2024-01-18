// Package xconfig 解析配置文件.
package xconfig

import (
	"os"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

// 系统运行环境以及该环境下默认配置文件路径.
const (
	// RuntimeENVWindows Windows系统运行环境
	RuntimeENVWindows = "windows"
	// RuntimeENVLinux Linux系统运行环境.
	RuntimeENVLinux = "linux"
	// RuntimeENVWindowsGlobalConfigPath Windows全局配置文件路径.
	RuntimeENVWindowsGlobalConfigPath = "F:\\code\\config\\goconfig"
	// RuntimeENVLinuxGlobalConfigPath Linux全局配置文件路径.
	RuntimeENVLinuxGlobalConfigPath = "/data/config"
)

// ParseIPXDBConfigPath 解析IP定位配置文件.
var ParseIPXDBConfigPath = ""

const defaultConfigFileType = "yaml"

var (
	configInstance     *Config
	configInstanceOnce sync.Once

	// runtimeOSFileSplit 运行环境文件分隔符.
	runtimeOSFileSplit = ""
)

func init() {
	switch runtime.GOOS {
	case RuntimeENVWindows:
		runtimeOSFileSplit = "\\"
		ParseIPXDBConfigPath = RuntimeENVWindowsGlobalConfigPath + "\\xdb\\ip2region.xdb"
	case RuntimeENVLinux:
		runtimeOSFileSplit = "/"
		ParseIPXDBConfigPath = RuntimeENVLinuxGlobalConfigPath + "/xdb/ip2region.xdb"
	default:
		runtimeOSFileSplit = "\\"
	}
}

// Config 对外暴露config.
type Config struct {
	config *viper.Viper
}

// InitConfig 初始化项目配置文件.
func InitConfig() *Config {
	configInstanceOnce.Do(func() {
		configInstance = &Config{}
		configInstance.initGlobalConfig()
	})
	return configInstance
}

func (c *Config) runtimeOSConfigPath() string {
	os := runtime.GOOS
	switch os {
	case RuntimeENVWindows:
		return RuntimeENVWindowsGlobalConfigPath
	case RuntimeENVLinux:
		return RuntimeENVLinuxGlobalConfigPath
	default:
		panic("unsupport runtime os:" + os)
	}
}

// Config 获取配置文件实例.
func (c *Config) Config() *viper.Viper {
	return c.config
}

// initGlobalConfig 初始化去哪聚配置文件.
func (c *Config) initGlobalConfig() {
	c.config = viper.New()
	c.config.SetConfigName("")
	c.config.SetConfigType(defaultConfigFileType)
	c.readAllConfig()
}

func (c *Config) readAllConfig() {
	configPath := c.runtimeOSConfigPath()
	file, err := os.Open(configPath)
	if err != nil {
		panic("open config path err:" + err.Error())
	}
	configs, err := file.Readdir(-1)
	if err != nil {
		panic("read config dir err:" + err.Error())
	}
	for _, fileInfo := range configs {
		if fileInfo.IsDir() {
			continue
		}
		path := configPath + runtimeOSFileSplit + fileInfo.Name()
		configFile, err := os.Open(path)
		if err != nil {
			panic("open config err:" + err.Error())
		}
		if err := c.config.ReadConfig(configFile); err != nil {
			panic("read config err:" + err.Error())
		}
	}
}
