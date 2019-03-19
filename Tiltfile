# -*- mode: Python -*-

username = str(local('whoami')).rstrip('\n')

def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (username, repr(file)))

k8s_yaml(m4_yaml('deploy/snack.yaml'))

## Part 2: Images

# most services we do docker_builds
docker_build('snack:tagged', 'snack').add('./snack/web', '/go/src/github.com/windmilleng/servantes/snack')
k8s_resource('snack', port_forwards=9000)
