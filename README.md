# Basic go bazel project

[![Test status](https://github.com/filmil/bazel-go-basic/workflows/Test/badge.svg)](https://github.com/filmil/bazel-go-basic/workflows/Test/badge.svg)

This is an empty go project that you can use for spinning off your own projects
that use `bazel` as a build system, and the go toolchain.  Of course you can add
other toolchains you need as your project grows.

## Features

- `bazel` fixed to a reasonably recent version.
- `.bazelrc` with support for `user.bazelrc`
- `.bazelrc` with support for a 3rd party public registry.
- Go toolchain set up with the latest go version at the moment
- Github test workflow with artifact caching.
- Github source and binary release workflow on release cut.
- LICENSE is present (Apache 2.0).
- README is present.
- `bazel build //...` works
- `bazel test //...` passes
- Workspace uses bzlmod (i.e. `MODULE.bazel` instead of `WORKSPACE`).
- Use `bazel run //:gazelle` to tidy up your build files.
- Use `bazel run //:buildifier` to format things.
- `.gitignore` ignores bazel ephemeral files.
