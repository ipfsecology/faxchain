package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/exchain/x/staking/types"
)

const (
	// UTC+8 Time: 2021-07-05 07:05:02
	blockTimestampEpoch = int64(1625439902)
	secondsPerWeek      = int64(60 * 60 * 24 * 7)
	weeksPerYear        = float64(52)
)

func calculateWeight(nowTime int64, tokens sdk.Dec) (shares types.Shares, sdkErr error) {
	nowWeek := (nowTime - blockTimestampEpoch) / secondsPerWeek
	rate := float64(nowWeek) / weeksPerYear
	weight := rate
	precision := fmt.Sprintf("%d", sdk.Precision)

	weightByDec, sdkErr := sdk.NewDecFromStr(fmt.Sprintf("%."+precision+"f", weight))
	if sdkErr == nil {
		shares = tokens.Mul(weightByDec)
	}
	return
}

func SimulateWeight(nowTime int64, tokens sdk.Dec) (votes types.Shares, sdkErr error) {
	return calculateWeight(nowTime, tokens)
}
