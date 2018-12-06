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

def get_username():
  return str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (repr(get_username()), repr(file)))

repo = local_git_repo('.')
k8s_yaml(m4_yaml('./global.yaml'))

# Frontend
# TODO(dmiller): cache
k8s_resource('fe', m4_yaml('fe/deployments/fe.yaml'), port_forwards=9000)
fe_img = 'gcr.io/windmill-public-containers/servantes/fe'
(fast_build(fe_img, 'Dockerfile.go.base', '/go/bin/fe --owner ' + get_username())
  .add(repo.path('fe'), '/go/src/github.com/windmilleng/servantes/fe')
  .run('go install github.com/windmilleng/servantes/fe'))

# Vigoda
k8s_resource('vigoda', m4_yaml('vigoda/deployments/vigoda.yaml'), port_forwards=9001)
vigoda_img = 'gcr.io/windmill-public-containers/servantes/vigoda'
(fast_build(vigoda_img, 'Dockerfile.go.base')
  .add(repo.path('vigoda'), '/go/src/github.com/windmilleng/servantes/vigoda')
  .run('go install github.com/windmilleng/servantes/vigoda'))

# Snack
k8s_resource('snack', m4_yaml('snack/deployments/snack.yaml'), port_forwards=9002)
docker_build('gcr.io/windmill-public-containers/servantes/snack', 'snack')

# Doggos
k8s_resource('doggos', m4_yaml('doggos/deployments/doggos.yaml'), port_forwards=9003)
doggos_img = 'gcr.io/windmill-public-containers/servantes/doggos'
(fast_build(doggos_img, 'Dockerfile.go.base')
  .add(repo.path('doggos'), '/go/src/github.com/windmilleng/servantes/doggos')
  .run('go install github.com/windmilleng/servantes/doggos'))

# Fortune
k8s_resource('fortune', m4_yaml('fortune/deployments/fortune.yaml'), port_forwards=9004)
fortune_img = 'gcr.io/windmill-public-containers/servantes/fortune'
(fast_build(fortune_img, 'Dockerfile.go.base')
  .add(repo.path('fortune'), '/go/src/github.com/windmilleng/servantes/fortune')
  .run('cd src/github.com/windmilleng/servantes/fortune && make proto')
  .run('go install github.com/windmilleng/servantes/fortune'))

# Hypothesizer
k8s_resource('hypothesizer', m4_yaml('hypothesizer/deployments/hypothesizer.yaml'), port_forwards=9005)
hyp_img = 'gcr.io/windmill-public-containers/servantes/hypothesizer'
(fast_build(hyp_img, 'Dockerfile.py.base')
  .add(repo.path('hypothesizer'), '/app')
  .run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt'))

# Spoonerisms
k8s_resource('spoonerisms', m4_yaml('spoonerisms/deployments/spoonerisms.yaml'), port_forwards=9006)
spoonerism_img = 'gcr.io/windmill-public-containers/servantes/spoonerisms'
(fast_build(spoonerism_img, 'Dockerfile.js.base', 'node /app/index.js')
  .add(repo.path('spoonerisms/src'), '/app')
  .add(repo.path('spoonerisms/package.json'), '/app/package.json')
  .add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock')
  .run('cd /app && yarn install', trigger=['spoonerisms/package.json', 'spoonerisms/yarn.lock']))

# Emoji
k8s_resource('emoji', m4_yaml('emoji/deployments/emoji.yaml'), port_forwards=9007)
emoji_img = 'gcr.io/windmill-public-containers/servantes/emoji'
(fast_build(emoji_img, 'Dockerfile.go.base')
  .add(repo.path('emoji'), '/go/src/github.com/windmilleng/servantes/emoji')
  .run('go install github.com/windmilleng/servantes/emoji'))


# Words
k8s_resource('words', m4_yaml('words/deployments/words.yaml'), port_forwards=9008)
words_img = 'gcr.io/windmill-public-containers/servantes/words'
(fast_build(words_img, 'Dockerfile.py.base')
  .add(repo.path('words'), '/app')
  .run('cd /app && pip install -r requirements.txt', trigger='words/requirements.txt'))

# Secrets
secrets_img = 'gcr.io/windmill-public-containers/servantes/secrets'
sfb = (fast_build(secrets_img, 'Dockerfile.go.base')
  .add(repo.path('secrets'), '/go/src/github.com/windmilleng/servantes/secrets')
  .run('go install github.com/windmilleng/servantes/secrets'))

k8s_resource('secret', image=sfb, port_forwards=9009)
