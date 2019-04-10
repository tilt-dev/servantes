# -*- mode: Python -*-

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
  'deploy/fe.yaml',
  'deploy/vigoda.yaml',
  'deploy/snack.yaml',
  'deploy/doggos.yaml',
  'deploy/fortune.yaml',
  'deploy/hypothesizer.yaml',
  'deploy/spoonerisms.yaml',
  'deploy/emoji.yaml',
  'deploy/words.yaml',
  'deploy/secrets.yaml',
  'deploy/job.yaml',
  'deploy/sleeper.yaml',
  'deploy/hello_world.yaml',
  'deploy/tick.yaml',
]

k8s_yaml([m4_yaml(f) for f in yamls])

## Part 2: Images

# most services are just a simple docker_build
docker_build('vigoda', 'vigoda')
docker_build('snack', 'snack')
docker_build('doggos', 'doggos')
docker_build('emoji', 'emoji')
docker_build('words', 'words')
docker_build('secrets', 'secrets')
docker_build('sleep', 'sleeper')
docker_build('sidecar', 'sidecar')

# we can add live_update steps on top of a docker_build for super fast in-place updates
docker_build('fe', 'fe',
  live_update=[
    sync('fe', '/go/src/github.com/windmilleng/servantes/fe'),
    run('go install github.com/windmilleng/servantes/fe'),
    restart_container(),
  ]
)

docker_build('hypothesizer', 'hypothesizer',
  live_update=[
    sync('hypothesizer', '/app'),
    run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt'),
    # no restart_container needed because hypothesizer is a flask app which hot-reloads its code
  ]
)
docker_build('fortune', 'fortune',
  live_update=[
    sync('fortune', '/go/src/github.com/windmilleng/servantes/fortune'),
    run('cd src/github.com/windmilleng/servantes/fortune && make proto'),
    run('go install github.com/windmilleng/servantes/fortune'),
    restart_container(),
  ]
)
docker_build('spoonerisms', 'spoonerisms',
  live_update=[
    sync('spoonerisms/src', '/app'),
    sync('spoonerisms/package.json', '/app/package.json'),
    sync('spoonerisms/yarn.lock', '/app/yarn.lock'),
    run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock']),
    restart_container(),
  ]
)

## Part 3: Resources
def add_ports(): # we want to add local ports to each service, starting at 9000
  port = 9000
  for name in ['fe', 'vigoda', 'snack', 'doggos', 'fortune', 'hypothesizer', 'spoonerisms', 'emoji', 'words', 'secrets']:
    k8s_resource(name, port_forwards=port)
    port += 1

add_ports()

## Part 4: other use cases

# here's a k8s_resource with only YAML and no associated docker_build that we
# can still port-forward. You can run, manipulate, and see logs for k8s
# resources out of the box!
k8s_resource('hello-world', port_forwards=9999)

# strip off the $USER- that we prepend to all deployment names
def resource_name(id):
  return id.name.replace(username + '-', '')
workload_to_resource_function(resource_name)
