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
	want := "hello test"
	cmd = exec.Command(tmpBin, "-second", fmt.Sprintf("%d,%d", (s+5)%60, (s+50)%60), "echo", want)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}
	close(done)
	equal(t, want+"\n", string(out), "stdout")
	want = fmt.Sprintf("Running at %s\n", now.Add(5*time.Second).Format(time.RFC1123))
	equal(t, want, stderr.String(), "stderr")
}

func equal(t *testing.T, want, got, msg string) {
	if want != got {
		t.Errorf(`%s
Expected: %s
Got:      %s
`, msg, want, got)
	}
}
