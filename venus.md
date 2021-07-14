
# 基于官方版本改的lotus-miner

变化功能点：
1. 发送消息到云上组件
2. 计算WinningPost证明
3. 禁用本地出块

位置https://github.com/ipfs-force-community/lotus/tree/force/venus_pool_v5

## configration change
```
[Venus]
  [Venus.Messager]
    Url = ""
    Token = ""
  [Venus.RegisterProofAPI]
    Url = ""
    Token = ""
  [Venus.Wallet]
    Url = ""
    Token = ""
```

## lotus-miner new api for push message in cmds

1. MessagerWaitMessage
2. MessagerPushMessage
3. MessagerGetMessage

lotus-miner提供的命令中需要发送，等待，搜索消息的命令全部通过调用lotus-miner新加的messager接口来实现

位置：

location：
cmd/lotus-storage-miner/*
api/api_storage.go
extern/storage-sealing/sealing.go
extern/storage-sealing/terminate_batch.go
node/modules/storageminer.go
storage/*

## add GasOverEstimate parameters to MessageSendSpec
位置：api/types.go
```go
type MessageSendSpec struct {
	MaxFee            abi.TokenAmount
	GasOverEstimation float64
}
```

## miner检查发送消息的地址是否存在
lotus-miner中原先需要判断地址存在的接口，需要改成调用messager的WalletHas接口
location: storage/addresses.go

## venus-messager client
venus-messager的客户端和lotus类似，都是使用的jsonrpc协议，可以同样使用lotus的jsonrpc库来接入。 需要注意的是命名空间是Messager
implement location： node/modules/messager/*

## connect gateway to receive ComputeProof request

这个功能用于监听云上组件下发的请求，在miner中是为了用于计算WinningPost证明
location：cli/wallet_client.go， node/impl/proof_client/*

## 非消息类型签名转到本地数据库（storageask， proposaldeal）
由于远程节点中不再存在私钥，一些不需要云上组件参与的签名（例如StorageAsk,DealProposal)需要连接钱包在本地进行签名
位置： markets/storageadapter/provider.go

## lotus-miner禁用本地出块
因为云上组件已经在出块，本地如果在出块会导致共识错误，因此，务必禁用本地出块
location: node/modules/storageminer.go 方法SetupBlockProducer