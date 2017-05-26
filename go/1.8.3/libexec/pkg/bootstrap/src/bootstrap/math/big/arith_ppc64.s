// Do not edit. Bootstrap copy of /private/tmp/go-20170524-61097-1ppddl4/go/src/math/big/arith_ppc64.s

//line /private/tmp/go-20170524-61097-1ppddl4/go/src/math/big/arith_ppc64.s:1
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !math_big_pure_go,ppc64

#include "textflag.h"

// This file provides fast assembly versions for the elementary
// arithmetic operations on vectors implemented in arith.go.

TEXT ·divWW(SB), NOSPLIT, $0
	BR ·divWW_g(SB)

