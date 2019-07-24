# -*- mode: Python -*-

k8s_yaml(['crd.yaml', 'widget.yaml'])
k8s_kind('widget', image_json_path='{.location.city}')

docker_build('localhost:32001/postgres', 'vigoda')
