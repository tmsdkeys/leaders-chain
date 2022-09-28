package keeper

import (
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

var _ types.QueryServer = Keeper{}
