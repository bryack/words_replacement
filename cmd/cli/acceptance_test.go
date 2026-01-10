package main_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/bryack/words/adapters/acceptance"
	"github.com/bryack/words/specifications"
)

var (
	buildOnce  sync.Once
	binaryPath string
	buildError error
)

func ensureBinary() (string, error) {
	buildOnce.Do(func() {
		binPath, err := buildBinaryPath()
		if err != nil {
			buildError = err
		}
		binaryPath = binPath
	})
	return binaryPath, buildError
}

func TestWordReplacerSpecification(t *testing.T) {
	binaryPath, err := ensureBinary()
	if err != nil {
		t.Fatal(err)
	}

	// t.Cleanup выполнится только ПОСЛЕ того, как завершатся все t.Run ниже
	t.Cleanup(func() {
		err := os.RemoveAll(filepath.Dir(binaryPath))
		if err != nil {
			t.Logf("Warning: failed to remove temp directory: %v", err)
		}
	})

	t.Run("should replace words in text", func(t *testing.T) {
		driver := &acceptance.Driver{BinaryPath: binaryPath}
		specifications.WordReplacerSpecification(t, driver)
	})
}

func buildBinaryPath() (string, error) {
	tempDir, err := os.MkdirTemp("", "test-binary-*")
	if err != nil {
		return "", fmt.Errorf("failed to make temp directory: %v", err)
	}
	binName := "temp-binary"
	binPath := filepath.Join(tempDir, binName)

	build := exec.Command("go", "build", "-cover", "-o", binPath, ".")
	build.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("GOCOVERDIR"))
	var stderr bytes.Buffer
	build.Stderr = &stderr

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Go build Error: \n%q\n", stderr.String())
		return "", fmt.Errorf("cannot build tool %s: %v", binName, err)
	}
	return binPath, nil
}
