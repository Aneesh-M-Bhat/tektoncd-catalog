// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../../gen/gen.go

package ed25519

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("crypto/ed25519", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name:  "PublicKeySize",
		Const: "32",
	}, {
		Name: "Valid",
		Params: []internal.Param{
			{Kind: adt.BytesKind | adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			publicKey, message, signature := c.Bytes(0), c.Bytes(1), c.Bytes(2)
			if c.Do() {
				c.Ret, c.Err = Valid(publicKey, message, signature)
			}
		},
	}},
}
