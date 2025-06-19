# https://debezium.io/documentation/reference/2.1/tutorial.html
STORAGE_DIR=${PWD}/storage
STORAGE_SURREAL_DIR=${STORAGE_DIR}/surreal
STORAGE_MYSQL_DIR=${STORAGE_DIR}/mysql
STORAGE_POSTGRES_DIR=${STORAGE_DIR}/postgres
STORAGE_CASSANDRA_DIR=${STORAGE_DIR}/cassandra
STORAGE_CONNECTORS_DIR=${STORAGE_DIR}/debezium

STORAGE_USERS_DIR=${STORAGE_POSTGRES_DIR}/users

# Infrastructure
start-network:
	docker network create global

# Storage - Users
start-storage-users:
	docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > ${STORAGE_USERS_DIR}/postgresql.conf
	echo 'wal_level = logical' >> ${STORAGE_USERS_DIR}/postgresql.conf
	docker run -it --rm --name storage-users \
		--network global \
		-v ${STORAGE_USERS_DIR}/postgresql.conf:/etc/postgresql/postgresql.conf \
		-p 5432:5432 \
		-e POSTGRES_DB=database \
		-e POSTGRES_USER=pguser \
		-e POSTGRES_PASSWORD=pwd \
		postgres \
		-c 'config_file=/etc/postgresql/postgresql.conf'

migrate-up-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose up 1

migrate-down-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose down 1

migrate-force-storage-users:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose force ${version}

# Messaging
start-zookeeper:
	docker run -it --rm --name zookeeper \
		--network global \
		-p 2181:2181 \
		-p 2888:2888 \
		-p 3888:3888 \
		quay.io/debezium/zookeeper:2.1

start-kafka:
	docker run -it --rm --name kafka \
		--network global \
		-p 9092:9092 \
		-e BROKER_ID=1 \
		-e ZOOKEEPER_CONNECT=zookeeper \
		quay.io/debezium/kafka:2.1

start-kafka-connect:
	docker run -it --rm --name connect \
		--network global \
		-p 8083:8083 \
		-e BOOTSTRAP_SERVERS=kafka:9092 \
		-e GROUP_ID=1 \
		-e CONFIG_STORAGE_TOPIC=my_connect_configs \
		-e OFFSET_STORAGE_TOPIC=my_connect_offsets \
		-e STATUS_STORAGE_TOPIC=my_connect_statuses \
		quay.io/debezium/connect:2.1

start-kafka-ui:
	docker run -it --rm --name kafka-ui \
	--network global \
	-p 8080:8080 \
	-e DYNAMIC_CONFIG_ENABLED=true \
	-e KAFKA_CLUSTERS_0_NAME=cluster \
    -e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka:9092 \
	provectuslabs/kafka-ui

register-storage-connectors:
	curl -i -X POST \
		-H "Accept:application/json" \
		-H "Content-Type:application/json" \
		localhost:8083/connectors/ \
		-d '@${STORAGE_CONNECTORS_DIR}/users.json'

# GRPC / Protobuf
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/users/v1.proto