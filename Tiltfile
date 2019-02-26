def get_username():
  return str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (repr(get_username()), repr(file)))

repo = local_git_repo('.')

k8s_yaml(m4_yaml('deploy/hypothesizer.yaml'))

hfb = custom_build(
  'gcr.io/windmill-public-containers/servantes/hypothesizer',
  'docker build -t $TAG hypothesizer',
  ['hypothesizer']).add_fast_build()
hfb.add(repo.path('hypothesizer'), '/app')
hfb.run('cd /app && pip install -r requirements.txt', trigger='hypothesizer/requirements.txt')
hfb.hot_reload()

k8s_resource('hypothesizer', image=hfb, port_forwards=5000)
