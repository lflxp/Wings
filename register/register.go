package register

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/lflxp/Wings/utils"
	"github.com/lflxp/curl"
	"github.com/lflxp/dbui/etcd"
)

//默认注册地址
var serviceUrl string = "/ams/main/services"
var args map[string]string = map[string]string{"name": "", "address": "", "monitor::port": "", "rpc::port": "", "server": ""}

func Go() {
	err := Register("conf/app.conf")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
}

func WatchDog() {
	Go()
	gcInterval, _ := time.ParseDuration("29s")
	ticker := time.NewTicker(gcInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				Go()
			}
		}
	}()
}

func GetConfig(path string) error {
	iniConfig, err := config.NewConfig("ini", path)
	if err != nil {
		return err
	}
	//验证中控机是否可访问 后期会改为api接口网关
	server := iniConfig.String("server")
	if strings.ContainsAny(server, ":") == false {
		server = fmt.Sprintf("%s:80", server)
	}
	if utils.ScannerPort(server) == false {
		return errors.New(fmt.Sprintf("中控机 %s 不可达", server))
	}
	for key, _ := range args {
		args[key] = iniConfig.String(key)
	}
	return nil
}

//注册rpc和api接口
func Register(path string) error {
	err := GetConfig(path)
	if err != nil {
		return err
	}

	//etcd 连接
	//etcd 服务器地址由中控机提供
	st := &etcd.EtcdUi{Endpoints: []string{curl.HttpGet(fmt.Sprintf("http://%s/api/v1/etcdhost", args["server"]))}}
	st.InitClientConn()
	//验证ip
	resp := st.Get("/ams/main/ansible/ip")
	ip := curl.HttpsGet(string(resp.Kvs[0].Value))
	if args["address"] != ip {
		fmt.Println(args["address"], string(resp.Kvs[0].Value))
		return errors.New(fmt.Sprintf("配置IP与真实IP不符 %s %s", args["address"], ip))
	}
	//注册服务
	err = st.Add(fmt.Sprintf("%s/%s", serviceUrl, args["name"]), args["name"])
	err = st.Add(fmt.Sprintf("%s/%s/rpc", serviceUrl, args["name"]), fmt.Sprintf("%s %s", args["name"], "rpc服务调用"))
	err = st.Add(fmt.Sprintf("%s/%s/api", serviceUrl, args["name"]), fmt.Sprintf("%s %s", args["name"], "服务调用(监控和自身)"))
	//注册rpc服务
	err = st.AddLease(fmt.Sprintf("%s/%s/rpc/tcp@%s:%s", serviceUrl, args["name"], args["address"], args["rpc::port"]), "test", 30)
	if err != nil {
		return err
	}
	//注册api监控服务
	err = st.AddLease(fmt.Sprintf("%s/%s/api/tcp@%s:%s", serviceUrl, args["name"], args["address"], args["monitor::port"]), fmt.Sprintf("%s:%s", args["address"], args["monitor:port"]), 30)
	if err != nil {
		return err
	}
	return nil
}
