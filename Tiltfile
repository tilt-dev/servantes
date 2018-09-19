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
  * Other notes: does a `pip install` for package dependencies. Only reinstalls dependencies, only if the dependencies have changed.
* Spoonerisms
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies, only if the dependencies have changed
"""

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms])

def go_service(name, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/%s' % name)
  path = '/go/src/github.com/windmilleng/servantes/%s' % name
  repo = local_git_repo('.')
  img.add(repo.path('./%s/' % name), path)

  if run_fn:
    run_fn(img)

  img.run('go install github.com/windmilleng/servantes/%s' % name)
  return k8s_service(yaml, img)

def python_service(name, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.py.base', image_name, 'python /app/app.py')
  repo = local_git_repo('.')
  img.add(repo.path('./%s/' % name), "/app")

  if run_fn:
    run_fn(img)

  return k8s_service(yaml, img)

def javascript_service(name, dir_with_code, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  img.add(repo.path('./%s/%s' % (name, dir_with_code)), '/app/')
  img.add(repo.path('./%s/package.json' % name), '/app/index.js')
  img.add(repo.path('./%s/yarn.lock' % name), '/app/yarn.lock')
  img.add(repo.path('./%s/' % name), "/app")

  if run_fn:
    run_fn(img)

  return k8s_service(yaml, img)

def fe():
  return go_service('servantes')

def vigoda():
  return go_service('vigoda')

def snack():
  return go_service('snack')

def doggos():
  return go_service('doggos')

def fortune():
  def runs(img):
    img.run('cd src/github.com/windmilleng/servantes/fortune && make proto')

  return go_service('fortune', runs)

def hypothesizer():
  def runs(img):
    img.run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')

  return python_service('hypothesizer', runs)

def spoonerisms():
  def runs(img):
    img.run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])

  return javascript_service('spoonerisms', 'src', runs)

def local_file(p):
  return local("cat %s" % p)
