package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	lockuptypes "github.com/osmosis-labs/osmosis/x/lockup/types"
)

func NewGauge(id uint64, isPerpetual bool, distrTo lockuptypes.QueryCondition, coins sdk.Coins, startTime time.Time, numEpochsPaidOver uint64, filledEpochs uint64, distrCoins sdk.Coins) Gauge {
	return Gauge{
		Id:                id,
		IsPerpetual:       isPerpetual,
		DistributeTo:      distrTo,
		Coins:             coins,
		StartTime:         startTime,
		NumEpochsPaidOver: numEpochsPaidOver,
		FilledEpochs:      filledEpochs,
		DistributedCoins:  distrCoins,
	}
}

func (gauge Gauge) IsUpcomingGauge(curTime time.Time) bool {
	if gauge.StartTime.After(curTime) {
		return true
	}
	return false
}

func (gauge Gauge) IsActiveGauge(curTime time.Time) bool {
	if gauge.StartTime.Before(curTime) {
		return false
	}
	return gauge.IsPerpetual || gauge.FilledEpochs < gauge.NumEpochsPaidOver
}

func (gauge Gauge) IsFinishedGauge(curTime time.Time) bool {
	return !gauge.IsUpcomingGauge(curTime) && !gauge.IsActiveGauge(curTime)
}
