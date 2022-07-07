package iface

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/*
	注册中心所需要的方法命名  ==> 暂时使用Nacos相关配置，如果后续需要支持其他注册中心，需要修改该接口文件方法
	- 1、注册服务
	- 2、获取全部服务实例
 	- 3、获取一个负载均衡服务实例
	- 4、订阅服务变化 -- 暂不需要
*/

type IRegisterCenter interface {
	RegisterInstance() (bool, error)

	SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error)

	SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error)
}
