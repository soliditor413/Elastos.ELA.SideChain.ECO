// Copyright 2015 The Elastos.ELA.SideChain.ECO Authors
// This file is part of the Elastos.ELA.SideChain.ECO library.
//
// The Elastos.ELA.SideChain.ECO library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Elastos.ELA.SideChain.ECO library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Elastos.ELA.SideChain.ECO library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.

// nmn005: enode://eca689dac06f246cc30b0aaa48d7750ed5f334d3c9e40c66adb2a1672d5374606e55e0abad2136b0bcdba2b08c59f1e7191c8bd0338d7d09650d0ae4f2c6d632@172.31.25.19:0?discport=20650
// nmn010: enode://bce5df778247c03f3066e2c7c30d49df18fa349b4ade9347568c02054ff013a8ce7232df4592828c23cf0c65fbe22fcf020d27addb077152ccb4aceeff5ab879@172.31.8.35:0?discport=20650
// nmn015: enode://395397db78687a067a3701cb4835a208d7016d6610d1a8206e1758188dfd8295a089a7c603da24428d99afc08abfbd0111779a8bb31b8f8209233cad3ce30c06@172.31.39.1:0?discport=20650
// nmn020: enode://4ae19dba00b2a459ed6ccfb34dbd6303d2c5ba72d5d6003e1ae7c8c5449c31bd39cb896b2fc1b76de3df1bd1eaefa6895eca617c656b5d4b8d8e960824eb941a@172.31.24.100:0?discport=20650
var MainnetBootnodes = []string{
	"enode://eca689dac06f246cc30b0aaa48d7750ed5f334d3c9e40c66adb2a1672d5374606e55e0abad2136b0bcdba2b08c59f1e7191c8bd0338d7d09650d0ae4f2c6d632@172.31.25.19:0?discport=20650",
	"enode://bce5df778247c03f3066e2c7c30d49df18fa349b4ade9347568c02054ff013a8ce7232df4592828c23cf0c65fbe22fcf020d27addb077152ccb4aceeff5ab879@172.31.8.35:0?discport=20650",
	"enode://395397db78687a067a3701cb4835a208d7016d6610d1a8206e1758188dfd8295a089a7c603da24428d99afc08abfbd0111779a8bb31b8f8209233cad3ce30c06@172.31.39.1:0?discport=20650",
	"enode://4ae19dba00b2a459ed6ccfb34dbd6303d2c5ba72d5d6003e1ae7c8c5449c31bd39cb896b2fc1b76de3df1bd1eaefa6895eca617c656b5d4b8d8e960824eb941a@172.31.24.100:0?discport=20650",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://01a22ea82977e9ce3eafd3d7144f9e3d1fed6fde5e7ece98a62d67565f0445f587e05499a73aa7f4d1da537668c1e03ab74ad01ba3a0f11713d045bb214ff04d@13.234.24.155:0?discport=20650",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
