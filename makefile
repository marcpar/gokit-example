.PHONY: help

help: ## This help.
	


myname:
	@echo "Hello!! I'm an super awesome makefile"

build:
	@echo "Hey! cool your building!!!"
	docker-compose build 
	@echo "you definitley build it congratulations!!"

dev:
	@echo "DEVELOPMENT MODE!!"
	docker-compose up -d
	@echo "Awesome! we can start developing"

test:
	@echo "Woah!! your testing, Cool!"
	
clean:
	@echo "Hey!! we are cleaning!!"
	

# TODO: pushing image to registry
push:
	@echo "pushing images to docker registry"
	# docker push 
