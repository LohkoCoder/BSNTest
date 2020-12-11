package main

import (
	"BSNTest/service"
	"BSNTest/src"
	"fmt"
	"log"
	"time"
)

const (
	api = "http://52.83.209.158:17502" //Node gateway address
	userCode = "USER0001202011301420587001697" //User No.
	appCode = "app001b9317f4b112c4ea781cf18481e3f338e" //Application No.
)

var (
	puk = src.ReadTxtFromFile("./BsnTestnetCert/userAppCert/secp256r1/public_key.pem") //Application public key
	prk = src.ReadTxtFromFile("./BsnTestnetCert/userAppCert/secp256r1/private_key.pem") //Application private key
	mspDir= "./BsnTestnetCert/fabricMsp" //Certificate storage directory
	cert = src.ReadTxtFromFile("./BsnTestnetCert/gatewayCert/gateway_public_cert_secp256r1.pem") //Certificate
	channelName = "channel202010310000001"
	chainCodeName = "cc_4a4d88419caa4ca5806f4266cf15a820"
)


func main()  {

	clientConfig := src.ClientConfig{
		Api:      api,
		UserCode: userCode,
		AppCode:  appCode,
		Puk:      puk,
		Prk:      prk,
		MspDir:   mspDir,
		Cert:     cert,
	}

	client, err := clientConfig.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	smartContract := service.SmartContract{
		ChainCodeName: chainCodeName,
		Client: client,
	}

	res := smartContract.QueryOrder("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")

	res = smartContract.RegisterOrder("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")
	time.Sleep(1000 * time.Millisecond)
	res = smartContract.ConfirmOrder("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")
	time.Sleep(1000 * time.Millisecond)
	res = smartContract.DownPayment("3201214","32012121", "zjj", "123", "万科", "天宇府", "89", "802")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")
	time.Sleep(1000 * time.Millisecond)
	res = smartContract.ConfirmDownPayment("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")
	time.Sleep(1000 * time.Millisecond)
	res = smartContract.FullPayment("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------waiting the tx to be mined--------")
	time.Sleep(1000 * time.Millisecond)
	res = smartContract.ConfirmFullPayment("3201214","32012121", "zjj", "123")
	fmt.Println(res)
	fmt.Println("--------done!!!!!!!!!!!!!!!!!!!!!!--------")
}

