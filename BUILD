load("@io_bazel_rules_k8s//k8s:object.bzl", "k8s_object")

k8s_object(
  name = "vigoda-server",
  kind = "deployment",

  # A template of a Kubernetes Deployment object yaml.
  template = ":deploy/vigoda.yaml",

  cluster = "gke_blorg-dev_us-central1-b_blorg",
)

k8s_object(
  name = "snack-server",
  kind = "deployment",

  # A template of a Kubernetes Deployment object yaml.
  template = ":deploy/snack.yaml",

  cluster = "gke_blorg-dev_us-central1-b_blorg",
)
