#!/bin/sh
# ref:
# https://github.com/mustafaileri/redis-cluster-with-sentinel/blob/master/sentinel/entrypoint.sh

sed -i "s/\$SENTINEL_PORT/$SENTINEL_PORT/g" /usr/local/etc/redis/sentinel.conf
sed -i "s/\$SENTINEL_QUORUM/$SENTINEL_QUORUM/g" /usr/local/etc/redis/sentinel.conf
sed -i "s/\$SENTINEL_DOWN_AFTER/$SENTINEL_DOWN_AFTER/g" /usr/local/etc/redis/sentinel.conf
sed -i "s/\$SENTINEL_FAILOVER/$SENTINEL_FAILOVER/g" /usr/local/etc/redis/sentinel.conf

exec docker-entrypoint.sh redis-server /usr/local/etc/redis/redis.conf
exec docker-entrypoint.sh redis-sentinel /usr/local/etc/redis/sentinel.conf
