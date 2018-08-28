# -*- mode: Python -*-

repo = local_git_repo('./')
path = '/go/src/github.com/windmilleng/servantes'

def servantes():
  return composite_service([fe, vigoda, fortune])

def service(name, extra_runs=[]):
  yaml = local_file('deployments/%s/%s.yaml' % (name, name))

  # right now, Servantes is only intended to work with local docker-for-desktop
  # or minikube, so we just make up an image name
  image_name = 'windmill.build/servantes/%s' % name

  img = build_docker_image('Dockerfile.base', image_name, '/go/bin/%s' % name)
  img.add(path, repo)

  for r in extra_runs:
    img.run(r)

  img.run('go install github.com/windmilleng/servantes/cmd/%s' % name)
  return k8s_service(yaml, img)

def fe():
  return service('servantes')

def vigoda():
  return service('vigoda')

def fortune():
  return service('fortune', ['cd src/github.com/windmilleng/servantes && make proto'])

def local_file(p):
  return local("cat %s" % p)