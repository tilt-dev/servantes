load("@io_bazel_rules_docker//go:image.bzl", "go_image")

package(default_visibility = ["//visibility:public"])

go_image(
    name = "image",
    srcs = ["main.go"],
    importpath = "github.com/windmilleng/servantes/vigoda",
    base = "@gobase//image",
)
