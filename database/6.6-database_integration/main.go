package main

import "6_6/example/postgres"

func main() {
	//redis.RedisString()
	//redis.RedisScan()
	postgres.PostgresSimpleQuery()
	postgres.PostgresJoinQuery()

}

//┌─────(free)─────(~/GOLANG/Slurm_course/golang-main/6.6)
//└> $ go run main.go
//QueryRow failed: ERROR: relation "users" does not exist (SQLSTATE 42P01)
//exit status 1
