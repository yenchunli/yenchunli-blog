---
title: "Write a request limiter in Go"
date: 2020-04-12T14:08:30+08:00
draft: false

tags: ["Go", "Redis", "Gin"]

summary: "A middleware which can restrict ip request in certain time window"

---

# Go-Request-limit

[Github](https://github.com/412988937/go-request-limit)

## How to test
```bash
docker run --name redis-lab -p 6379:6379 -d redis
cd <project-name>
go build main.go //if necessary
./main
```

## Design Idea

There are the scenerio on how to prevent URL from users' request.

1. Request come
2. Check if user's ip is in the Redis list
3. If not, create a redis list, whose key is user's ip. Return.
4. If yes, check the redis list with key=(user's ip) and calcuate its length.
5. If length is over request limit, return status `too many reqest`.
6. If length is valid, push user's ip to the list and return.

* Return data including `X-RateLimit-Remaining` and `X-RateLimit-Reset`. The value of `X-RateLimit-Reset` is based on the TTL of certain redis list. The TTL will start when the first request comes and will expire in an hour. After an hour, the redis list of certain key will disappear.


