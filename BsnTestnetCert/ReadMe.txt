Certificate download zip includes two parts:  off-chain application certificate and gateway certificate：
1、userAppCert directory: used to store off-chain application certificates, the files are as follows:
【1】secp256r1 directory：Use the directory certificate for signing and validation If the algorithm type of your published test network service is secp256r1.
     【1】public_cert.pem：The public key file of the application certificate is mainly used by the gateway to verify the received signature data or disclose it to other users.
     【2】private_key.pem：The private key file of the application certificate is mainly used for signing and verifing data when requesting the gateway. 
【2】secp256k1directory：Use the directory certificate for signing and validation If the algorithm type of your published test network service is secp256k1.
        【1】public_cert.pem：The public key file of the application certificate is mainly used by the gateway to verify the received signature data or disclose it to other users.。
        【2】private_key.pem：The private key file of the application certificate is mainly used for signing and verifing data when requesting the gateway.
【3】sm2 directory：Use the directory certificate for signing and validation If the algorithm type of your published test network service is sm2.
        【1】public_cert.pem：The public key file of the application certificate is mainly used by the gateway to verify the received signature data or disclose it to other users.
        【2】private_key.pem：The private key file of the application certificate is mainly used for signing and verifing data when requesting the gateway.
2、The gatewayCert directory: used to store the gateway certificate, the specific files are as follows:
     【1】gateway_public_cert_secp256r1.pem：The gateway public key (secp256r1) is mainly used to verify the data when the gateway responds if you are participating in a service algorithm of type SECP256R1.
     【2】gateway_public_cert_secp256k1.pem：The gateway public key (secp256k1) is mainly used to verify the data when the gateway responds if you are participating in a service algorithm of type SECP256K1.
     【3】gateway_public_cert_sm2.pem：The gateway public key(sm2) is mainly used to verify the data when the gateway responds if you are participating in a service algorithm of type SM2.
3、fabricMs directory：Upload a local certificate of public key upload mode for Fabric, the specific files are as follows:
    【1】keystore directory：The private key of the user's certificate.
    【2】{testuser}@{AppCode}-cert.pem：The user certificate file for Fabric。
    【Note】How to use: Change the {testuser}@{AppCode}-cert.pem file in fabricMsp to【child user name@AppCode-cert.pem】，and configure the folder path to the mspDir parameter in config.
4、This type of certificate file is only used for the node gateway interactive signature and data verification, not for any application business processing! 

------------------------------------------------------------------------
证书下载压缩包含用户链下应用证书和网关证书两大部分：
1、userAppCert目录：用于存放链下应用证书，具体文件如下：
     【1】secp256r1目录：如果您发布测试网服务算法类型为secp256r1时,使用该目录证书进行签名以及验证。
        【1】public_cert.pem：应用证书公钥文件，主要用于网关对接受的签名数据进行验证或公开给其它使用者。
        【2】private_key.pem：应用证书私钥文件，主要用于请求网关时对数据进行签名以及验证。
     【2】secp256k1目录：如果您发布测试网服务算法类型为secp256k1时,使用该目录证书进行签名。
        【1】public_cert.pem：应用证书公钥文件，主要用于网关对接受的签名数据进行验证或公开给其它使用者。
        【2】private_key.pem：应用证书私钥文件，主要用于请求网关时对数据进行签名。
     【3】sm2目录：如果您发布测试网服务算法类型为国密(sm2)时,使用该目录证书进行签名以及验证。
        【1】public_cert.pem：应用证书公钥文件，主要用于网关对接受的签名数据进行验证或公开给其它使用者。
        【2】private_key.pem：应用证书私钥文件，主要用于请求网关时对数据进行签名。
2、gatewayCert目录：用于存放网关证书，具体文件如下：
     【1】gateway_public_cert_secp256r1.pem：网关公钥secp256r1证书，如果您发布测试网服务算法类型为secp256r1时，用该公钥证书对网关响应时的数据进行验签。
     【2】gateway_public_cert_secp256k1.pem：网关公钥secp256k1证书，如果您发布测试网服务算法类型为secp256k1时，用该公钥证书对网关响应时的数据进行验签。
     【3】gateway_public_cert_sm2.pem：网关公钥国密(sm2)证书，如果您发布测试网服务算法类型为国密(sm2)时，用该公钥证书对网关响应时的数据进行验签。
3、fabricMsp目录：为Fabric上传公钥模式中的本地证书，具体文件如下：
    【1】keystore目录：用户证书的私钥。
    【2】{testuser}@{AppCode}-cert.pem：Fabric的用户证书文件。
    【注】使用方式: 为将fabricMsp中的{testuser}@{AppCode}-cert.pem文件修改为【子用户名@AppCode-cert.pem】，并将该文件夹路径配置到config中的mspDir参数。
4、此类证书文件，只限于测试网与节点网关数据交互签名与验签使用，不用于任何应用业务处理！
