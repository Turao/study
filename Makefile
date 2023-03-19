# https://debezium.io/documentation/reference/2.1/tutorial.html
SCHEMAS_DIR=${PWD}/schemas
USERS_SCHEMAS_DIR=${SCHEMAS_DIR}/postgres/users

start-zookeeper:
	docker run -it --rm --name zookeeper -p 2181:2181 -p 2888:2888 -p 3888:3888 quay.io/debezium/zookeeper:2.1

start-kafka:
	docker run -it --rm --name kafka -p 9092:9092 --link zookeeper:zookeeper quay.io/debezium/kafka:2.1

start-database:
	docker run -it --rm --name postgres -p 5432:5432 -e POSTGRES_DB=database -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pwd postgres

database-migrate-up:
	docker run -v ${USERS_SCHEMAS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://user:pwd@localhost:5432/database?sslmode=disable -verbose up 1

database-migrate-down:
	docker run -v ${USERS_SCHEMAS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://user:pwd@localhost:5432/database?sslmode=disable -verbose down 1

start-kafka-connect:
	docker run -it --rm --name connect -p 8083:8083 -e GROUP_ID=1 -e CONFIG_STORAGE_TOPIC=my_connect_configs -e OFFSET_STORAGE_TOPIC=my_connect_offsets -e STATUS_STORAGE_TOPIC=my_connect_statuses --link kafka:kafka --link mysql:mysql quay.io/debezium/connect:2.1

register-connector:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d '{ "name": "inventory-connector", "config": { "connector.class": "io.debezium.connector.mysql.MySqlConnector", "tasks.max": "1", "database.hostname": "mysql", "database.port": "3306", "database.user": "debezium", "database.password": "dbz", "database.server.id": "184054", "topic.prefix": "dbserver1", "database.include.list": "inventory", "schema.history.internal.kafka.bootstrap.servers": "kafka:9092", "schema.history.internal.kafka.topic": "schemahistory.inventory" } }'

start-dependencies:
	start-database
	start-zookeeper
	start-kafka
	start-kafka-connect
	register-connector
