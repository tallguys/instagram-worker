# instagram-worker

## Desc
Download images and videos from an instagram account.

## Start
```
$ go run cmd/main.go [username]
```

## Dependency

### PostgreSQL
```
$ docker volume create pgdata
$ docker run -it --rm -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=XPN]un -p 5432:5432 -d postgres

# My OS is Windows with WSL, so it's exposed with the WSL host
$ docker run --rm -p 8080:8080 -d adminer 
```
