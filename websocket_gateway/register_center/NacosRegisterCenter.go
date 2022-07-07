package register_center

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/*
	Nacos 注册中心
	TODO 这个类还需要斟酌一下配置项抽取出来进一步封装
*/

type Nacos struct {
	Client naming_client.INamingClient
}

var NacosClient Nacos = Nacos{}

func GetNacosClient() *Nacos {
	return &NacosClient
}

func (n *Nacos) RegisterInstance() (bool, error) {
	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("D:\\PROJECT\\go_project\\im_go\\log"),
		constant.WithCacheDir("D:\\PROJECT\\go_project\\im_go\\cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),
		constant.WithPassword("admin123"),
	)

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			Scheme:      "http",
			IpAddr:      "nas.huerpu.top",
			Port:        28848,
			ContextPath: "/nacos",
		},
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(" 创建Nacos 服务失败, 失败信息 :" + err.Error())
		return false, err
	}

	_, errR := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        8080,
		ServiceName: "websocket_gateway_1",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "beijing"},
		//ClusterName: "cluster-a", // 默认值DEFAULT
		//GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})

	if errR != nil {
		panic(" 注册实例失败，异常信息：" + errR.Error())
		return false, errR
	}

	n.Client = namingClient
	return false, nil
}

func (n *Nacos) SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error) {
	instances, err := n.Client.SelectInstances(param)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (n *Nacos) SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	healthyInstance, err := n.Client.SelectOneHealthyInstance(param)
	if err != nil {
		return nil, err
	}
	return healthyInstance, nil
}
