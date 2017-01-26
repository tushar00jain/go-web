Start the Docker containers
```
docker-compose run --service-ports go
docker-compose up -d db
```
# TODO

  - go
  ```
  cd server
  bash gateway/stub-proxy.sh
  cd ..

  go run main.go
  bash ${GOPATH}/bin/service
  ```

# DEBUG

  - db
  ```
  docker exec -it db bash
  psql -U test
  ```
