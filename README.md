# bazel-differ

`bazel-differ` is a command line interface for Bazel that helps you do incremental builds across different Git versions. `bazel-differ` is a mostly a pure Go port of the excellent [bazel-diff](https://github.com/tinder/bazel-diff) and should be able to function as a drop-in replacement for most use cases.

# FAQ

1. Why did you port `bazel-diff`?

A couple reasons:

*  The projects I work with don't have any JVM dependencies so using `bazel-diff` was adding some additional build time I wanted to eliminate.
* With the exception of the Bazel server itself, much of the tooling in the Bazel ecosystem is written in Go.

2. What are the differences from `bazel-diff`?

* Due to some implementation differences, the actual hashes generated by the two programs are different
* `bazel-differ` isn't using `streamed_proto` output from Bazel query (I'm not sure if there is a Go implemntation of this?). In some minimal testing against some large repositories, `bazel-differ` still seems to outperform `bazel-diff` by 2x.
