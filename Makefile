shell:
	docker-compose run backend bash
up:
	docker-compose up --build
down:
	docker-compose down
dbconsole:
	docker-compose run backend ./bin/db_console
dbmigrate:
	docker-compose run backend ./bin/db_migrate
dbrollback:
	docker-compose run backend ./bin/db_rollback

