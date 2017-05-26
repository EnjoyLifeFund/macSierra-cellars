// Do not edit. Bootstrap copy of /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/ppc64/opt.go

//line /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/ppc64/opt.go:1
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

// Many Power ISA arithmetic and logical instructions come in four
// standard variants. These bits let us map between variants.
const (
	V_CC = 1 << 0 // xCC (affect CR field 0 flags)
	V_V  = 1 << 1 // xV (affect SO and OV flags)
)
