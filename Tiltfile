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
  * Other notes: Has a JS component
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

def get_username():
  return str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (repr(get_username()), repr(file)))

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
  'deploy/sleeper.yaml'
]

k8s_yaml([m4_yaml(f) for f in yamls])

## Part 2: Images

# most services we do docker_builds
docker_build('gcr.io/windmill-public-containers/servantes/vigoda', 'vigoda')
docker_build('gcr.io/windmill-public-containers/servantes/snack', 'snack')
docker_build('gcr.io/windmill-public-containers/servantes/doggos', 'doggos')
docker_build('gcr.io/windmill-public-containers/servantes/emoji', 'emoji')
docker_build('gcr.io/windmill-public-containers/servantes/words', 'words')
docker_build('gcr.io/windmill-public-containers/servantes/secrets', 'secrets')
docker_build('gcr.io/windmill-public-containers/servantes/sleep', 'sleeper')

# fast builds show how we can handle complex cases quickly
(fast_build('gcr.io/windmill-public-containers/servantes/fe',
            'Dockerfile.go.base', '/go/bin/fe --owner ' + get_username())
  .add(repo.path('fe'), '/go/src/github.com/windmilleng/servantes/fe')
  .run('go install github.com/windmilleng/servantes/fe'))
(fast_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'Dockerfile.py.base')
  .add(repo.path('hypothesizer'), '/app')
  .run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt'))
(fast_build('gcr.io/windmill-public-containers/servantes/fortune', 'Dockerfile.go.base')
  .add(repo.path('fortune'), '/go/src/github.com/windmilleng/servantes/fortune')
  .run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  .run('go install github.com/windmilleng/servantes/fortune'))
(fast_build('gcr.io/windmill-public-containers/servantes/spoonerisms', 'Dockerfile.js.base', 'node /app/index.js')
  .add(repo.path('spoonerisms/src'), '/app')
  .add(repo.path('spoonerisms/package.json'), '/app/package.json')
  .add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')
  .run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock']))

## Part 3: Resources
def add_ports(): # we want to add local ports to each service, starting at 9000
  port = 9000
  for name in ['fe', 'vigoda', 'snack', 'doggos', 'fortune', 'hypothesizer', 'spoonerisms', 'emoji', 'words', 'secrets']:
    k8s_resource(name, port_forwards=port)
    port += 1

add_ports()
