import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';

// Run the direct Go test process. A zero exit without the asynchronous pipe
// assertions completing must not turn into a false-positive suite result.
const [command, ...args] = process.argv.slice(2);
assert.ok(command, 'Usage: node test/streams-result.mjs COMMAND [ARG ...]');
const result = spawnSync(command, args, {
  encoding: 'utf8', timeout: 30_000, maxBuffer: 4 * 1024 * 1024,
});
process.stdout.write(result.stdout ?? '');
process.stderr.write(result.stderr ?? '');
assert.ifError(result.error);
assert.equal(result.status, 0, `Test process failed (${result.signal ?? result.status})`);
const lines = result.stdout.split(/\r?\n/);
for (const name of ['multichunk', 'empty']) {
  assert.ok(lines.includes(`PASS file pipe ${name}`), `Missing completed file pipe assertion: ${name}`);
}
