# ============= Master ============= #
build-master:
	docker build -t redis:master -f Dockerfile-master .
run-master:
	docker run --name redis-master \
	--net=host \
	-p 6379:6379 \
	-p 26379:26379 \
	-d redis:master
	sleep 1
	docker ps -a
remove-master:
	docker stop redis-master
	docker rm redis-master
restart-master: remove-master build-master run-master

log-master:
	docker logs redis-master
connect-master:
	docker exec -it redis-master bash

# ============== Slave ============== #
build-slave:
	docker build -t redis:slave -f Dockerfile-slave .
run-slave:
	docker run --name redis-slave \
	--net=host \
	-p 6380:6380 \
	-p 26380:26380 \
	-d redis:slave
	sleep 1
	docker ps -a
remove-slave:
	docker stop redis-slave
	docker rm redis-slave
restart-slave: remove-slave build-slave run-slave

log-slave:
	docker logs redis-slave
connect-slave:
	docker exec -it redis-slave bash

# ======== Miscs ========== #
stop-the-world:
	docker stop redis-master
	docker stop redis-slave
clean-failed:
	docker rmi $(docker images | grep "^<none>" | awk '{print $3}');
redis-check-conncted-slaves:

	redis-cli info | grep connected_slaves
