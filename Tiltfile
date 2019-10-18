# -*- mode: Python -*-
enable_feature("snapshots")

k8s_resource_assembly_version(2)

"""
This Tiltfile contains one external-facing service which depends on a number of internal services.
Here's a quick rundown of these services and their properties:

* Frontend
  * Language: Go
  * Other notes: presents a grid of the results of calling all of the other services
* Vigoda
  * Language: Go
* Snack
  * Language: Go
  * Other notes: Uses static_build
* Doggos
  * Language: Go
  * Other notes: Has a JS component, and a sidecar that yells a lot
* Fortune
  * Language: Go
  * Other notes: Uses protobufs
* Hypothesizer
  * Language: Python
  * Other notes: does a `pip install` for package dependencies. Reinstalls dependencies, only if the dependencies have changed.
* Spoonerisms
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies, only if the dependencies have changed
"""

enable_feature("events")

# If you get push errors, you can change the default_registry.
# Create tilt_option.json with contents: {"default_registry": "gcr.io/my-personal-project"}
# (with your registry inserted). tilt_option.json is gitignore'd, unlike Tiltfile
default_registry(read_json('tilt_option.json', {})
                 .get('default_registry', 'gcr.io/windmill-public-containers/servantes'))

username = str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (username, repr(file)))

## Part 1: kubernetes yamls
yamls = [
  'deploy/vigoda.yaml',
]

k8s_yaml([m4_yaml(f) for f in yamls])

## Part 2: Images

# most services are just a simple docker_build
docker_build('vigodabase', 'vigoda', dockerfile='./vigoda/Dockerfile.base')
docker_build('vigoda', 'vigoda', dockerfile='./vigoda/Dockerfile.child')

k8s_resource('vigoda', port_forwards=9000)

# strip off the $USER- that we prepend to all deployment names
def resource_name(id):
  return id.name.replace(username + '-', '')
workload_to_resource_function(resource_name)
