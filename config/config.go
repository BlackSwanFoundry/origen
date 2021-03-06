package bsfconfig

import (
	"flag"
	"log"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)

	config.SetCoinType(CoinType)

	genUnit := sdk.OneDec()
	err := sdk.RegisterDenom(HumanCoinUnit, genUnit)
	if err != nil {
		log.Fatal(err)
	}

	davidsonUnit := sdk.NewDecWithPrec(1, int64(GenExponent)) // 10^-18 (davidson)
	err = sdk.RegisterDenom(BaseCoinUnit, davidsonUnit)

	if err != nil {
		log.Fatal(err)
	}

	config.Seal()
}

var testingConfigState = struct {
	mtx   sync.Mutex
	isSet bool
}{
	isSet: false,
}

func SetTestingConfig() {
	if !isGoTest() {
		panic("SetTestingConfig called but not running go test")
	}

	testingConfigState.mtx.Lock()
	defer testingConfigState.mtx.Unlock()

	if testingConfigState.isSet {
		return
	}

	SetConfig()

	testingConfigState.isSet = true
}

func isGoTest() bool {
	return flag.Lookup("test.v") != nil
}
