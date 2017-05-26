// Do not edit. Bootstrap copy of /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/gc/builtin_test.go

//line /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/gc/builtin_test.go:1
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc_test

import (
	"bytes"
	"internal/testenv"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestBuiltin(t *testing.T) {
	testenv.MustHaveGoRun(t)

	old, err := ioutil.ReadFile("builtin.go")
	if err != nil {
		t.Fatal(err)
	}

	new, err := exec.Command(testenv.GoToolPath(t), "run", "mkbuiltin.go", "-stdout").Output()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(old, new) {
		t.Fatal("builtin.go out of date; run mkbuiltin.go")
	}
}
