// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package extension

import (
	"github.com/vechain/thor/state"
	"github.com/vechain/thor/thor"
)

// Extension implements native methods of `Extension` contract.
type Extension struct {
	addr  thor.Address
	state *state.State
}

// New create a new instance.
func New(addr thor.Address, state *state.State) *Extension {
	return &Extension{addr, state}
}
