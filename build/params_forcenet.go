//go:build forcenet
// +build forcenet

package build

import (
	"os"
	"strconv"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const GenesisNetworkVersion = network.Version14

const BootstrappersFile = ""
const GenesisFile = ""

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3
const UpgradeTapeHeight = -4

var UpgradeAssemblyHeight = abi.ChainEpoch(-5)

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = -6

const UpgradeKumquatHeight = -7
const UpgradeCalicoHeight = -9
const UpgradePersianHeight = -10
const UpgradeOrangeHeight = -11
const UpgradeClausHeight = -12
const UpgradeTrustHeight = -13
const UpgradeNorwegianHeight = -14
const UpgradeTurboHeight = -15
const UpgradeHyperdriveHeight = -16
const UpgradeChocolateHeight = -17

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg8MiBV1,
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
	)
	policy.SetPreCommitChallengeDelay(10)

	SetAddressNetwork(address.Testnet)

	Devnet = true
	BuildType |= BuildTypeForcenet

	baseExpireLine := builtin.EpochsInDay

	if baseExpireLineString := os.Getenv("FORCE_BASE_EXPIRE_LINE"); baseExpireLineString != "" {
		tmp, err := strconv.Atoi(baseExpireLineString)
		if err != nil {
			panic(xerrors.Errorf("parse FORCE_BASE_EXPIRE_LINE failed %w", err))
		}
		baseExpireLine = tmp
	}

	miner2.MaxProveCommitDuration = map[abi.RegisteredSealProof]abi.ChainEpoch{
		abi.RegisteredSealProof_StackedDrg32GiBV1:  abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay, // PARAM_SPEC
		abi.RegisteredSealProof_StackedDrg2KiBV1:   abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg8MiBV1:   abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg512MiBV1: abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg64GiBV1:  abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,

		abi.RegisteredSealProof_StackedDrg32GiBV1_1:  abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay, // PARAM_SPEC
		abi.RegisteredSealProof_StackedDrg2KiBV1_1:   abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg8MiBV1_1:   abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg512MiBV1_1: abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
		abi.RegisteredSealProof_StackedDrg64GiBV1_1:  abi.ChainEpoch(baseExpireLine) + miner2.PreCommitChallengeDelay,
	}
	miner2.MaxPreCommitRandomnessLookback = abi.ChainEpoch(baseExpireLine) + miner2.ChainFinality
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 1

var WhitelistedBlock = cid.Undef
