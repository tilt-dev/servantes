# -*- mode: Python -*-

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms])

def go_service(name, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  # right now, Servantes is only intended to work with local docker-for-desktop
  # or minikube, so we just make up an image name
  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/%s' % name)
  path = '/go/src/github.com/windmilleng/servantes/%s' % name
  repo = local_git_repo('./%s/' % name)
  img.add(repo, path)

  if run_fn:
    run_fn(img)

  img.run('go install github.com/windmilleng/servantes/%s' % name)
  return k8s_service(yaml, img)

def python_service(name, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.py.base', image_name, 'python /app/app.py')
  repo = local_git_repo('./%s/' % name)
  img.add(repo, "/app")

  if run_fn:
    run_fn(img)

  return k8s_service(yaml, img)

def javascript_service(name, run_fn=None):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  image_name = 'gcr.io/windmill-public-containers/servantes/%s' % name

  img = build_docker_image('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('./%s/' % name)
  img.add(repo, "/app")

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

  return javascript_service('spoonerisms', runs)

def local_file(p):
  return local("cat %s" % p)
