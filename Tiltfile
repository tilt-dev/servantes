# -*- mode: Python -*-

# curl localhost:9999 or navigate to `hello-world` in the UI and hit `b` (for "browser")
k8s_resource('hello-world', yaml='deploy/hello_world.yaml', port_forwards=9999)

### Alternately, you could do it like this:
# # add to the global pool of k8s yaml
# k8s_yaml('deploy/hello_world.yaml')
#
# # first arg is Deployment name to match by; we apply the port forward to matching k8s entities
# k8s_resource('hello-world', port_forwards=9999)
