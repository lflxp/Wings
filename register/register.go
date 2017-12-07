package register

import (
	"fmt"
	"errors"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/curl"
	"github.com/astaxie/beego/config"
)

//默认注册地址
var serviceUrl string = "/ams/main/services"
var args map[string]string = map[string]string{"name":"name","address":"address","monitor::port":"monitor::port","rpc::port":"rpc::port","etcd::host":"etcd::host"}

func GetConfig(path string) error {
	iniConfig,err := config.NewConfig("ini",path)
	if err != nil {
		return err
	}
	for _,key := range args {
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
	st := &etcd.EtcdUi{Endpoints:[]string{args["etcd::host"]}}
	st.InitClientConn()
	//验证ip
	resp := st.Get("/ams/main/ansible/ip")
	ip := curl.HttpsGet(string(resp.Kvs[0].Value))
	if args["address"] != ip {
		fmt.Println(args["address"],string(resp.Kvs[0].Value))
		return errors.New(fmt.Sprintf("配置IP与真实IP不符 %s %s",args["address"],ip))
	}
	//注册服务
    err = st.Add(fmt.Sprintf("%s/%s",serviceUrl,args["name"]),args["name"])	
    err = st.Add(fmt.Sprintf("%s/%s/rpc",serviceUrl,args["name"]),fmt.Sprintf("%s %s",args["name"],"rpc服务调用"))	
    err = st.Add(fmt.Sprintf("%s/%s/api",serviceUrl,args["name"]),fmt.Sprintf("%s %s",args["name"],"服务调用(监控和自身)"))	
	//注册rpc服务
	err = st.Add(fmt.Sprintf("%s/%s/rpc/tcp@%s:%s",serviceUrl,args["name"],args["address"],args["rpc::port"]),"test")
	if err != nil {
		return err
	}
	//注册api监控服务
	err = st.Add(fmt.Sprintf("%s/%s/api/tcp@%s:%s",serviceUrl,args["name"],args["address"],args["monitor::port"]),fmt.Sprintf("%s:%s",args["address"],args["monitor:port"]))
	if err != nil {
		return err
	}
	return nil
}