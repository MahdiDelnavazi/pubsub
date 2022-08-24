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
<br/>

For start rabbitmq publisher server :
```
$ make serverRabbitmq
```

For start rabbitmq subscriber server :
```
$ make serverSubscriberRabbitmq
```
<br/>

For start go publisher and subscriber server :
```
$ make serverGo
```

<br/><br/>
benchmark golang, redis, kafka or rabbitmq:
```
 ab -k -p json.txt -T aplication/json -n 100000 -c 100 -t 1000  http://127.0.0.1:3000/pubsub-{YOUR-METHOD}}}/
```

<br/>

**results**

|      name    | concurrency level |   longest     |   90% time   |   99% time   |   request count   |   time   |   memory   |   cpu%   |
|    :---:     |     :---:         |    :---:      |:---:      |:---:      |:---:      |:---:      |:---:      |:---:      | 
| golang       |32                 | 71ms          |3ms       |8ms       |50k      |2.551s      |200mb      |40%      |-      |-      |
| kafka        |32                 | 241ms         |4ms       |12ms      |50k      |4.495s      |100mb      |90%      |-      |-      |
| redis        |32                 | 211ms         |6ms       |19ms      |50k      |6.524s      |200mb      |90%      |-      |-      |
| rabbitmq     |32                 | 232ms         |3ms       |9ms       |50k      |4.120s      |220mb      |90%      |-      |-      |
| golang       |100                | 64ms          |10ms      |22ms      |50k      | 2.603s     |100mb      |20%      |-      |-      |
| kafka        |100                | 236ms         |10ms      |36ms      |50k      |3.231s      |400mb      |40%      |-      |-      |
| redis        |100                | 282ms         |17ms      |32ms      |50k      | 4.762s     |300mb      |30%      |-      |-      |
| rabbitmq     |100                | 134ms         |9ms       |37ms      |50k      |3.328s      |400mb      |40%      |-      |-      |
 




