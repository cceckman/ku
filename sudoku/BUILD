load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(
    default_visibility = ["//:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "collection.go",
        "puzzle.go",
    ],
)

go_test(
    name = "sudoku_test",
    size = "small",
    srcs = [
        "sudoku_test.go",
    ],
    library = ":go_default_library",
)

