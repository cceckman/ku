load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "verify",
    srcs = ["main.go"],
    deps = ["//sudoku:go_default_library"],
)

sh_test(
    name = "suite_a_test",
    srcs = ["suite_a_test.sh"],
    data = [
        ":verify",
        "//testdata:suite-a",
    ],
)
