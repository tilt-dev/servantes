# -*- mode: Python -*-

repo = local_git_repo('.')

k8s_yaml('deploy/hypothesizer.yaml')

# docker_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'hypothesizer')
# hfb = docker_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'hypothesizer').add_fast_build()
# hfb.add(repo.path('hypothesizer'), '/app')
docker_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'hypothesizer')


# docker_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'hypothesizer')
# custom_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'docker build -t $TAG hypothesizer', ['hypothesizer'])
# hfb = custom_build('gcr.io/windmill-public-containers/servantes/hypothesizer', 'docker build -t $TAG hypothesizer', ['hypothesizer']).add_fast_build()
# hfb.add(repo.path('hypothesizer'), '/app').run('sleep 3')


k8s_resource('hypothesizer', port_forwards=9005)
