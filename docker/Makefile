REPO=darrylwest
IMAGE=$(REPO)/spotcached:latest
PORT=19501
SPOT_HOME=$(HOME)/.spotcache

docker:
	docker build -f Dockerfile.scratch -t $(IMAGE) .

publish:
	docker push $(IMAGE)

run:
	docker run -d -p $(PORT):19501 -v $(SPOT_HOME):/.spotcache --name spotcached-$(PORT) $(IMAGE)

