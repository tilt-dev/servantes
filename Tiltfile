# -*- mode: Python -*-

def servantes():
  return composite_service([fe, vigoda, fortune, doggos, snack, hypothesizer, spoonerisms])

def fe():
  yaml = read_file('servantes/deployments/servantes.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/servantes'

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/servantes')
  path = '/go/src/github.com/windmilleng/servantes/servantes'
  repo = local_git_repo('.')
  img.add(repo.path('servantes/'), path)

  img.run('go install github.com/windmilleng/servantes/servantes')
  return k8s_service(yaml, img)

def vigoda():
  yaml = read_file('vigoda/deployments/vigoda.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/vigoda'

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/vigoda')
  path = '/go/src/github.com/windmilleng/servantes/vigoda'
  repo = local_git_repo('.')
  img.add(repo.path('vigoda/'), path)

  img.run('go install github.com/windmilleng/servantes/vigoda')

  return k8s_service(yaml, img)

def snack():
  yaml = read_file('snack/deployments/snack.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/snack'

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/snack')
  path = '/go/src/github.com/windmilleng/servantes/snack'
  repo = local_git_repo('.')
  img.add(repo.path('snack/'), path)

  img.run('go install github.com/windmilleng/servantes/snack')

  return k8s_service(yaml, img)

def doggos():
  yaml = read_file('doggos/deployments/doggos.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/doggos'

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/doggos')
  path = '/go/src/github.com/windmilleng/servantes/doggos'
  repo = local_git_repo('.')
  img.add(repo.path('doggos/'), path)

  img.run('go install github.com/windmilleng/servantes/doggos')

  return k8s_service(yaml, img)

def fortune():
  yaml = read_file('fortune/deployments/fortune.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/fortune'

  img = build_docker_image('Dockerfile.go.base', image_name, '/go/bin/fortune')
  path = '/go/src/github.com/windmilleng/servantes/fortune'
  repo = local_git_repo('.')
  img.add(repo.path('fortune/'), path)

  img.run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  img.run('go install github.com/windmilleng/servantes/fortune')

  return k8s_service(yaml, img)

def hypothesizer():
  yaml = read_file('hypothesizer/deployments/hypothesizer.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/hypothesizer'

  img = build_docker_image('Dockerfile.py.base', image_name, 'python /app/app.py')
  repo = local_git_repo('.')
  img.add(repo.path('hypothesizer/'), "/app")

  img.run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')

  return k8s_service(yaml, img)

def spoonerisms():
  yaml = read_file('spoonerisms/deployments/spoonerisms.yaml')

  image_name = 'gcr.io/windmill-public-containers/servantes/spoonerisms'

  img = build_docker_image('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  img.add(repo.path('spoonerisms/src'), '/app/')
  img.add(repo.path('spoonerisms/package.json'), '/app/package.json')
  img.add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')

  img.run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock'])

  return k8s_service(yaml, img)
