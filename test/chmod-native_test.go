package chmod_test

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	async "gopurs/output/async"
	"gopurs/output/gopurs_runtime"
	sync "gopurs/output/sync"
)

func fixture(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(path, []byte("permissions"), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func requireMode(t *testing.T, path string, want os.FileMode) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != want {
		t.Fatalf("file mode = %04o, want %04o", got, want)
	}
}

func panicError(t *testing.T, action func()) (err error) {
	t.Helper()
	defer func() {
		value := recover()
		var ok bool
		err, ok = value.(error)
		if !ok {
			t.Fatalf("expected error panic, got %v", value)
		}
	}()
	action()
	return nil
}

func chmodResult(t *testing.T, path string, mode any) error {
	t.Helper()
	type result struct{ err, value any }
	completed := make(chan result, 1)
	async.ChmodImpl(path, mode, func(err, value any) any {
		completed <- result{err, value}
		return nil
	})
	select {
	case value := <-completed:
		if value.value != nil {
			t.Fatalf("chmod callback returned unexpected result: %v", value.value)
		}
		if value.err == nil {
			return nil
		}
		err, ok := value.err.(error)
		if !ok {
			t.Fatalf("chmod callback error has type %T", value.err)
		}
		return err
	case <-time.After(3 * time.Second):
		t.Fatal("chmod callback did not complete")
		return nil
	}
}

func TestSyncChmodChangesPermissions(t *testing.T) {
	path := fixture(t)
	sync.ChmodSyncImpl(path, "0444")
	requireMode(t, path, 0444)
	sync.ChmodSyncImpl(path, "0640")
	requireMode(t, path, 0640)
}

func TestSyncChmodReportsErrors(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing")
		err := panicError(t, func() { sync.ChmodSyncImpl(path, "0444") })
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("missing file cause was lost: %v", err)
		}
	})
	t.Run("invalid octal", func(t *testing.T) {
		path := fixture(t)
		err := panicError(t, func() { sync.ChmodSyncImpl(path, "0888") })
		var numberError *strconv.NumError
		if !errors.As(err, &numberError) {
			t.Fatalf("invalid mode did not return a parse error: %v", err)
		}
		requireMode(t, path, 0600)
	})
}

func TestAsyncChmodChangesPermissions(t *testing.T) {
	for _, representation := range []struct {
		name string
		wrap func(string) any
	}{
		{"native string", func(s string) any { return s }},
		{"boxed PureScript string", func(s string) any { return gopurs_runtime.Str(s) }},
	} {
		t.Run(representation.name, func(t *testing.T) {
			path := fixture(t)
			for _, mode := range []struct {
				text string
				want os.FileMode
			}{{"0444", 0444}, {"0640", 0640}} {
				if err := chmodResult(t, path, representation.wrap(mode.text)); err != nil {
					t.Fatal(err)
				}
				requireMode(t, path, mode.want)
			}
		})
	}
}

func TestAsyncChmodReportsErrors(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing")
		err := chmodResult(t, path, gopurs_runtime.Str("0444"))
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("missing file cause was lost: %v", err)
		}
	})
	t.Run("invalid octal", func(t *testing.T) {
		path := fixture(t)
		err := chmodResult(t, path, gopurs_runtime.Str("0888"))
		var numberError *strconv.NumError
		if !errors.As(err, &numberError) {
			t.Fatalf("invalid mode did not return a parse error: %v", err)
		}
		requireMode(t, path, 0600)
	})
}
