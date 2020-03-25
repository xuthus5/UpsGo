package utils

// 该处定义了如何从根目录下获取并解析配置信息

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

// E 定义了读取配置文件信息的根结构
type E struct {
	Environments `yaml:"environments"`
}

// Environments 项目主要配置项[子项] 如果需要扩展 在这里添加结构来实现yaml的解析
type Environments struct {
	ProjectName string `yaml:"project_name"` //项目名称
	Debug       bool   `yaml:"debug"`        //是否开启debug模式
	Server      string `yaml:"server"`       //服务运行的host:port
	User        User   `yaml:"user"`         //用户相关配置
	Ups         Ups    `yaml:"ups"`          //ups服务相关
}

// User 用户配置项
type User struct {
	Username string `yaml:"username"` //用户名
	Password string `yaml:"password"` //密码
	Token    string `yaml:"token"`    //token验证
}

// File 文件相关配置
type Ups struct {
	Bucket   string `yaml:"Bucket"`   //服务名称
	Operator string `yaml:"Operator"` //权的操作员名称
	Password string `yaml:"Password"` //授权的操作员密码
}

// conf 是一个全局的配置信息实例 项目运行只读取一次 是一个单例
var conf *E
var once sync.Once

// GetConfig 调用该方法会实例化conf 项目运行会读取一次配置文件 确保不会有多余的读取损耗
func GetConfig() *E {
	once.Do(func() {
		conf = new(E)
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			//读取配置文件失败,停止执行
			panic("read config file error:" + err.Error())
		}
	})
	return conf
}
