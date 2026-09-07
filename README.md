# Basic go bazel project

[![Test](https://github.com/filmil/bazel-go-basic/actions/workflows/test.yml/badge.svg)](https://github.com/filmil/bazel-go-basic/actions/workflows/test.yml)
[![Tag and Release](https://github.com/filmil/bazel-go-basic/actions/workflows/tag-and-release.yml/badge.svg)](https://github.com/filmil/bazel-go-basic/actions/workflows/tag-and-release.yml)
[![Publish on Bazel Central Registry](https://github.com/filmil/bazel-go-basic/actions/workflows/publish-bcr.yml/badge.svg)](https://github.com/filmil/bazel-go-basic/actions/workflows/publish-bcr.yml)
[![Publish to my Bazel registry](https://github.com/filmil/bazel-go-basic/actions/workflows/publish.yml/badge.svg)](https://github.com/filmil/bazel-go-basic/actions/workflows/publish.yml)

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
- Protobuf support via `rules_proto` and `go_proto_library`.
- Use `bazel run //:gazelle` to tidy up your build files.
- Use `bazel run //:buildifier` to format things.
- `.gitignore` ignores bazel ephemeral files.
- A setup for a coding assistant. (It's fashionable at the moment!)
- A setup for auto-publishing the module to Bazel central registry, and a
  privately owned, but public, secondary registry.
- A workflow that cuts a new release each week, and publishes to my bazel registry.
- A workflow that gives an option to publish to BCR.
- A registry test module in `integration/`, which the Bazel Central
  Registry presubmit builds against the published archive. Run it locally
  with `cd integration && bazel test //...`.

## Documentation

- [AI Assistant BCR Publishing Rules](ai/AI_BCR.md)
- [AI Assistant Git Rules](ai/AI_GIT.md)
- [Gemini Configuration](GEMINI.md)
- [Hello BZL Reference](hello.md)
