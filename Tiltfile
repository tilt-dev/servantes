# -*- mode: Python -*-

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms])

def go_service(name, extra_runs=[]):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  # right now, Servantes is only intended to work with local docker-for-desktop
  # or minikube, so we just make up an image name
  image_name = 'windmill.build/servantes/%s' % name

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/%s' % name)
  path = '/go/src/github.com/windmilleng/servantes/%s' % name
  repo = local_git_repo('./%s/' % name)
  img.add(path, repo)

  for r in extra_runs:
    img.run(r)

  img.run('go install github.com/windmilleng/servantes/%s' % name)
  return k8s_service(yaml, img)

def python_service(name, extra_runs=[]):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  # right now, Servantes is only intended to work with local docker-for-desktop
  # or minikube, so we just make up an image name
  image_name = 'windmill.build/servantes/%s' % name

  img = build_docker_image('Dockerfile.py.base', image_name, 'python /app/app.py')
  repo = local_git_repo('./%s/' % name)
  img.add("/app", repo)

  for r in extra_runs:
    img.run(r)

  return k8s_service(yaml, img)

def javascript_service(name, extra_runs=[]):
  yaml = local_file('%s/deployments/%s.yaml' % (name, name))

  # right now, Servantes is only intended to work with local docker-for-desktop
  # or minikube, so we just make up an image name
  image_name = 'windmill.build/servantes/%s' % name

  img = build_docker_image('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('./%s/' % name)
  img.add("/app", repo)

  for r in extra_runs:
    img.run(r)

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
  return go_service('fortune', ['cd src/github.com/windmilleng/servantes/fortune && make proto'])

def hypothesizer():
  return python_service('hypothesizer', ['cd /app && pip install -r requirements.txt'])

def spoonerisms():
  return javascript_service('spoonerisms', ['cd /app && yarn install'])

def local_file(p):
  return local("cat %s" % p)
