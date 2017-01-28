Start the Docker containers
```
docker-compose up -d db
docker-compose run --service-ports go_service
docker-compose run --service-ports go_main
```
# TODO

  - go

  ```
  bash stub-proxy.sh
  go run gateway/proxy.go

  go run main.go
  ```

# DEBUG

  - db
  ```
  docker exec -it db bash
  psql -U test
  ```
