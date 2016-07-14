load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(
    default_visibility = ["//:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = ["puzzle.go", "collection.go"],
)

go_test(
    name = "puzzle_test",
    srcs = ["puzzle_test.go"],
    library = ":go_default_library",
    size = "small",
    )
