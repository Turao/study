# https://debezium.io/documentation/reference/2.1/tutorial.html
STORAGE_DIR=${PWD}/storage
STORAGE_USERS_DIR=${STORAGE_DIR}/postgres/users

start-zookeeper:
	docker run -it --rm --name zookeeper -p 2181:2181 -p 2888:2888 -p 3888:3888 quay.io/debezium/zookeeper:2.1

start-kafka:
	docker run -it --rm --name kafka -p 9092:9092 --link zookeeper:zookeeper quay.io/debezium/kafka:2.1

start-database:
	docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > ${STORAGE_USERS_DIR}/postgresql.conf
	echo 'wal_level = logical' >> ${STORAGE_USERS_DIR}/postgresql.conf
	docker run -it --rm --name database -v ${STORAGE_USERS_DIR}/postgresql.conf:/etc/postgresql/postgresql.conf -p 5432:5432 -e POSTGRES_DB=database -e POSTGRES_USER=pguser -e POSTGRES_PASSWORD=pwd postgres -c 'config_file=/etc/postgresql/postgresql.conf'

database-migrate-up:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose up 1

database-migrate-down:
	docker run -v ${STORAGE_USERS_DIR}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://pguser:pwd@localhost:5432/database?sslmode=disable -verbose down 1

start-kafka-connect:
	docker run -it --rm --name connect -p 8083:8083 -e GROUP_ID=1 -e CONFIG_STORAGE_TOPIC=my_connect_configs -e OFFSET_STORAGE_TOPIC=my_connect_offsets -e STATUS_STORAGE_TOPIC=my_connect_statuses --link kafka:kafka --link database:database quay.io/debezium/connect:2.1

register-connector:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d '{ "name": "user-connector", "config": { "connector.class": "io.debezium.connector.postgresql.PostgresConnector", "tasks.max": "1", "database.hostname": "database", "database.port": "5432", "database.user": "pguser", "database.password": "pwd", "database.dbname":"database", "database.server.id": "184054", "topic.prefix": "dbserver1", "database.include.list": "user", "schema.history.internal.kafka.bootstrap.servers": "kafka:9092", "schema.history.internal.kafka.topic": "schemahistory.user" } }'

start-dependencies:
	start-database
	start-zookeeper
	start-kafka
	start-kafka-connect
	register-connector
