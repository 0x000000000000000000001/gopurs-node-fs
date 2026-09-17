package Node_FS_Async

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func statResult(t *testing.T, path string) (error, os.FileInfo) {
	t.Helper()
	type result struct {
		err  any
		info any
	}
	completed := make(chan result, 1)
	StatImpl(path, func(err, info any) any {
		completed <- result{err, info}
		return nil
	})
	select {
	case value := <-completed:
		if value.err != nil {
			if value.info != nil {
				t.Fatal("failed stat returned file info")
			}
			return value.err.(error), nil
		}
		return nil, value.info.(os.FileInfo)
	case <-time.After(3 * time.Second):
		t.Fatal("stat callback did not complete")
		return nil, nil
	}
}

func TestStatMissingFileIncludesENOENTAndPreservesCause(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing")
	err, _ := statResult(t, path)
	if err == nil || !strings.HasPrefix(err.Error(), "ENOENT:") {
		t.Fatalf("missing file error lacks ENOENT: %v", err)
	}
	var pathError *os.PathError
	if !errors.Is(err, os.ErrNotExist) || !errors.As(err, &pathError) || pathError.Path != path {
		t.Fatalf("original missing-file error was lost: %v", err)
	}
}

func TestStatExistingPathsAndOtherErrors(t *testing.T) {
	directory := t.TempDir()
	file := filepath.Join(directory, "file")
	if err := os.WriteFile(file, []byte("hello"), 0600); err != nil {
		t.Fatal(err)
	}
	if err, info := statResult(t, file); err != nil || info.Size() != 5 || !info.Mode().IsRegular() {
		t.Fatalf("file stat: info=%v err=%v", info, err)
	}
	if err, info := statResult(t, directory); err != nil || !info.IsDir() {
		t.Fatalf("directory stat: info=%v err=%v", info, err)
	}
	err, _ := statResult(t, filepath.Join(file, "child"))
	if err == nil || !errors.Is(err, syscall.ENOTDIR) || strings.Contains(err.Error(), "ENOENT") {
		t.Fatalf("non-ENOENT error changed: %v", err)
	}
}
