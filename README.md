Start the Docker containers
```
docker-compose run --service-ports go_main
docker-compose up -d db
```
# TODO

  - go

  ```
  cd server
  bash gateway/stub-proxy.sh
  cd ..
  go run main.go

  docker exec -it go_main bash
  ${GOPATH}/bin/service
  ```

# DEBUG

  - db
  ```
  docker exec -it db bash
  psql -U test
  ```
