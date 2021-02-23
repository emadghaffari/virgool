# virgool

virgool is a microservice application written by golang and go-kit
each service need own database, this is an example and we use one database for all services, in production you must use local databases for each service

# exports

environment="development"
for development mode we read from file

environment="production"
for production mode we read from file (address and token) then read data from vault

###
change the config.example.yaml to config.yaml

# tools:
 ## database:
 - mysql:5.7
 - phpmyadmin
 - redis
 - rediscommander/redis-commander:latest
 
 ## tracer:
 - jaegertracing/all-in-one:1.20
 
 ## message brokers:
 - zookeeper:3.4.13
 - zookeeper:3.4.13
 - wurstmeister/kafka
 - wurstmeister/kafka
 - wurstmeister/kafka

 ## services monitoring:
 - weaveworks/scope:1.13.1
 
 ## searching with monitoring:
 - docker.elastic.co/elasticsearch/elasticsearch:7.7.1
 - docker.elastic.co/kibana/kibana:7.7.1
 - docker.elastic.co/logstash/logstash:7.7.1
 - docker.elastic.co/beats/filebeat:7.7.1


### auth:
with the auth service you can Register, Login And Verify you users

### notification
with notif service you can send notif to users with SMS or Email
notif service is consumer kafka topic for get notifs from other services

### blog
with blog service you can CRUD posts, tags and upload images 
