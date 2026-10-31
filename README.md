# gopurs-node-fs

## Local Go development

This checkout is part of the gopurs library family. Use the
[local Go development guide](../gopurs/README.md#develop-one-library-locally)
for toolchain setup, sibling dependencies, Spago configuration and Go commands.
The existing npm, Bower and Dhall commands below retain their JavaScript or
upstream roles.

Default `createReadStream` / `createWriteStream` preserve the opened file and
open errors through the paired `gopurs-node-streams` native adapters. The pipe
regression copies arbitrary multichunk UTF-8 content and an empty file, waits for
destination completion, and asserts exact contents. The Go test wrapper rejects
zero exits without both completed assertions. This change does not implement
the options-based or file-descriptor stream constructors.


[![Latest release](http://img.shields.io/github/release/purescript-node/purescript-node-fs.svg)](https://github.com/purescript-node/purescript-node-fs/releases)
[![Build status](https://github.com/purescript-node/purescript-node-fs/workflows/CI/badge.svg?branch=master)](https://github.com/purescript-node/purescript-node-fs/actions?query=workflow%3ACI+branch%3Amaster)
[![Pursuit](https://pursuit.purescript.org/packages/purescript-node-fs/badge)](https://pursuit.purescript.org/packages/purescript-node-fs)

PureScript bindings to node's `fs` module.

## Installation

```
spago install node-fs
```

## Documentation

Module documentation is [published on Pursuit](http://pursuit.purescript.org/packages/purescript-node-fs).
