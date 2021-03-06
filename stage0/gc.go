// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build linux

package stage0

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/appc/spec/schema/types"

	"github.com/coreos/rkt/common"
)

// GC enters the pod by fork/exec()ing the stage1's /gc similar to /init.
// /gc can expect to have its CWD set to the pod root.
func GC(pdir string, uuid *types.UUID, debug bool) error {
	ep, err := getStage1Entrypoint(pdir, gcEntrypoint)
	if err != nil {
		log.Fatalf("error determining gc entrypoint: %v", err)
	}

	args := []string{filepath.Join(common.Stage1RootfsPath(pdir), ep)}
	if debug {
		args = append(args, "--debug")
	}
	args = append(args, uuid.String())

	c := exec.Cmd{
		Path:   args[0],
		Args:   args,
		Stderr: os.Stderr,
		Dir:    pdir,
	}
	return c.Run()
}
