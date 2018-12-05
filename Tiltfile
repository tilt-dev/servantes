# -*- mode: Python -*-

# """
# This Tiltfile contains 1 composite service which depends on a number of regular services.
# Here's a quick rundown of these services and their properties:
# * Frontend
#   * Language: Go
#   * Other notes: presents a grid of the results of calling all of the other services
# * Vigoda
#   * Language: Go
# * Snack
#   * Language: Go
#   * Other notes: Uses static_build
# * Doggos
#   * Language: Go
#   * Other notes: Has a JS component
# * Fortune
#   * Language: Go
#   * Other notes: Uses protobufs
# * Hypothesizer
#   * Language: Python
#   * Other notes: does a `pip install` for package dependencies. Reinstalls dependencies, only if the dependencies have changed.
# * Spoonerisms
#   * Language: JavaScript
#   * Other notes: Uses yarn. Does a `yarn install` for package dependencies, only if the dependencies have changed
# """


username = str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (username, file))

repo = local_git_repo('.')

yaml_files = [
  'fe/deployments/fe.yaml',
  'vigoda/deployments/vigoda.yaml',
  'snack/deployments/snack.yaml',
  'doggos/deployments/doggos.yaml',
  'fortune/deployments/fortune.yaml',
  'hypothesizer/deployments/hypothesizer.yaml',
  'spoonerisms/deployments/spoonerisms.yaml',
  'emoji/deployments/emoji.yaml',
  'global.yaml',
]

k8s_yaml([m4_yaml(f) for f in yaml_files])

# docker_build('gcr.io/windmill-public-containers/servantes/fe', 'fe')




# docker_build('gcr.io/windmill-public-containers/servantes/vigoda', 'vigoda')
# docker_build('gcr.io/windmill-public-containers/servantes/snack', 'snack')
# docker_build('gcr.io/windmill-public-containers/servantes/doggos', 'doggos')

k8s_resource('fe', port_forwards=9000)


# fe
# fast_build('gcr.io/windmill-public-containers/servantes/fe', 'Dockerfile.go.base', '/go/bin/fe --owner ' + username) \
#   .add(repo.path('fe'), '/go/src/github.com/windmilleng/servantes/fe') \
#   .run('go install github.com/windmilleng/servantes/fe')







# # vigoda
# fast_build('gcr.io/windmill-public-containers/servantes/vigoda', 'Dockerfile.go.base') \
#   .add(repo.path('vigoda'), '/go/src/github.com/windmilleng/servantes/vigoda') \
#   .run('go install github.com/windmilleng/servantes/vigoda')
# k8s_resource('vigoda', m4_yaml('vigoda/deployments/vigoda.yaml'), port_forwards=9001)

# # snack
# docker_build('gcr.io/windmill-public-containers/servantes/snack', 'snack')
# k8s_resource('snack', m4_yaml('snack/deployments/snack.yaml'), port_forwards=9002)

# # doggos
# fast_build('gcr.io/windmill-public-containers/servantes/doggos', 'Dockerfile.go.base') \
#   .add(repo.path('doggos'), '/go/src/github.com/windmilleng/servantes/doggos') \
#   .run('go install github.com/windmilleng/servantes/doggos')
# k8s_resource('doggos', m4_yaml('doggos/deployments/doggos.yaml'), port_forwards=9003)

# # fortune
# fast_build('gcr.io/windmill-public-containers/servantes/fortune', 'Dockerfile.go.base') \
#   .add(repo.path('fortune'), '/go/src/github.com/windmilleng/servantes/fortune') \
#   .run('cd src/github.com/windmilleng/servantes/fortune && make proto') \
#   .run('go install github.com/windmilleng/servantes/fortune')
# k8s_resource('fortune', m4_yaml('fortune/deployments/fortune.yaml'), port_forwards=9004)

# # hypothesizer
# fast_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'Dockerfile.py.base') \
#   .add(repo.path('hypothesizer'), '/app') \
#   .run('cd /app && pip install -r requirements.txt')
# # FIXME(dbentley): handle trigger
# k8s_resource('hypothesizer', m4_yaml('hypothesizer/deployments/hypothesizer.yaml'), port_forwards=9005)

# # spoonerisms
# fast_build('gcr.io/windmill-public-containers/servantes/spoonerisms', 'Dockerfile.js.base', 'node /app/index.js') \
#   .add(repo.path('spoonerisms/src'), '/app') \
#   .add(repo.path('spoonerisms/package.json'), '/app/package.json') \
#   .add(repo.path('spoonerisms/yarn.lock'), '/app/yarn.lock') \
#   .run('cd /app && yarn install') #FIXME(dbentley): trigger
# k8s_resource('spoonerisms', m4_yaml('spoonerisms/deployments/spoonerisms.yaml'), port_forwards=9006)

# # emoji
# fast_build('gcr.io/windmill-public-containers/servantes/emoji', 'Dockerfile.go.base') \
#   .add(repo.path('emoji'), '/go/src/github.com/windmilleng/servantes/emoji') \
#   .run('go install github.com/windmilleng/servantes/emoji')
# k8s_resource('emoji', m4_yaml('emoji/deployments/emoji.yaml'), port_forwards=9007)



# k8s_resource('general', m4_yaml('global.yaml'))


# # # words
# words_img = fast_build('gcr.io/windmill-public-containers/servantes/words', 'Dockerfile.py.base') \
#   .add(repo.path('words'), '/app') \
#   .run('cd /app && pip install -r requirements.txt')
# # FIXME(dbentley): handle trigger
# # k8s_resource('words', image=words_img, port_forwards=9008)

# # # secrets
# secrets_img = fast_build('gcr.io/windmill-public-containers/servantes/secrets', 'Dockerfile.go.base') \
#   .add(repo.path('secrets'), '/go/src/github.com/windmilleng/servantes/secrets') \
#   .run('go install github.com/windmilleng/servantes/secrets')
