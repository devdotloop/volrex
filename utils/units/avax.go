// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package units

// Denominations of value
const (
	NanoVolrex  uint64 = 1
	MicroVolrex uint64 = 1000 * NanoVolrex
	Schmeckle   uint64 = 49*MicroVolrex + 463*NanoVolrex
	MilliVolrex uint64 = 1000 * MicroVolrex
	Volrex      uint64 = 1000 * MilliVolrex
	KiloVolrex  uint64 = 1000 * Volrex
	MegaVolrex  uint64 = 1000 * KiloVolrex
)
