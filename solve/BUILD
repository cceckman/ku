load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

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
    srcs = ["solve_test.go"],
    library = ":go_default_library",
    deps = [
        "//puzzle:go_default_library",
        "//verify:go_default_library",
    ],
    size = "small",
)
