
Start the Docker containers
```
docker-compose run --service-ports go
docker-compose up -d db
```


#TODO
- Automate 
  Inside the docker container, run
  ```
  go install
  go get github.com/lib/pq
  go run main.go

  docker exec -it db bash
  psql -U test
  ```
