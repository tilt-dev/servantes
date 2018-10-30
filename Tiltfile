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

gy = read_file("hellokubernetes.yaml")
global_yaml(gy)

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms, emoji])

def get_username():
  return local('whoami').rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -DOWNER=%s %s' % (repr(get_username()), repr(file)))

def fe():
  yaml = m4_yaml('fe/deployments/fe.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/fe'

  start_fast_build('Dockerfile.go.base', image_name, '/go/bin/fe --owner ' + get_username())
  path = '/go/src/github.com/windmilleng/servantes/fe'
  repo = local_git_repo('.')
  add(repo.path('fe'), path)

  run('go install github.com/windmilleng/servantes/fe')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9000)
  return s

def vigoda():
  yaml = m4_yaml('vigoda/deployments/vigoda.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/vigoda'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/vigoda'
  repo = local_git_repo('.')
  add(repo.path('vigoda'), path)

  run('go install github.com/windmilleng/servantes/vigoda')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9001)
  return s

def snack():
  yaml = m4_yaml('snack/deployments/snack.yaml')
  img = static_build('snack/Dockerfile',
                     'gcr.io/windmill-public-containers/servantes/snack')
  s = k8s_service(yaml, img)
  s.port_forward(9002)
  return s

def doggos():
  yaml = m4_yaml('doggos/deployments/doggos.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/doggos'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/doggos'
  repo = local_git_repo('.')
  add(repo.path('doggos'), path)

  run('go install github.com/windmilleng/servantes/doggos')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9003)
  return s

def fortune():
  yaml = m4_yaml('fortune/deployments/fortune.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/fortune'

  img = start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/fortune'
  repo = local_git_repo('.')
  add(repo.path('fortune'), path)

  run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  run('go install github.com/windmilleng/servantes/fortune')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9004)
  return s

def hypothesizer():
  yaml = m4_yaml('hypothesizer/deployments/hypothesizer.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/hypothesizer'

  start_fast_build('Dockerfile.py.base', image_name)
  repo = local_git_repo('.')
  add(repo.path('hypothesizer'), "/app")

  run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9005)
  return s

def spoonerisms():
  yaml = m4_yaml('spoonerisms/deployments/spoonerisms.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/spoonerisms'

  start_fast_build('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  add(repo.path('spoonerisms/src'), '/app')
  add(repo.path('spoonerisms/package.json'), '/app/package.json')
  add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')

  run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9006)
  return s

def emoji():
  yaml = m4_yaml('emoji/deployments/emoji.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/emoji'

  start_fast_build('Dockerfile.go.base', image_name)
  path = '/go/src/github.com/windmilleng/servantes/emoji'
  repo = local_git_repo('.')
  add(repo.path('emoji'), path)

  run('go install github.com/windmilleng/servantes/emoji')
  img = stop_build()

  s = k8s_service(yaml, img)
  s.port_forward(9007)
  return s
