# Groups
* [](#)
  * [Closing](#Closing)
  * [Discover](#Discover)
  * [Session](#Session)
  * [Shutdown](#Shutdown)
  * [Version](#Version)
* [Auth](#Auth)
  * [AuthNew](#AuthNew)
  * [AuthVerify](#AuthVerify)
* [Beacon](#Beacon)
  * [BeaconGetEntry](#BeaconGetEntry)
* [Chain](#Chain)
  * [ChainDeleteObj](#ChainDeleteObj)
  * [ChainExport](#ChainExport)
  * [ChainGetBlock](#ChainGetBlock)
  * [ChainGetBlockMessages](#ChainGetBlockMessages)
  * [ChainGetGenesis](#ChainGetGenesis)
  * [ChainGetMessage](#ChainGetMessage)
  * [ChainGetMessagesInTipset](#ChainGetMessagesInTipset)
  * [ChainGetNode](#ChainGetNode)
  * [ChainGetParentMessages](#ChainGetParentMessages)
  * [ChainGetParentReceipts](#ChainGetParentReceipts)
  * [ChainGetPath](#ChainGetPath)
  * [ChainGetRandomnessFromBeacon](#ChainGetRandomnessFromBeacon)
  * [ChainGetRandomnessFromTickets](#ChainGetRandomnessFromTickets)
  * [ChainGetTipSet](#ChainGetTipSet)
  * [ChainGetTipSetByHeight](#ChainGetTipSetByHeight)
  * [ChainHasObj](#ChainHasObj)
  * [ChainHead](#ChainHead)
  * [ChainNotify](#ChainNotify)
  * [ChainReadObj](#ChainReadObj)
  * [ChainSetHead](#ChainSetHead)
  * [ChainStatObj](#ChainStatObj)
  * [ChainTipSetWeight](#ChainTipSetWeight)
* [Client](#Client)
  * [ClientCalcCommP](#ClientCalcCommP)
  * [ClientCancelDataTransfer](#ClientCancelDataTransfer)
  * [ClientCancelRetrievalDeal](#ClientCancelRetrievalDeal)
  * [ClientDataTransferUpdates](#ClientDataTransferUpdates)
  * [ClientDealPieceCID](#ClientDealPieceCID)
  * [ClientDealSize](#ClientDealSize)
  * [ClientFindData](#ClientFindData)
  * [ClientGenCar](#ClientGenCar)
  * [ClientGetDealInfo](#ClientGetDealInfo)
  * [ClientGetDealStatus](#ClientGetDealStatus)
  * [ClientGetDealUpdates](#ClientGetDealUpdates)
  * [ClientGetRetrievalUpdates](#ClientGetRetrievalUpdates)
  * [ClientHasLocal](#ClientHasLocal)
  * [ClientImport](#ClientImport)
  * [ClientListDataTransfers](#ClientListDataTransfers)
  * [ClientListDeals](#ClientListDeals)
  * [ClientListImports](#ClientListImports)
  * [ClientListRetrievals](#ClientListRetrievals)
  * [ClientMinerQueryOffer](#ClientMinerQueryOffer)
  * [ClientQueryAsk](#ClientQueryAsk)
  * [ClientRemoveImport](#ClientRemoveImport)
  * [ClientRestartDataTransfer](#ClientRestartDataTransfer)
  * [ClientRetrieve](#ClientRetrieve)
  * [ClientRetrieveTryRestartInsufficientFunds](#ClientRetrieveTryRestartInsufficientFunds)
  * [ClientRetrieveWithEvents](#ClientRetrieveWithEvents)
  * [ClientStartDeal](#ClientStartDeal)
  * [ClientStatelessDeal](#ClientStatelessDeal)
* [Create](#Create)
  * [CreateBackup](#CreateBackup)
* [Gas](#Gas)
  * [GasBatchEstimateMessageGas](#GasBatchEstimateMessageGas)
  * [GasEstimateFeeCap](#GasEstimateFeeCap)
  * [GasEstimateGasLimit](#GasEstimateGasLimit)
  * [GasEstimateGasPremium](#GasEstimateGasPremium)
  * [GasEstimateMessageGas](#GasEstimateMessageGas)
* [I](#I)
  * [ID](#ID)
* [Log](#Log)
  * [LogList](#LogList)
  * [LogSetLevel](#LogSetLevel)
* [Market](#Market)
  * [MarketAddBalance](#MarketAddBalance)
  * [MarketGetReserved](#MarketGetReserved)
  * [MarketReleaseFunds](#MarketReleaseFunds)
  * [MarketReserveFunds](#MarketReserveFunds)
  * [MarketWithdraw](#MarketWithdraw)
* [Miner](#Miner)
  * [MinerCreateBlock](#MinerCreateBlock)
  * [MinerGetBaseInfo](#MinerGetBaseInfo)
* [Mpool](#Mpool)
  * [MpoolBatchPush](#MpoolBatchPush)
  * [MpoolBatchPushMessage](#MpoolBatchPushMessage)
  * [MpoolBatchPushUntrusted](#MpoolBatchPushUntrusted)
  * [MpoolCheckMessages](#MpoolCheckMessages)
  * [MpoolCheckPendingMessages](#MpoolCheckPendingMessages)
  * [MpoolCheckReplaceMessages](#MpoolCheckReplaceMessages)
  * [MpoolClear](#MpoolClear)
  * [MpoolGetConfig](#MpoolGetConfig)
  * [MpoolGetNonce](#MpoolGetNonce)
  * [MpoolPending](#MpoolPending)
  * [MpoolPublishByAddr](#MpoolPublishByAddr)
  * [MpoolPublishMessage](#MpoolPublishMessage)
  * [MpoolPush](#MpoolPush)
  * [MpoolPushMessage](#MpoolPushMessage)
  * [MpoolPushUntrusted](#MpoolPushUntrusted)
  * [MpoolSelect](#MpoolSelect)
  * [MpoolSelects](#MpoolSelects)
  * [MpoolSetConfig](#MpoolSetConfig)
  * [MpoolSub](#MpoolSub)
* [Msig](#Msig)
  * [MsigAddApprove](#MsigAddApprove)
  * [MsigAddCancel](#MsigAddCancel)
  * [MsigAddPropose](#MsigAddPropose)
  * [MsigApprove](#MsigApprove)
  * [MsigApproveTxnHash](#MsigApproveTxnHash)
  * [MsigCancel](#MsigCancel)
  * [MsigCreate](#MsigCreate)
  * [MsigGetAvailableBalance](#MsigGetAvailableBalance)
  * [MsigGetPending](#MsigGetPending)
  * [MsigGetVested](#MsigGetVested)
  * [MsigGetVestingSchedule](#MsigGetVestingSchedule)
  * [MsigPropose](#MsigPropose)
  * [MsigRemoveSigner](#MsigRemoveSigner)
  * [MsigSwapApprove](#MsigSwapApprove)
  * [MsigSwapCancel](#MsigSwapCancel)
  * [MsigSwapPropose](#MsigSwapPropose)
* [Net](#Net)
  * [NetAddrsListen](#NetAddrsListen)
  * [NetAgentVersion](#NetAgentVersion)
  * [NetAutoNatStatus](#NetAutoNatStatus)
  * [NetBandwidthStats](#NetBandwidthStats)
  * [NetBandwidthStatsByPeer](#NetBandwidthStatsByPeer)
  * [NetBandwidthStatsByProtocol](#NetBandwidthStatsByProtocol)
  * [NetBlockAdd](#NetBlockAdd)
  * [NetBlockList](#NetBlockList)
  * [NetBlockRemove](#NetBlockRemove)
  * [NetConnect](#NetConnect)
  * [NetConnectedness](#NetConnectedness)
  * [NetDisconnect](#NetDisconnect)
  * [NetFindPeer](#NetFindPeer)
  * [NetPeerInfo](#NetPeerInfo)
  * [NetPeers](#NetPeers)
  * [NetPubsubScores](#NetPubsubScores)
* [Node](#Node)
  * [NodeStatus](#NodeStatus)
* [Paych](#Paych)
  * [PaychAllocateLane](#PaychAllocateLane)
  * [PaychAvailableFunds](#PaychAvailableFunds)
  * [PaychAvailableFundsByFromTo](#PaychAvailableFundsByFromTo)
  * [PaychCollect](#PaychCollect)
  * [PaychGet](#PaychGet)
  * [PaychGetWaitReady](#PaychGetWaitReady)
  * [PaychList](#PaychList)
  * [PaychNewPayment](#PaychNewPayment)
  * [PaychSettle](#PaychSettle)
  * [PaychStatus](#PaychStatus)
  * [PaychVoucherAdd](#PaychVoucherAdd)
  * [PaychVoucherCheckSpendable](#PaychVoucherCheckSpendable)
  * [PaychVoucherCheckValid](#PaychVoucherCheckValid)
  * [PaychVoucherCreate](#PaychVoucherCreate)
  * [PaychVoucherList](#PaychVoucherList)
  * [PaychVoucherSubmit](#PaychVoucherSubmit)
* [State](#State)
  * [StateAccountKey](#StateAccountKey)
  * [StateAllMinerFaults](#StateAllMinerFaults)
  * [StateCall](#StateCall)
  * [StateChangedActors](#StateChangedActors)
  * [StateCirculatingSupply](#StateCirculatingSupply)
  * [StateCompute](#StateCompute)
  * [StateDealProviderCollateralBounds](#StateDealProviderCollateralBounds)
  * [StateDecodeParams](#StateDecodeParams)
  * [StateGetActor](#StateGetActor)
  * [StateListActors](#StateListActors)
  * [StateListMessages](#StateListMessages)
  * [StateListMiners](#StateListMiners)
  * [StateLookupID](#StateLookupID)
  * [StateMarketBalance](#StateMarketBalance)
  * [StateMarketDeals](#StateMarketDeals)
  * [StateMarketParticipants](#StateMarketParticipants)
  * [StateMarketStorageDeal](#StateMarketStorageDeal)
  * [StateMinerActiveSectors](#StateMinerActiveSectors)
  * [StateMinerAvailableBalance](#StateMinerAvailableBalance)
  * [StateMinerDeadlines](#StateMinerDeadlines)
  * [StateMinerFaults](#StateMinerFaults)
  * [StateMinerInfo](#StateMinerInfo)
  * [StateMinerInitialPledgeCollateral](#StateMinerInitialPledgeCollateral)
  * [StateMinerPartitions](#StateMinerPartitions)
  * [StateMinerPower](#StateMinerPower)
  * [StateMinerPreCommitDepositForPower](#StateMinerPreCommitDepositForPower)
  * [StateMinerProvingDeadline](#StateMinerProvingDeadline)
  * [StateMinerRecoveries](#StateMinerRecoveries)
  * [StateMinerSectorAllocated](#StateMinerSectorAllocated)
  * [StateMinerSectorCount](#StateMinerSectorCount)
  * [StateMinerSectors](#StateMinerSectors)
  * [StateNetworkName](#StateNetworkName)
  * [StateNetworkVersion](#StateNetworkVersion)
  * [StateReadState](#StateReadState)
  * [StateReplay](#StateReplay)
  * [StateSearchMsg](#StateSearchMsg)
  * [StateSectorExpiration](#StateSectorExpiration)
  * [StateSectorGetInfo](#StateSectorGetInfo)
  * [StateSectorPartition](#StateSectorPartition)
  * [StateSectorPreCommitInfo](#StateSectorPreCommitInfo)
  * [StateVMCirculatingSupplyInternal](#StateVMCirculatingSupplyInternal)
  * [StateVerifiedClientStatus](#StateVerifiedClientStatus)
  * [StateVerifiedRegistryRootKey](#StateVerifiedRegistryRootKey)
  * [StateVerifierStatus](#StateVerifierStatus)
  * [StateWaitMsg](#StateWaitMsg)
* [Sync](#Sync)
  * [SyncCheckBad](#SyncCheckBad)
  * [SyncCheckpoint](#SyncCheckpoint)
  * [SyncIncomingBlocks](#SyncIncomingBlocks)
  * [SyncMarkBad](#SyncMarkBad)
  * [SyncState](#SyncState)
  * [SyncSubmitBlock](#SyncSubmitBlock)
  * [SyncUnmarkAllBad](#SyncUnmarkAllBad)
  * [SyncUnmarkBad](#SyncUnmarkBad)
  * [SyncValidateTipset](#SyncValidateTipset)
* [Wallet](#Wallet)
  * [WalletBalance](#WalletBalance)
  * [WalletDefaultAddress](#WalletDefaultAddress)
  * [WalletDelete](#WalletDelete)
  * [WalletExport](#WalletExport)
  * [WalletHas](#WalletHas)
  * [WalletImport](#WalletImport)
  * [WalletList](#WalletList)
  * [WalletNew](#WalletNew)
  * [WalletSetDefault](#WalletSetDefault)
  * [WalletSign](#WalletSign)
  * [WalletSignMessage](#WalletSignMessage)
  * [WalletValidateAddress](#WalletValidateAddress)
  * [WalletVerify](#WalletVerify)
## 


### Closing


Perms: read

Inputs: `null`

Response: `{}`

### Discover


Perms: read

Inputs: `null`

Response:
```json
{
  "info": {
    "title": "Lotus RPC API",
    "version": "1.2.1/generated=2020-11-22T08:22:42-06:00"
  },
  "methods": [],
  "openrpc": "1.2.6"
}
```

### Session


Perms: read

Inputs: `null`

Response: `"07070707-0707-0707-0707-070707070707"`

### Shutdown


Perms: admin

Inputs: `null`

Response: `{}`

### Version


Perms: read

Inputs: `null`

Response:
```json
{
  "Version": "string value",
  "APIVersion": 131328,
  "BlockDelay": 42
}
```

## Auth


### AuthNew


Perms: admin

Inputs:
```json
[
  null
]
```

Response: `"Ynl0ZSBhcnJheQ=="`

### AuthVerify


Perms: read

Inputs:
```json
[
  "string value"
]
```

Response: `null`

## Beacon
The Beacon method group contains methods for interacting with the random beacon (DRAND)


### BeaconGetEntry
BeaconGetEntry returns the beacon entry for the given filecoin epoch. If
the entry has not yet been produced, the call will block until the entry
becomes available


Perms: read

Inputs:
```json
[
  10101
]
```

Response:
```json
{
  "Round": 42,
  "Data": "Ynl0ZSBhcnJheQ=="
}
```

## Chain
The Chain method group contains methods for interacting with the
blockchain, but that do not require any form of state computation.


### ChainDeleteObj
ChainDeleteObj deletes node referenced by the given CID


Perms: admin

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `{}`

### ChainExport
ChainExport returns a stream of bytes with CAR dump of chain data.
The exported chain data includes the header chain from the given tipset
back to genesis, the entire genesis state, and the most recent 'nroots'
state trees.
If oldmsgskip is set, messages from before the requested roots are also not included.


Perms: read

Inputs:
```json
[
  10101,
  true,
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response: `"Ynl0ZSBhcnJheQ=="`

### ChainGetBlock
ChainGetBlock returns the block specified by the given CID.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "Miner": "f01234",
  "Ticket": {
    "VRFProof": "Ynl0ZSBhcnJheQ=="
  },
  "ElectionProof": {
    "WinCount": 9,
    "VRFProof": "Ynl0ZSBhcnJheQ=="
  },
  "BeaconEntries": null,
  "WinPoStProof": null,
  "Parents": null,
  "ParentWeight": "0",
  "Height": 10101,
  "ParentStateRoot": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ParentMessageReceipts": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Messages": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "BLSAggregate": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  },
  "Timestamp": 42,
  "BlockSig": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  },
  "ForkSignaling": 42,
  "ParentBaseFee": "0"
}
```

### ChainGetBlockMessages
ChainGetBlockMessages returns messages stored in the specified block.

Note: If there are multiple blocks in a tipset, it's likely that some
messages will be duplicated. It's also possible for blocks in a tipset to have
different messages from the same sender at the same nonce. When that happens,
only the first message (in a block with lowest ticket) will be considered
for execution

NOTE: THIS METHOD SHOULD ONLY BE USED FOR GETTING MESSAGES IN A SPECIFIC BLOCK

DO NOT USE THIS METHOD TO GET MESSAGES INCLUDED IN A TIPSET
Use ChainGetParentMessages, which will perform correct message deduplication


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "BlsMessages": null,
  "SecpkMessages": null,
  "Cids": null
}
```

### ChainGetGenesis
ChainGetGenesis returns the genesis tipset.


Perms: read

Inputs: `null`

Response:
```json
{
  "Cids": null,
  "Blocks": null,
  "Height": 0
}
```

### ChainGetMessage
ChainGetMessage reads a message referenced by the specified CID from the
chain blockstore.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "Version": 42,
  "To": "f01234",
  "From": "f01234",
  "Nonce": 42,
  "Value": "0",
  "GasLimit": 9,
  "GasFeeCap": "0",
  "GasPremium": "0",
  "Method": 1,
  "Params": "Ynl0ZSBhcnJheQ==",
  "CID": {
    "/": "bafy2bzacebbpdegvr3i4cosewthysg5xkxpqfn2wfcz6mv2hmoktwbdxkax4s"
  }
}
```

### ChainGetMessagesInTipset
ChainGetMessagesInTipset returns message stores in current tipset


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response: `null`

### ChainGetNode


Perms: read

Inputs:
```json
[
  "string value"
]
```

Response:
```json
{
  "Cid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Obj": {}
}
```

### ChainGetParentMessages
ChainGetParentMessages returns messages stored in parent tipset of the
specified block.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `null`

### ChainGetParentReceipts
ChainGetParentReceipts returns receipts for messages in parent tipset of
the specified block. The receipts in the list returned is one-to-one with the
messages returned by a call to ChainGetParentMessages with the same blockCid.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `null`

### ChainGetPath
ChainGetPath returns a set of revert/apply operations needed to get from
one tipset to another, for example:
```
       to
        ^
from   tAA
  ^     ^
tBA    tAB
 ^---*--^
     ^
    tRR
```
Would return `[revert(tBA), apply(tAB), apply(tAA)]`


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ],
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response: `null`

### ChainGetRandomnessFromBeacon
ChainGetRandomnessFromBeacon is used to sample the beacon for randomness.


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ],
  2,
  10101,
  "Ynl0ZSBhcnJheQ=="
]
```

Response: `null`

### ChainGetRandomnessFromTickets
ChainGetRandomnessFromTickets is used to sample the chain for randomness.


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ],
  2,
  10101,
  "Ynl0ZSBhcnJheQ=="
]
```

Response: `null`

### ChainGetTipSet
ChainGetTipSet returns the tipset specified by the given TipSetKey.


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response:
```json
{
  "Cids": null,
  "Blocks": null,
  "Height": 0
}
```

### ChainGetTipSetByHeight
ChainGetTipSetByHeight looks back for a tipset at the specified epoch.
If there are no blocks at the specified epoch, a tipset at an earlier epoch
will be returned.


Perms: read

Inputs:
```json
[
  10101,
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response:
```json
{
  "Cids": null,
  "Blocks": null,
  "Height": 0
}
```

### ChainHasObj
ChainHasObj checks if a given CID exists in the chain blockstore.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `true`

### ChainHead
ChainHead returns the current head of the chain.


Perms: read

Inputs: `null`

Response:
```json
{
  "Cids": null,
  "Blocks": null,
  "Height": 0
}
```

### ChainNotify
ChainNotify returns channel with chain head updates.
First message is guaranteed to be of len == 1, and type == 'current'.


Perms: read

Inputs: `null`

Response: `null`

### ChainReadObj
ChainReadObj reads ipld nodes referenced by the specified CID from chain
blockstore and returns raw bytes.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `"Ynl0ZSBhcnJheQ=="`

### ChainSetHead
ChainSetHead forcefully sets current chain head. Use with caution.


Perms: admin

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response: `{}`

### ChainStatObj
ChainStatObj returns statistics about the graph referenced by 'obj'.
If 'base' is also specified, then the returned stat will be a diff
between the two objects.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "Size": 42,
  "Links": 42
}
```

### ChainTipSetWeight
ChainTipSetWeight computes weight for the specified tipset.


Perms: read

Inputs:
```json
[
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ]
]
```

Response: `"0"`

## Client
The Client methods all have to do with interacting with the storage and
retrieval markets as a client


### ClientCalcCommP
ClientCalcCommP calculates the CommP for a specified file


Perms: write

Inputs:
```json
[
  "string value"
]
```

Response:
```json
{
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 1024
}
```

### ClientCancelDataTransfer
ClientCancelDataTransfer cancels a data transfer with the given transfer ID and other peer


Perms: write

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### ClientCancelRetrievalDeal
ClientCancelRetrievalDeal cancels an ongoing retrieval deal based on DealID


Perms: write

Inputs:
```json
[
  5
]
```

Response: `{}`

### ClientDataTransferUpdates


Perms: write

Inputs: `null`

Response:
```json
{
  "TransferID": 3,
  "Status": 1,
  "BaseCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "IsInitiator": true,
  "IsSender": true,
  "Voucher": "string value",
  "Message": "string value",
  "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Transferred": 42,
  "Stages": {
    "Stages": null
  }
}
```

### ClientDealPieceCID
ClientCalcCommP calculates the CommP and data size of the specified CID


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "PayloadSize": 9,
  "PieceSize": 1032,
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
}
```

### ClientDealSize
ClientDealSize calculates real deal data size


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "PayloadSize": 9,
  "PieceSize": 1032
}
```

### ClientFindData
ClientFindData identifies peers that have a certain file, and returns QueryOffers (one per peer).


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  null
]
```

Response: `null`

### ClientGenCar
ClientGenCar generates a CAR file for the specified file.


Perms: write

Inputs:
```json
[
  {
    "Path": "string value",
    "IsCAR": true
  },
  "string value"
]
```

Response: `{}`

### ClientGetDealInfo
ClientGetDealInfo returns the latest information about a given deal.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 42,
  "Message": "string value",
  "DealStages": {
    "Stages": null
  },
  "Provider": "f01234",
  "DataRef": {
    "TransferType": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceCid": null,
    "PieceSize": 1024,
    "RawBlockSize": 42
  },
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 42,
  "PricePerEpoch": "0",
  "Duration": 42,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "Verified": true,
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": null
    }
  }
}
```

### ClientGetDealStatus
ClientGetDealStatus returns status given a code


Perms: read

Inputs:
```json
[
  42
]
```

Response: `"string value"`

### ClientGetDealUpdates
ClientGetDealUpdates returns the status of updated deals


Perms: write

Inputs: `null`

Response:
```json
{
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 42,
  "Message": "string value",
  "DealStages": {
    "Stages": null
  },
  "Provider": "f01234",
  "DataRef": {
    "TransferType": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceCid": null,
    "PieceSize": 1024,
    "RawBlockSize": 42
  },
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 42,
  "PricePerEpoch": "0",
  "Duration": 42,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "Verified": true,
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": null
    }
  }
}
```

### ClientGetRetrievalUpdates
ClientGetRetrievalUpdates returns status of updated retrieval deals


Perms: write

Inputs: `null`

Response:
```json
{
  "PayloadCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ID": 5,
  "PieceCID": null,
  "PricePerByte": "0",
  "UnsealPrice": "0",
  "Status": 0,
  "Message": "string value",
  "Provider": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "BytesReceived": 42,
  "BytesPaidFor": 42,
  "TotalPaid": "0",
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": null
    }
  }
}
```

### ClientHasLocal
ClientHasLocal indicates whether a certain CID is locally stored.


Perms: write

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `true`

### ClientImport
ClientImport imports file under the specified path into filestore.


Perms: admin

Inputs:
```json
[
  {
    "Path": "string value",
    "IsCAR": true
  }
]
```

Response:
```json
{
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ImportID": 50
}
```

### ClientListDataTransfers
ClientListTransfers returns the status of all ongoing transfers of data


Perms: write

Inputs: `null`

Response: `null`

### ClientListDeals
ClientListDeals returns information about the deals made by the local client.


Perms: write

Inputs: `null`

Response: `null`

### ClientListImports
ClientListImports lists imported files and their root CIDs


Perms: write

Inputs: `null`

Response: `null`

### ClientListRetrievals
ClientListRetrievals returns information about retrievals made by the local client


Perms: write

Inputs: `null`

Response: `null`

### ClientMinerQueryOffer
ClientMinerQueryOffer returns a QueryOffer for the specific miner and file.


Perms: read

Inputs:
```json
[
  "f01234",
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  null
]
```

Response:
```json
{
  "Err": "string value",
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Piece": null,
  "Size": 42,
  "MinPrice": "0",
  "UnsealPrice": "0",
  "PaymentInterval": 42,
  "PaymentIntervalIncrease": 42,
  "Miner": "f01234",
  "MinerPeer": {
    "Address": "f01234",
    "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "PieceCID": null
  }
}
```

### ClientQueryAsk
ClientQueryAsk returns a signed StorageAsk from the specified miner.


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "f01234"
]
```

Response:
```json
{
  "Price": "0",
  "VerifiedPrice": "0",
  "MinPieceSize": 1032,
  "MaxPieceSize": 1032,
  "Miner": "f01234",
  "Timestamp": 10101,
  "Expiry": 10101,
  "SeqNo": 42
}
```

### ClientRemoveImport
ClientRemoveImport removes file import


Perms: admin

Inputs:
```json
[
  50
]
```

Response: `{}`

### ClientRestartDataTransfer
ClientRestartDataTransfer attempts to restart a data transfer with the given transfer ID and other peer


Perms: write

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### ClientRetrieve
ClientRetrieve initiates the retrieval of a file, as specified in the order.


Perms: admin

Inputs:
```json
[
  {
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Piece": null,
    "Size": 42,
    "LocalStore": 12,
    "Total": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "Client": "f01234",
    "Miner": "f01234",
    "MinerPeer": {
      "Address": "f01234",
      "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "PieceCID": null
    }
  },
  {
    "Path": "string value",
    "IsCAR": true
  }
]
```

Response: `{}`

### ClientRetrieveTryRestartInsufficientFunds
ClientRetrieveTryRestartInsufficientFunds attempts to restart stalled retrievals on a given payment channel
which are stuck due to insufficient funds


Perms: write

Inputs:
```json
[
  "f01234"
]
```

Response: `{}`

### ClientRetrieveWithEvents
ClientRetrieveWithEvents initiates the retrieval of a file, as specified in the order, and provides a channel
of status updates.


Perms: admin

Inputs:
```json
[
  {
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Piece": null,
    "Size": 42,
    "LocalStore": 12,
    "Total": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "Client": "f01234",
    "Miner": "f01234",
    "MinerPeer": {
      "Address": "f01234",
      "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "PieceCID": null
    }
  },
  {
    "Path": "string value",
    "IsCAR": true
  }
]
```

Response:
```json
{
  "Event": 5,
  "Status": 0,
  "BytesReceived": 42,
  "FundsSpent": "0",
  "Err": "string value"
}
```

### ClientStartDeal
ClientStartDeal proposes a deal with a miner.


Perms: admin

Inputs:
```json
[
  {
    "Data": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": null,
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "Wallet": "f01234",
    "Miner": "f01234",
    "EpochPrice": "0",
    "MinBlocksDuration": 42,
    "ProviderCollateral": "0",
    "DealStartEpoch": 10101,
    "FastRetrieval": true,
    "VerifiedDeal": true
  }
]
```

Response: `null`

### ClientStatelessDeal
ClientStatelessDeal fire-and-forget-proposes an offline deal to a miner without subsequent tracking.


Perms: write

Inputs:
```json
[
  {
    "Data": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": null,
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "Wallet": "f01234",
    "Miner": "f01234",
    "EpochPrice": "0",
    "MinBlocksDuration": 42,
    "ProviderCollateral": "0",
    "DealStartEpoch": 10101,
    "FastRetrieval": true,
    "VerifiedDeal": true
  }
]
```

Response: `null`

## Create


### CreateBackup
CreateBackup creates node backup onder the specified file name. The
method requires that the lotus daemon is running with the
LOTUS_BACKUP_BASE_PATH environment variable set to some path, and that
the path specified when calling CreateBackup is within the base path


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

## Gas


### GasBatchEstimateMessageGas


