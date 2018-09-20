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

def fe():
  yaml = read_file('fe/deployments/fe.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/fe'

  img = build_docker_image('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/fe'
  repo = local_git_repo('.')
  img.add(repo.path('fe'), path)

  img.run('go install github.com/windmilleng/servantes/fe')
  return k8s_service(yaml, img)

def vigoda():
  yaml = read_file('vigoda/deployments/vigoda.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/vigoda'

  img = build_docker_image('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/vigoda'
  repo = local_git_repo('.')
  img.add(repo.path('vigoda'), path)

  img.run('go install github.com/windmilleng/servantes/vigoda')

  return k8s_service(yaml, img)

def snack():
  yaml = read_file('snack/deployments/snack.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/snack'

  img = build_docker_image('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/snack'
  repo = local_git_repo('.')
  img.add(repo.path('snack'), path)

  img.run('go install github.com/windmilleng/servantes/snack')

  return k8s_service(yaml, img)

def doggos():
  yaml = read_file('doggos/deployments/doggos.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/doggos'

  img = build_docker_image('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/doggos'
  repo = local_git_repo('.')
  img.add(repo.path('doggos'), path)

  img.run('go install github.com/windmilleng/servantes/doggos')

  return k8s_service(yaml, img)

def fortune():
  yaml = read_file('fortune/deployments/fortune.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/fortune'

  img = build_docker_image('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/fortune'
  repo = local_git_repo('.')
  img.add(repo.path('fortune'), path)

  img.run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  img.run('go install github.com/windmilleng/servantes/fortune')

  return k8s_service(yaml, img)

def hypothesizer():
  yaml = read_file('hypothesizer/deployments/hypothesizer.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/hypothesizer'

  img = build_docker_image('Dockerfile.py.base', image_name)
  repo = local_git_repo('.')
  img.add(repo.path('hypothesizer'), "/app")

  img.run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')

  return k8s_service(yaml, img)

def spoonerisms():
  yaml = read_file('spoonerisms/deployments/spoonerisms.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/spoonerisms'

  img = build_docker_image('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  img.add(repo.path('spoonerisms/src'), '/app')
  img.add(repo.path('spoonerisms/package.json'), '/app/package.json')
  img.add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')

  img.run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])

  return k8s_service(yaml, img)
