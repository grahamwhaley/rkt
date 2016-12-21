// Copyright 2016 The rkt Authors
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

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/coreos/rkt/tests/testutils"
)

func TestRktStop(t *testing.T) {
	image := patchTestACI("rkt-stop-test.aci", "--name=rkt-stop-test", "--exec=/inspect --read-stdin --silent-sigterm")
	defer os.Remove(image)

	ctx := testutils.NewRktRunCtx()
	defer ctx.Cleanup()

	tmpDir := createTempDirOrPanic("rkt-TestRktStop-")
	defer os.RemoveAll(tmpDir)

	// Define tests
	tests := []struct {
		cmd        string
		uuidFile   bool
		expectKill bool
	}{
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
		// Test regular stop
		{
			"stop",
			false,
			false,
		},
		// Test forced stop
		{
			"stop --force",
			false,
			true,
		},
		// Test uuid-file
		{
			"stop",
			true,
			false,
		},
	}

	// Run tests
	for i, tt := range tests {
		//let's time it
		startTime := time.Now()
		// Prepare image
		cmd := fmt.Sprintf("%s --insecure-options=image prepare %s", ctx.Cmd(), image)
		podUUID := runRktAndGetUUID(t, cmd)

		// Run image
		cmd = fmt.Sprintf("%s --insecure-options=image run-prepared --interactive %s", ctx.Cmd(), podUUID)
		child := spawnOrFail(t, cmd)

		// Wait for prompt to make sure the pod is started
		if err := expectTimeoutWithOutput(child, "Enter text:", time.Minute); err != nil {
			t.Fatalf("Can't start pod")
		}

		//Just dump initial podinfo for reference
		var podInfo *podInfo
		//podInfo = getPodInfoVerbose(t, ctx, podUUID)

		args := podUUID
		if tt.uuidFile {
			uuidFile, err := ioutil.TempFile(tmpDir, "uuid-file")
			if err != nil {
				panic(fmt.Sprintf("Cannot create uuid-file: %v", err))
			}
			uuidFilePath := uuidFile.Name()
			if err := ioutil.WriteFile(uuidFilePath, []byte(podUUID), 0600); err != nil {
				panic(fmt.Sprintf("Cannot write pod UUID to uuid-file: %v", err))
			}
			args = fmt.Sprintf("--uuid-file=%s", uuidFilePath)
		}

		runCmd := fmt.Sprintf("%s %s %s", ctx.Cmd(), tt.cmd, args)
		t.Logf("Running test #%d, %s", i, runCmd)
		spawnTime := time.Now()
		spawnOrFail(t, runCmd)

		// Make sure the pod is stopped
		exitedSuccessfully := false
		//Wait for 5 seconds - fine granularity (100mS)
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			podInfo = getPodInfo(t, ctx, podUUID)
			if podInfo.state == "exited" {
				exitedSuccessfully = true
				break
			}
		}
		if !exitedSuccessfully {
			t.Logf("Expected pod %q to be exited, but it is %q", podUUID, podInfo.state)
			t.Logf(" wait some more...")
			//wait another good while to see if we were just too hasty
			//we don't need such a fine granularity here...
			for i := 0; i < 120; i++ {
				time.Sleep(500 * time.Millisecond)
				podInfo = getPodInfo(t, ctx, podUUID)
				if podInfo.state == "exited" {
					exitedSuccessfully = true
					break
				}
			}
		}

		stopTime := time.Now()
		//and dump the podinfo again, for reference
		//podInfo = getPodInfoVerbose(t, ctx, podUUID)

		if !exitedSuccessfully {
			//And now dump the info we got...to try and diagnose!
			podInfo = getPodInfoVerbose(t, ctx, podUUID)
			t.Fatalf("Expected pod %q to be exited, but it is %q after %s", podUUID, podInfo.state, stopTime.Sub(startTime))
		}

		t.Logf(" test #%d took %s to spawn", i, spawnTime.Sub(startTime))
		t.Logf(" test #%d took %s overall", i, stopTime.Sub(startTime))

		exitStatus := 0
		if tt.expectKill {
			exitStatus = -1
		}
		waitOrFail(t, child, exitStatus)

		//And remove the pod - be nice and free up the resources
		// otherwise on some systems we can run out of network addresses for instance
		cmd = fmt.Sprintf("%s rm %s", ctx.Cmd(), podUUID)
		spawnAndWaitOrFail(t, cmd, 0)
	}
}
