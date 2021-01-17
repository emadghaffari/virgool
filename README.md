# virgool

virgool is a microservice application written by golang and go-kit
each service need own database, this is an example and we use one database for all services, in production you must use local databases for each service

# exports

environment="development"
for development mode we read from file

environment="production"
for production mode we read from file (address and token) then read data from vault

# kafka

kafka service

## create topic:

kafka-topics.sh --zookeeper zookeeper:2181 --topic notifications --create --partitions 3 --replication-factor 1

## list topics:

kafka-topics.sh --zookeeper zookeeper:2181 --list
