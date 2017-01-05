package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	pkg := "qvl.io/sleepto"
	tmpBin := filepath.Join(os.TempDir(), "sleepto-"+strconv.FormatInt(time.Now().Unix(), 10))

	// build temporary binary
	cmd := exec.Command("go", "build", "-o", tmpBin, pkg)
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
	defer func() {
		if err := os.Remove(tmpBin); err != nil {
			t.Error(err)
		}
	}()

	done := make(chan struct{})
	now := time.Now()
	s := now.Second()

	// Check binary timing
	go func() {
		select {
		case <-done:
			t.Error("Waiting too short")
		case <-time.After(4 * time.Second):
		}

		select {
		case <-time.After(5 * time.Second):
			t.Error("Waiting too long")
		case <-done:
		}
	}()

	// Run binary
	cmd = exec.Command(tmpBin, "-second", fmt.Sprintf("%d,%d", (s+5)%60, (s+50)%60))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}
	close(done)

	// Check binary output
	if len(out) > 0 {
		t.Errorf("no stdout expected but got: %s", out)
	}
	info := fmt.Sprintf("Running at %s\n", now.Add(5*time.Second).Format(time.RFC1123))
	msg := stderr.String()
	if msg != info {
		t.Errorf(`Unexpected message
Expected: %s
Got:      %s
`, info, msg)
	}
}
