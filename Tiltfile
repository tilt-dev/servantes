# repros https://github.com/tilt-dev/tilt-extensions/issues/391
# 1. tilt up
# 2. edit fortune/main.go
# 3. observe:
#   a. the process isn't restarted
#   b. the deployment spec's command is "/go/bin/fortune" rather than the restart_process wrapper

load('ext://restart_process', 'docker_build_with_restart')

k8s_custom_deploy(
  'fortune-depl',
  apply_cmd='m4 -Dvarowner="$(whoami)" deploy/fortune.yaml | sed -e"s|image: fortune|image: ${TILT_IMAGE_0}|" | kubectl apply -f - -oyaml',
  delete_cmd='m4 -Dvarowner="$(whoami)" deploy/fortune.yaml | kubectl delete -f -',
  image_deps=['fortune'],
  deps='deploy/fortune.yaml')

docker_build_with_restart('fortune', 'fortune', '/go/bin/fortune',
  live_update=[
    sync('fortune', '/go/src/github.com/tilt-dev/servantes/fortune'),
    run('cd src/github.com/tilt-dev/servantes/fortune && make proto'),
    run('cd src/github.com/tilt-dev/servantes/fortune && go install .'),
  ]
)
