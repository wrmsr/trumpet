.PHONY: docker
docker:
	docker build --tag willremind101/trumpet-test-postgres .

.PHONY: launch
launch:
	docker-compose up
