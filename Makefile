shell:
	docker-compose run backend bash
up:
	docker-compose up --build
down:
	docker-compose down
console:
	docker-compose run backend ./dbconsole
