# -*- mode: Python -*-

"""
This Tiltfile contains 1 composite service which depends on a number of regular services.
Here's a quick rundown of these services and their properties:

* Frontend
  * Language: Go
  * Other notes: presents a grid of the results of calling all of the other services
* Vigoda
  * Language: Go
* Snack
  * Language: Go
* Doggos
  * Language: Go
* Fortune
  * Language: Go
* Hypothesizer
  * Language: Python
  * Other notes: does a `pip install` for package dependencies. Reinstalls dependencies, only if the dependencies have changed.
* Spoonerisms
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies, only if the dependencies have changed
"""

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms])

def get_username():
  return local('whoami').rstrip('\n')

def kustomized_yaml(dir):
  have_kustomize = local('kustomize > /dev/null || echo not found')
  if len(have_kustomize) > 0:
    print("servantes requires kustomize. Install via: go get sigs.k8s.io/kustomize")
  return local('kustomize build ' + dir + ' | m4 -DOWNER="' + get_username() + '"')

def fe():
  yaml = kustomized_yaml('fe/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/fe'

  start_fast_build('Dockerfile.go.base', image_name, '/go/bin/fe --owner ' + get_username())
  path = '/go/src/github.com/windmilleng/servantes/fe'
  repo = local_git_repo('.')
  add(repo.path('fe'), path)

  run('go install github.com/windmilleng/servantes/fe')
  img = stop_build()

  return k8s_service(yaml, img)

def vigoda():
  yaml = kustomized_yaml('vigoda/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/vigoda'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/vigoda'
  repo = local_git_repo('.')
  add(repo.path('vigoda'), path)

  run('go install github.com/windmilleng/servantes/vigoda')
  img = stop_build()

  return k8s_service(yaml, img)

def snack():
  yaml = kustomized_yaml('snack/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/snack'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/snack'
  repo = local_git_repo('.')
  add(repo.path('snack'), path)

  run('go install github.com/windmilleng/servantes/snack')
  img = stop_build()

  return k8s_service(yaml, img)

def doggos():
  yaml = kustomized_yaml('doggos/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/doggos'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/doggos'
  repo = local_git_repo('.')
  add(repo.path('doggos'), path)

  run('go install github.com/windmilleng/servantes/doggos')
  img = stop_build()

  return k8s_service(yaml, img)

def fortune():
  yaml = kustomized_yaml('fortune/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/fortune'

  img = start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/fortune'
  repo = local_git_repo('.')
  add(repo.path('fortune'), path)

  run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  run('go install github.com/windmilleng/servantes/fortune')
  img = stop_build()

  return k8s_service(yaml, img)

def hypothesizer():
  yaml = kustomized_yaml('hypothesizer/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/hypothesizer'

  start_fast_build('Dockerfile.py.base', image_name)
  repo = local_git_repo('.')
  add(repo.path('hypothesizer'), "/app")

  run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')
  img = stop_build()

  return k8s_service(yaml, img)

def spoonerisms():
  yaml = kustomized_yaml('spoonerisms/deployments')

  image_name = 'gcr.io/windmill-public-containers/servantes/spoonerisms'

  start_fast_build('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  add(repo.path('spoonerisms/src'), '/app')
  add(repo.path('spoonerisms/package.json'), '/app/package.json')
  add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')

  run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])
  img = stop_build()

  return k8s_service(yaml, img)
