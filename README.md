Queue and Runner for docker actions
-----------------------------------

Use a redis queue to execute command via docker in a pub/sub client/server relationship.


Building
--------

Start a redis server

    docker run -d -p 6379:6379 -t redis

Start the queue

    go run queue.go

Start a runner

    go run runner.go


