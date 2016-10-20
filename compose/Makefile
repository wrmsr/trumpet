DOCKER_USER=willremind101

.PHONY: docker
docker:
	docker build --tag $(DOCKER_USER)/trumpet-test-postgres trumpet-test-postgres

.PHONY: launch
launch:
	docker-compose up
