load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_bundle")

package(default_visibility = ["//visibility:public"])

go_image(
    name = "image",
    srcs = ["main.go"],
    importpath = "github.com/windmilleng/servantes/vigoda",
)

container_bundle(
  name = "tilt",
  images = {
    "gcr.io/windmill-public-containers/servantes/vigoda:tilt-bazel": "image",
  }
)
