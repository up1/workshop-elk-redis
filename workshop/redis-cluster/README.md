# Redis cluster = Replication + Shading data

## Cluster
Required minimum nodes = 6 nodes (production)
* Master = 3
* Slave = 3

## 1. Start Redis Node with Cluster mode
```
$docker-compose up -d redis1
$docker-compose up -d redis2
$docker-compose up -d redis3
$docker-compose up -d redis4
$docker-compose up -d redis5
$docker-compose up -d redis6

$docker-compose ps

            Name                           Command               State     Ports
----------------------------------------------------------------------------------
redis-cluster_redis1_1          docker-entrypoint.sh redis ...   Up       6379/tcp
redis-cluster_redis2_1          docker-entrypoint.sh redis ...   Up       6379/tcp
redis-cluster_redis3_1          docker-entrypoint.sh redis ...   Up       6379/tcp
redis-cluster_redis4_1          docker-entrypoint.sh redis ...   Up       6379/tcp
redis-cluster_redis5_1          docker-entrypoint.sh redis ...   Up       6379/tcp
redis-cluster_redis6_1          docker-entrypoint.sh redis ...   Up       6379/tcp
```

## 2. Create cluster with replication = 1
* Master 3 nodes
* Slave 3 nodes
```
$docker-compose up redis-cluster

Creating redis-cluster_redis-cluster_1 ... done
Attaching to redis-cluster_redis-cluster_1
redis-cluster_1  | >>> Performing hash slots allocation on 6 nodes...
redis-cluster_1  | Master[0] -> Slots 0 - 5460
redis-cluster_1  | Master[1] -> Slots 5461 - 10922
redis-cluster_1  | Master[2] -> Slots 10923 - 16383
redis-cluster_1  | Adding replica 172.16.238.14:7000 to 172.16.238.10:7000
redis-cluster_1  | Adding replica 172.16.238.15:7000 to 172.16.238.11:7000
redis-cluster_1  | Adding replica 172.16.238.13:7000 to 172.16.238.12:7000
redis-cluster_1  | M: 3705403ec9205209819a9319d835a58bc25b837b 172.16.238.10:7000
redis-cluster_1  |    slots:[0-5460] (5461 slots) master
redis-cluster_1  | M: 1813742a6c2c6a13ba651e1c52ae00877dcca669 172.16.238.11:7000
redis-cluster_1  |    slots:[5461-10922] (5462 slots) master
redis-cluster_1  | M: 9ba2c1f319340e43cbed0646c1ab33008f993dee 172.16.238.12:7000
redis-cluster_1  |    slots:[10923-16383] (5461 slots) master
redis-cluster_1  | S: 143cffd4342dc37ec18eb40f1235c87bece05e14 172.16.238.13:7000
redis-cluster_1  |    replicates 9ba2c1f319340e43cbed0646c1ab33008f993dee
redis-cluster_1  | S: aae82971f6b6c32dcec393ca15e803d225677004 172.16.238.14:7000
redis-cluster_1  |    replicates 3705403ec9205209819a9319d835a58bc25b837b
redis-cluster_1  | S: fb76c0865c1415463fe3f3b6d8c6dae716e65b9f 172.16.238.15:7000
redis-cluster_1  |    replicates 1813742a6c2c6a13ba651e1c52ae00877dcca669
redis-cluster_1  | Can I set the above configuration? (type 'yes' to accept): >>> Nodes configuration updated
redis-cluster_1  | >>> Assign a different config epoch to each node
redis-cluster_1  | >>> Sending CLUSTER MEET messages to join the cluster
redis-cluster_1  | Waiting for the cluster to join
redis-cluster_1  | ..
redis-cluster_1  | >>> Performing Cluster Check (using node 172.16.238.10:7000)
redis-cluster_1  | M: 3705403ec9205209819a9319d835a58bc25b837b 172.16.238.10:7000
redis-cluster_1  |    slots:[0-5460] (5461 slots) master
redis-cluster_1  |    1 additional replica(s)
redis-cluster_1  | S: aae82971f6b6c32dcec393ca15e803d225677004 172.16.238.14:7000
redis-cluster_1  |    slots: (0 slots) slave
redis-cluster_1  |    replicates 3705403ec9205209819a9319d835a58bc25b837b
redis-cluster_1  | M: 1813742a6c2c6a13ba651e1c52ae00877dcca669 172.16.238.11:7000
redis-cluster_1  |    slots:[5461-10922] (5462 slots) master
redis-cluster_1  |    1 additional replica(s)
redis-cluster_1  | S: fb76c0865c1415463fe3f3b6d8c6dae716e65b9f 172.16.238.15:7000
redis-cluster_1  |    slots: (0 slots) slave
redis-cluster_1  |    replicates 1813742a6c2c6a13ba651e1c52ae00877dcca669
redis-cluster_1  | M: 9ba2c1f319340e43cbed0646c1ab33008f993dee 172.16.238.12:7000
redis-cluster_1  |    slots:[10923-16383] (5461 slots) master
redis-cluster_1  |    1 additional replica(s)
redis-cluster_1  | S: 143cffd4342dc37ec18eb40f1235c87bece05e14 172.16.238.13:7000
redis-cluster_1  |    slots: (0 slots) slave
redis-cluster_1  |    replicates 9ba2c1f319340e43cbed0646c1ab33008f993dee
redis-cluster_1  | [OK] All nodes agree about slots configuration.
redis-cluster_1  | >>> Check for open slots...
redis-cluster_1  | >>> Check slots coverage...
redis-cluster_1  | [OK] All 16384 slots covered.

```

Check your cluster on node 1
```
$docker container exec -it redis-cluster_redis1_1 bash
#redis-cli -c -p 7000
>set name hello

-> Redirected to slot [5798] located at 127.0.0.1:7001
OK

>get name
"hello"

```