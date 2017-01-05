package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"testing"
	"time"
)

const pkg = "qvl.io/sleepto"

var tmpbin string

// Provide temporary binary
func TestMain(m *testing.M) {
	tmpbin = filepath.Join(os.TempDir(), "sleepto-"+strconv.FormatInt(time.Now().Unix(), 10))
	cmd := exec.Command("go", "build", "-o", tmpbin, pkg)
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	code := m.Run()

	if err := os.Remove(tmpbin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(code)
}

func TestEcho(t *testing.T) {
	done := make(chan struct{})

	// Run binary
	go func() {
		now := time.Now()
		s := now.Second()
		want := "hello test"
		cmd := exec.Command(tmpbin, "-second", fmt.Sprintf("%d,%d", (s+3)%60, (s+50)%60), "echo", want)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		close(done)
		equal(t, want+"\n", string(out), "stdout")
		want = fmt.Sprintf("sleeping until: %s\n", now.Add(3*time.Second).Format(time.RFC1123))
		equal(t, want, stderr.String(), "stderr")
	}()

	if err := timing(done, 2, 3); err != nil {
		t.Error(err)
	}
}

func TestAlarm(t *testing.T) {
	done := make(chan struct{})
	m := strconv.Itoa(int(time.Now().Month()))
	want := "hello alarm"
	cmd := exec.Command(tmpbin, "-month", m, "-silent", "echo", want)

	// Run binary
	go func() {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		out, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		close(done)
		equal(t, want+"\n", string(out), "stdout")
		equal(t, "", stderr.String(), "stderr")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		if err := cmd.Process.Signal(syscall.SIGALRM); err != nil {
			t.Error(err)
		}
	}()

	if err := timing(done, 2, 3); err != nil {
		t.Error(err)
	}
}

func equal(t *testing.T, want, got, msg string) {
	if want != got {
		t.Errorf(`%s
Expected: %s
Got:      %s
`, msg, want, got)
	}
}

func timing(done chan struct{}, min, max time.Duration) error {
	select {
	case <-done:
		return errors.New("too quick")
	case <-time.After(min * time.Second):
	}
	select {
	case <-time.After((max - min) * time.Second):
		return errors.New("too slow")
	case <-done:
	}
	return nil
}
