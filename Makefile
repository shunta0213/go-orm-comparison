up:
	docker compose up --build -d


down:
	docker compose down --rmi local --volumes --remove-orphans