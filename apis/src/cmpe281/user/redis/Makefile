# ref: https://docs.docker.com/docker-cloud/builds/push-images/
# ============= Master ============= #
build-master:
	docker build -t botime/redis-master -f Dockerfile-master0 .
remove-master-image:
	docker rmi botime/redis-master
run-master-bk:
	docker run --name redis-master \
	--network host \
	-p 6379:6379 \
	-p 16379:16379 \
	-d botime/redis-master
	sleep 1
	docker ps -a
	redis-cli ping
run-master:
	docker run --name redis-master \
	--net=bridge \
	-p 6379:6379 \
	-p 16379:16379 \
	-d botime/redis-master
	sleep 1
	docker ps -a
	redis-cli ping
remove-master:
	docker stop redis-master
	docker rm redis-master
restart-master: remove-master build-master run-master-bk

log-master:
	docker logs redis-master
connect-master:
	docker exec -it redis-master bash

# ============== Slave ============== #
build-slave:
	#docker rmi botime/redis-slave
	docker build -t botime/redis-slave -f Dockerfile-slave .
run-slave-bk:
	docker run --name redis-slave \
	--net=host \
	-d botime/redis-slave
	sleep 1
	docker ps -a
run-slave-0:
	docker run --name redis-slave \
	--net=bridge \
	-p 6390:6379 \
	-p 16390:16379 \
	--privileged \
	-d botime/redis-slave
	sleep 1
	docker ps -a
run-slave-1:
	docker run --name redis-slave-1 \
	--net=bridge \
	-p 6391:6379 \
	-p 16391:16379 \
	-d botime/redis-slave
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
push-master:
	docker tag botime/redis-master botime/redis-master
	docker push botime/redis-master
push-slave:
	docker tag botime/redis-slave botime/redis-slave
	docker push botime/redis-slave
pull-master:
	docker pull botime/redis-master
pull-slave:
	docker pull botime/redis-slave
block-master:
	sudo iptables -A INPUT -s 10.0.1.87 -j DROP
unblock-master:
	sudo iptables -D INPUT -s 10.0.1.87 -j DROP
block-slaves:
	sudo iptables -A INPUT -s 10.0.1.186 -j DROP;
	sudo iptables -A INPUT -s 10.0.1.125 -j DROP;
unblock-slaves:
	sudo iptables -D INPUT -s 10.0.1.186 -j DROP;
	sudo iptables -D INPUT -s 10.0.1.125 -j DROP;
block-ports:
	sudo iptables -A INPUT -p tcp --destination-port 6379 -j DROP;
	sudo iptables -A INPUT -p tcp --destination-port 16379 -j DROP;
unblock-ports:
	sudo iptables -D INPUT -p tcp --destination-port 6379 -j DROP;
	sudo iptables -D INPUT -p tcp --destination-port 16379 -j DROP;
