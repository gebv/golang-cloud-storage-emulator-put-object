backend1:
	@echo firebase emulator storage
	docker-compose up -d firebase-emulator
	sleep 10
	STORAGE_EMULATOR_SKIP_CREATE_BUCKET=1 STORAGE_EMULATOR_HOST=localhost:9199 go run main.go

backend2:
	@echo fsouza/fake-gcs-server backend
	docker-compose up -d fake-gcs-server
	sleep 3
	STORAGE_EMULATOR_HOST=localhost:8081 go run main.go
