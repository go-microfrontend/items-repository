up-db:
	sudo docker-compose up -d db

up-app:
	sudo docker-compose up -d app

up: up-db up-app

down:
	sudo docker-compose down

restart: down up

db-shell:
	sudo docker-compose exec db psql -U postgres -d postgres

app-shell:
	sudo docker-compose exec app sh

