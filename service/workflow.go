package service

import (
	"BSNTest/src"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	"log"
)

type SmartContract struct {
	ChainCodeName string
	Client *fabric.FabricClient
}



func (smartContract *SmartContract) setReqBody(funcName string, args []string) (body node.TransReqDataBody) {
	body = node.TransReqDataBody{
		UserName:     "zjj",
		Nonce:        src.GetNonce(),
		ChainCode:    smartContract.ChainCodeName,
		FuncName:     funcName,
		Args:         args,
		TransientMap: nil,
	}
	return
}
func (smartContract *SmartContract) baseFunc(funcName, orderId, customerId, name, phoneNum string, extension ...string) string {

	body := smartContract.setReqBody(funcName, append([]string{orderId, customerId, name, phoneNum}, extension...))
	resData, err := smartContract.Client.SdkTran(body)
	if err !=nil {
		log.Fatal(err)
	}

	if resData.Header.Code != 0 || resData.Header.Msg != "success" || resData.Body.CCRes.CCCode != 200 {
		log.Fatal(err)
	}
	return resData.Body.CCRes.CCData
}

func (smartContract *SmartContract) QueryOrder(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("QueryOrder", orderId, customerId, name, phoneNum)
}

func (smartContract *SmartContract) RegisterOrder(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("RegisterOrder", orderId, customerId, name, phoneNum)
}

func (smartContract *SmartContract) ConfirmOrder(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("ConfirmOrder", orderId, customerId, name, phoneNum)
}

func (smartContract *SmartContract) DownPayment(orderId, customerId, name, phoneNum, estateOrg, neighborhood, buildingId, rootId string) string {
	return smartContract.baseFunc("DownPayment", orderId, customerId, name, phoneNum, estateOrg, neighborhood, buildingId, rootId)
}

func (smartContract *SmartContract) ConfirmDownPayment(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("ConfirmDownPayment", orderId, customerId, name, phoneNum)
}

func (smartContract *SmartContract) FullPayment(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("FullPayment", orderId, customerId, name, phoneNum)
}

func (smartContract *SmartContract) ConfirmFullPayment(orderId, customerId, name, phoneNum string) string {
	return smartContract.baseFunc("ConfirmFullPayment", orderId, customerId, name, phoneNum)
}
