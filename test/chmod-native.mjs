import { execFileSync } from "node:child_process";
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";

const runtimeRoot = process.env.GOPURS_ROOT ?? fileURLToPath(new URL("../../gopurs/", import.meta.url));
const workspace = mkdtempSync(join(tmpdir(), "gopurs-fs-chmod-"));
try {
  mkdirSync(join(workspace, "gopurs_runtime"));
  copyFileSync(join(runtimeRoot, "runtime/runtime.go"), join(workspace, "gopurs_runtime/runtime.go"));
  for (const [name, source] of [["sync", "Sync.go"], ["async", "Async.go"]]) {
    mkdirSync(join(workspace, name));
    copyFileSync(fileURLToPath(new URL(`../src/Node/FS/${source}`, import.meta.url)), join(workspace, name, "ffi.go"));
  }
  copyFileSync(fileURLToPath(new URL("./chmod-native_test.go", import.meta.url)), join(workspace, "chmod_test.go"));
  writeFileSync(join(workspace, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  execFileSync("go", ["test", "-count=1", "./..."], { cwd: workspace, stdio: "inherit" });
} finally {
  rmSync(workspace, { recursive: true, force: true });
}
