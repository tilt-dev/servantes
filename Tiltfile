# -*- mode: Python -*-

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


# Read an option from the file tilt_option.json
# Create this file with json like { "default_registry": "gcr.io/my-personal-project" }
# tilt_option.json is .gitignore'd, so you'll have a clean status
default_registry(read_json('tilt_option.json', {})
                 .get('default_registry', 'gcr.io/windmill-public-containers/servantes'))

username = str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (username, repr(file)))

repo = local_git_repo('.')

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

# most services we do docker_builds
docker_build('vigoda', 'vigoda')
docker_build('snack', 'snack')
docker_build('doggos', 'doggos')
docker_build('emoji', 'emoji')
docker_build('words', 'words')
docker_build('secrets', 'secrets')
docker_build('sleep', 'sleeper')
docker_build('sidecar', 'sidecar')

# fast builds show how we can handle complex cases quickly
(fast_build('fe',
            'Dockerfile.go.base', '/go/bin/fe --owner ' + username)
  .add(repo.path('fe'), '/go/src/github.com/windmilleng/servantes/fe')
  .run('go install github.com/windmilleng/servantes/fe'))
(fast_build('hypothesizer', 'Dockerfile.py.base')
  .add(repo.path('hypothesizer'), '/app')
  .run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt'))
(fast_build('fortune', 'Dockerfile.go.base')
  .add(repo.path('fortune'), '/go/src/github.com/windmilleng/servantes/fortune')
  .run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  .run('go install github.com/windmilleng/servantes/fortune'))

# You can also ADD fast build instructions to a normal docker build; we use the fast build
# instructions to update the live container when possible, otherwise use the docker build
# instructions for a fresh build + deploy
spoonerisms = docker_build('gcr.io/windmill-public-containers/servantes/spoonerisms', 'spoonerisms')
spoonerisms.add(repo.path('spoonerisms/src'), '/app')
spoonerisms.add(repo.path('spoonerisms/package.json'), '/app/package.json')
spoonerisms.add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')
spoonerisms.run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])

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
