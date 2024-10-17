prepare-environment:
	(docker build . -t ansible_server:latest -f ansible/Dockerfile && docker build . -t jenkins_server:latest -f jenkins/Dockerfile && docker build . -t k8hosts:latest -f kubernetes/Dockerfile)
