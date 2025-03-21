// Ignored because implementation relies on external (ffi) lib
package bls_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet/key"
	"github.com/filecoin-project/lotus/lib/sigs"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
)

func TestRoundtrip(t *testing.T) {
	pk, err := sigs.Generate(crypto.SigTypeBLS)
	require.NoError(t, err)

	ki := types.KeyInfo{
		Type:       types.KTBLS,
		PrivateKey: pk,
	}
	k, err := key.NewKey(ki)
	require.NoError(t, err)

	p := []byte("potato")

	si, err := sigs.Sign(crypto.SigTypeBLS, pk, p)
	require.NoError(t, err)

	err = sigs.Verify(si, k.Address, p)
	require.NoError(t, err)
}

func TestUncompressedFails(t *testing.T) {
	// compressed
	err := sigs.Verify(&crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: []byte{0x99, 0x27, 0x44, 0x4b, 0xfc, 0xff, 0xdc, 0xa3, 0x4a, 0xf5, 0x7b, 0x78, 0x75, 0x7b, 0x9b, 0x90, 0xf1, 0xcd, 0x28, 0xd2, 0xa3, 0xae, 0xed, 0x2a, 0xa6, 0xbd, 0xe2, 0x99, 0xf8, 0xbb, 0xb9, 0x18, 0x47, 0x56, 0xf2, 0x28, 0x7b, 0x5, 0x88, 0xe6, 0xd3, 0xf2, 0x86, 0xd, 0x2b, 0xb2, 0x6, 0x6e, 0xc, 0x59, 0x77, 0x8c, 0x1e, 0x64, 0x4f, 0xb2, 0xcf, 0xb3, 0x5f, 0xba, 0x8f, 0x9, 0xfa, 0x82, 0x4a, 0x9e, 0xd8, 0x25, 0x10, 0x8c, 0x82, 0xff, 0x4b, 0xf6, 0x34, 0xc1, 0x3, 0x7e, 0xea, 0xf1, 0x85, 0xf4, 0x56, 0x73, 0xd4, 0xa1, 0xc1, 0xc6, 0xee, 0xb7, 0x12, 0xb7, 0xd7, 0x2a, 0x54, 0x98},
	}, mustAddr("f3tcgq5scpfhdwh4dbalwktzf6mbv3ng2nw7tyzni5cyrsgvineid6jybnweecpa6misa6lk4tvwtxj2gkwpzq"), []byte{0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f})
	require.NoError(t, err)

	// compressed byte changed
	err = sigs.Verify(&crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: []byte{0x99, 0x27, 0x44, 0x4b, 0xfc, 0xff, 0xdc, 0xa3, 0x4a, 0xf5, 0x7b, 0x78, 0x75, 0x7b, 0x9b, 0x90, 0xf1, 0xcd, 0x28, 0xd2, 0xa3, 0xae, 0xed, 0x2a, 0xa6, 0xbd, 0xe2, 0x99, 0xf8, 0xbb, 0xb9, 0x18, 0x47, 0x56, 0xf2, 0x28, 0x7b, 0x5, 0x88, 0xf6, 0xd3, 0xf2, 0x86, 0xd, 0x2b, 0xb2, 0x6, 0x6e, 0xc, 0x59, 0x77, 0x8c, 0x1e, 0x64, 0x4f, 0xb2, 0xcf, 0xb3, 0x5f, 0xba, 0x8f, 0x9, 0xfa, 0x82, 0x4a, 0x9e, 0xd8, 0x25, 0x10, 0x8c, 0x82, 0xff, 0x4b, 0xf6, 0x34, 0xc1, 0x3, 0x7e, 0xea, 0xf1, 0x85, 0xf4, 0x56, 0x73, 0xd4, 0xa1, 0xc1, 0xc6, 0xee, 0xb7, 0x12, 0xb7, 0xd7, 0x2a, 0x54, 0x98},
	}, mustAddr("f3tcgq5scpfhdwh4dbalwktzf6mbv3ng2nw7tyzni5cyrsgvineid6jybnweecpa6misa6lk4tvwtxj2gkwpzq"), []byte{0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f})
	require.Error(t, err)

	// compressed prefix
	err = sigs.Verify(&crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: []byte{0x99, 0x27, 0x44, 0x4b, 0xfc, 0xff, 0xdc, 0xa3, 0x4a, 0xf5, 0x7b, 0x78, 0x75, 0x7b, 0x9b, 0x90, 0xf1, 0xcd, 0x28, 0xd2, 0xa3, 0xae, 0xed, 0x2a, 0xa6, 0xbd, 0xe2, 0x99, 0xf8, 0xbb, 0xb9, 0x18, 0x47, 0x56, 0xf2, 0x28, 0x7b, 0x5, 0x88, 0xe6, 0xd3, 0xf2, 0x86, 0xd, 0x2b, 0xb2, 0x6, 0x6e, 0xc, 0x59, 0x77, 0x8c, 0x1e, 0x64, 0x4f, 0xb2, 0xcf, 0xb3, 0x5f, 0xba, 0x8f, 0x9, 0xfa, 0x82, 0x4a, 0x9e, 0xd8, 0x25, 0x10, 0x8c, 0x82, 0xff, 0x4b, 0xf6, 0x34, 0xc1, 0x3, 0x7e, 0xea, 0xf1, 0x85, 0xf4, 0x56, 0x73, 0xd4, 0xa1, 0xc1, 0xc6, 0xee, 0xb7, 0x12, 0xb7, 0xd7, 0x2a, 0x54, 0x98, 0x55},
	}, mustAddr("f3tcgq5scpfhdwh4dbalwktzf6mbv3ng2nw7tyzni5cyrsgvineid6jybnweecpa6misa6lk4tvwtxj2gkwpzq"), []byte{0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f})
	require.Error(t, err)

	// uncompressed
	err = sigs.Verify(&crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: []byte{0x19, 0x27, 0x44, 0x4b, 0xfc, 0xff, 0xdc, 0xa3, 0x4a, 0xf5, 0x7b, 0x78, 0x75, 0x7b, 0x9b, 0x90, 0xf1, 0xcd, 0x28, 0xd2, 0xa3, 0xae, 0xed, 0x2a, 0xa6, 0xbd, 0xe2, 0x99, 0xf8, 0xbb, 0xb9, 0x18, 0x47, 0x56, 0xf2, 0x28, 0x7b, 0x5, 0x88, 0xe6, 0xd3, 0xf2, 0x86, 0xd, 0x2b, 0xb2, 0x6, 0x6e, 0xc, 0x59, 0x77, 0x8c, 0x1e, 0x64, 0x4f, 0xb2, 0xcf, 0xb3, 0x5f, 0xba, 0x8f, 0x9, 0xfa, 0x82, 0x4a, 0x9e, 0xd8, 0x25, 0x10, 0x8c, 0x82, 0xff, 0x4b, 0xf6, 0x34, 0xc1, 0x3, 0x7e, 0xea, 0xf1, 0x85, 0xf4, 0x56, 0x73, 0xd4, 0xa1, 0xc1, 0xc6, 0xee, 0xb7, 0x12, 0xb7, 0xd7, 0x2a, 0x54, 0x98, 0x8, 0x94, 0x23, 0x78, 0xdb, 0xce, 0x2a, 0xd7, 0x2e, 0x87, 0xdf, 0x8, 0x3b, 0x66, 0xc6, 0x31, 0xc1, 0x8c, 0x58, 0x2f, 0x9f, 0x9e, 0x10, 0x4d, 0x2a, 0x7e, 0x13, 0xe7, 0x9c, 0xbb, 0x22, 0xde, 0xcc, 0xf6, 0x77, 0x77, 0xb0, 0x9c, 0x25, 0x5d, 0x5d, 0xe6, 0x88, 0x9, 0x8c, 0x63, 0x35, 0xd4, 0xa, 0x85, 0x76, 0x8d, 0xb7, 0x66, 0xa6, 0xc6, 0xec, 0xe6, 0xde, 0x2a, 0x9f, 0x34, 0x87, 0x28, 0x1a, 0x48, 0xfe, 0xca, 0xb1, 0x47, 0x2, 0xf6, 0x51, 0x26, 0x52, 0x70, 0x9d, 0x7e, 0xdb, 0x7e, 0x8b, 0xc9, 0xf6, 0x41, 0xaa, 0xa8, 0x3b, 0x7e, 0x8a, 0xfd, 0x7a, 0xe4, 0x79, 0xe6, 0x59, 0xe4},
	}, mustAddr("f3tcgq5scpfhdwh4dbalwktzf6mbv3ng2nw7tyzni5cyrsgvineid6jybnweecpa6misa6lk4tvwtxj2gkwpzq"), []byte{0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f})
	require.Error(t, err)

	// uncompressed one byte change
	err = sigs.Verify(&crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: []byte{0x19, 0x27, 0x44, 0x4b, 0xfc, 0xff, 0xdc, 0xa3, 0x4a, 0xf5, 0x7b, 0x78, 0x75, 0x7b, 0x9b, 0x90, 0xf1, 0xcd, 0x28, 0xd2, 0xa3, 0xae, 0xed, 0x2a, 0xa6, 0xbd, 0xe2, 0x99, 0xf8, 0xbb, 0xb9, 0x18, 0x47, 0x56, 0xf2, 0x28, 0x7b, 0x5, 0x88, 0xe6, 0xd3, 0xf2, 0x86, 0xd, 0x2b, 0xb2, 0x6, 0x6e, 0xc, 0x59, 0x77, 0x8c, 0x1e, 0x64, 0x4f, 0xb2, 0xcf, 0xb3, 0x5f, 0xba, 0x8f, 0x9, 0xfa, 0x82, 0x4a, 0x9e, 0xd8, 0x25, 0x10, 0x8c, 0x82, 0xff, 0x4b, 0xf6, 0x34, 0xc1, 0x3, 0x7e, 0xea, 0xf1, 0x85, 0xf4, 0x56, 0x73, 0xd4, 0xa1, 0xc1, 0xc6, 0xee, 0xb7, 0x12, 0xb7, 0xd7, 0x2a, 0x54, 0x98, 0x8, 0x94, 0x23, 0x78, 0xdb, 0xce, 0x2a, 0xd7, 0x2e, 0x87, 0xdf, 0x8, 0x3b, 0x66, 0xc6, 0x31, 0xc1, 0x8c, 0x58, 0x2f, 0x9f, 0x9e, 0x10, 0x4d, 0x2a, 0x7e, 0x13, 0xe7, 0x9c, 0xbb, 0x22, 0xde, 0xcc, 0xf6, 0x77, 0x77, 0xb0, 0x9c, 0x25, 0x5d, 0x5d, 0xe6, 0x88, 0x9, 0x8c, 0x63, 0x35, 0xd4, 0xa, 0x85, 0x66, 0x8d, 0xb7, 0x66, 0xa6, 0xc6, 0xec, 0xe6, 0xde, 0x2a, 0x9f, 0x34, 0x87, 0x28, 0x1a, 0x48, 0xfe, 0xca, 0xb1, 0x47, 0x2, 0xf6, 0x51, 0x26, 0x52, 0x70, 0x9d, 0x7e, 0xdb, 0x7e, 0x8b, 0xc9, 0xf6, 0x41, 0xaa, 0xa8, 0x3b, 0x7e, 0x8a, 0xfd, 0x7a, 0xe4, 0x79, 0xe6, 0x59, 0xe4},
	}, mustAddr("f3tcgq5scpfhdwh4dbalwktzf6mbv3ng2nw7tyzni5cyrsgvineid6jybnweecpa6misa6lk4tvwtxj2gkwpzq"), []byte{0x70, 0x6f, 0x74, 0x61, 0x74, 0x6f})
	require.Error(t, err)
}

func mustAddr(a string) address.Address {
	ad, _ := address.NewFromString(a)
	return ad
}
