import { execFileSync } from "node:child_process";
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";

const workspace = mkdtempSync(join(tmpdir(), "gopurs-fs-stat-"));
try {
  mkdirSync(join(workspace, "gopurs_runtime"));
  copyFileSync(fileURLToPath(new URL("../../gopurs/runtime/runtime.go", import.meta.url)), join(workspace, "gopurs_runtime/runtime.go"));
  copyFileSync(fileURLToPath(new URL("../src/Node/FS/Async.go", import.meta.url)), join(workspace, "async.go"));
  copyFileSync(fileURLToPath(new URL("./stat-native_test.go", import.meta.url)), join(workspace, "async_test.go"));
  writeFileSync(join(workspace, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  execFileSync("go", ["test", "-count=1", "."], { cwd: workspace, stdio: "inherit" });
} finally {
  rmSync(workspace, { recursive: true, force: true });
}
