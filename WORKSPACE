# Support Golang.
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.0.3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

bind(
    name = "rules_go.bzl",
    actual = "@io_bazel_rules_go//go:def.bzl",
)


