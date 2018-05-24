# Redis geo.lua example using Go

## To run

- `git clone https://github.com/abhirockzz/redis-geo.lua-golang.git`
- `cd redis-geo.lua-golang`
- Get `geo.lua` script - `curl -O https://raw.githubusercontent.com/RedisLabs/geo.lua/master/geo.lua`
- `eval $(docker-machine env)`
- Start Redis on Docker - `docker run --name redis --rm -p 6379:6379 redis`
- Run the example - `go run redis-geo-lua-example.go -redis-server=$(docker-machine ip):6379`