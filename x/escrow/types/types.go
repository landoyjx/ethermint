package types

import (
	"os"

	"github.com/tendermint/tendermint/libs/log"
)

var Logger log.Logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

type DayState uint8

const (
	BET DayState = iota
	DRAWING
	PAYOUT
	INVALID
)

// Structure with all the current bets information in a contest period (e.g. day)
type BetDay struct {
	GrandPrize uint64 // total prize for a day
}
