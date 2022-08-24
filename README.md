# pubsub
simple publisher / subscriber application with different methods 


For start redis publisher server :
```
$ make serverRedis
```

For start redis subscriber server :
```
$ make serverSubscriberRedis
```
<br/>

For start kafka publisher server :
```
$ make serverKafka
```

For start kafka subscriber server :
```
$ make serverSubscriberKafka
```

For start rabbitmq publisher server :
```
$ make serverRabbitmq
```

For start rabbitmq subscriber server :
```
$ make serverSubscriberRabbitmq
```

For start go publisher and subscriber server :
```
$ make serverGo
```

benchmark GO:
```
 ab -k -p json.txt -T aplication/json -n 100000 -c 100 -t 1000  http://127.0.0.1:3000/pubsub-golang}}/
```

benchmark Redis:
```
 ab -k -p json.txt -T aplication/json -n 100000 -c 100 -t 1000  http://127.0.0.1:3000/pubsub-redis}}/
```

benchmark rabbitmq:
```
 ab -k -p json.txt -T aplication/json -n 100000 -c 100 -t 1000  http://127.0.0.1:3000/pubsub-rabbitmq}}/
```


benchmark kafka:
```
 ab -k -p json.txt -T aplication/json -n 100000 -c 100 -t 1000  http://127.0.0.1:3000/pubsub-kafka}}/
```

