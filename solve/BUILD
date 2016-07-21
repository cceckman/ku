load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(
    default_visibility = ["//:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = ["solve.go"],
    deps = [
        "//puzzle:go_default_library",
        "//verify:go_default_library",
    ],
)

go_test(
    name = "solve_test",
    size = "small",
    srcs = ["solve_test.go"],
    library = ":go_default_library",
    deps = [
        "//puzzle:go_default_library",
        "//verify:go_default_library",
    ],
)
