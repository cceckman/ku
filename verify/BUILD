load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

package(
    default_visibility = ["//:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = ["verify.go"],
    deps = ["//puzzle:go_default_library"],
)

go_test(
    name = "verify_test",
    size = "small",
    srcs = ["verify_test.go"],
    data = [
        "//testdata",
    ],
    library = ":go_default_library",
    deps = [
        "//puzzle:go_default_library",
    ],
)
