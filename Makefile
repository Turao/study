# https://debezium.io/documentation/reference/2.1/tutorial.html
STORAGE_DIR=${PWD}/storage
STORAGE_SURREAL_DIR=${STORAGE_DIR}/surreal
STORAGE_MYSQL_DIR=${STORAGE_DIR}/mysql
STORAGE_POSTGRES_DIR=${STORAGE_DIR}/postgres
STORAGE_CASSANDRA_DIR=${STORAGE_DIR}/cassandra
STORAGE_CONNECTORS_DIR=${STORAGE_DIR}/debezium

STORAGE_USERS_DIR=${STORAGE_POSTGRES_DIR}/users


# Storage - Users
start-storage-users:
	docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > ${STORAGE_USERS_DIR}/postgresql.conf
	echo 'wal_level = logical' >> ${STORAGE_USERS_DIR}/postgresql.conf
	docker run -it --rm --name storage-users -v ${STORAGE_USERS_DIR}/postgresql.conf:/etc/postgresql/postgresql.conf -p 5432:5432 -e POSTGRES_DB=database -e POSTGRES_USER=pguser -e POSTGRES_PASSWORD=pwd postgres -c 'config_file=/etc/postgresql/postgresql.conf'

migrate-up-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose up 1

migrate-down-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose down 1

migrate-force-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose force ${version}

# Messaging
start-zookeeper:
	docker run -it --rm --name zookeeper -p 2181:2181 -p 2888:2888 -p 3888:3888 quay.io/debezium/zookeeper:2.1

start-kafka:
	docker run -it --rm --name kafka -p 9092:9092 --link zookeeper:zookeeper quay.io/debezium/kafka:2.1

start-kafka-connect:
	docker run -it --rm --name connect -p 8083:8083 -e GROUP_ID=1 -e CONFIG_STORAGE_TOPIC=my_connect_configs -e OFFSET_STORAGE_TOPIC=my_connect_offsets -e STATUS_STORAGE_TOPIC=my_connect_statuses --link kafka:kafka --link storage-users:storage-users quay.io/debezium/connect:2.1

register-storage-connectors:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d '@${STORAGE_CONNECTORS_DIR}/users.json'


# GRPC / Protobuf
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/users/v1.proto