#！/bin/bash


mkdir -p  ../docker/mariadb/data

docker-compose -f ../docker/mariadb.yaml up -d
