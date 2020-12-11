package src

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"log"
)

type ClientConfig struct {
	Api string
	UserCode string
	AppCode string
	Puk string
	Prk string
	MspDir string
	Cert string
}

const chainCodeName = "cc_4a4d88419caa4ca5806f4266cf15a820"

func (cc *ClientConfig) InitConfig() (client *fabric.FabricClient, err error) {

	clientConfig, err := config.NewConfig(cc.Api, cc.UserCode, cc.AppCode, cc.Puk, cc.Prk, cc.MspDir, cc.Cert )
	if err !=nil{
		log.Fatal(err)
		return nil, err
	}
	err = clientConfig.Init()
	if err != nil{
		log.Fatal(err)
		return nil, err
	}

	client, err = fabric.InitFabricClient(clientConfig)
	if err != nil{
		log.Fatal(err)
		return nil, err
	}
	return
}


