.PHONY: version foo

version:
	docker-compose up -d develop
	docker-compose exec develop java --version

foo:
	docker-compose up -d develop
	docker-compose exec develop javac Foo.java && docker-compose exec develop java Foo
