load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "verify",
    srcs = ["main.go"],
    deps = ["//puzzle:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["verify.go"],
    deps = ["//puzzle:go_default_library"],
)

go_test(
    name = "verify_test",
    srcs = ["verify_test.go"],
    data = [
        "//testdata:suite-a",
    ],
    library = ":go_default_library",
    deps = [
        "//puzzle:go_default_library",
    ],
)
