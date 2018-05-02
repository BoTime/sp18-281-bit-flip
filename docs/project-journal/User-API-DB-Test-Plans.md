**Database: Redis**

### Test 1 - Single Node (Local & EC2)
|Step#|Description|Status|Date|
| --- | --- | --- | --- |
|1| Create a single node Redis Server locally. Test reads and writes. | [x] | April 15 |
|2| Create a single node Redis Server on EC2. Test reads and writes. | [x] | May 1 |

---

### Test 2 - Two Nodes (Local)
|Step#|Description|Status|Date|
| --- | --- | --- | --- |
|1| Create a two nodes Redis Cluster locally. One Master and one Slave. | [] |  |
|2|  Write to Master node and check if Slave will receive writes. | []| ... |

---

### Test 3 - Three Nodes (EC2)
|Step#|Description|Status|Date|
| --- | --- | --- | --- |
|1| Create a two nodes Redis Cluster on EC2. One Master and two Slaves. | [] |  |
|2| Write `key1` to M1 and make sure R2, R3 receives `key1`. | [] | ... |
|3| Create network partitions. M1 in partition 1, R1 and R2 in partition 2. | [] | ... |
|4| Wait for sentinels to detect failure. Wait for failover. How long? | [] | ... |
|5| Write `key2` to M1 again and confirm.| [] | ... |
|6| Remove network partition. And wait for the three nodes to recover. | [] | ... |
|7| Read `key2` from M1. Now `key2` should be `nil`. | [] | ... |

```
Before partition:

    +----+
    | M1 |
    | S1 |
    +----+
       |
    +----+    |    +----+
    | R2 |----+----| R3 |
    | S2 |         | S3 |
    +----+         +----+

Configuration: quorum = 2

After partition:

    +----+
    | M1 |
    | S1 | <- C1 (writes will be lost)
    +----+
       |
     @@@@@
     @@@@@
       |
    +------+    |    +----+
    | [M2] |----+----| R3 |
    | S2   |         | S3 |
    +------+         +----+
```
--
### Test 4 - Five Nodes (EC2)
|Step#|Description|Status|Date|
| --- | --- | --- | --- |
|1| Create a two nodes Redis Cluster on EC2. One Master and four Slave nodes. | [] |  |
|2| Write `key1` to M1 and make sure R2, R3, R4, R5 receives `key1`. | [] | ... |
|3| Create network partitions. M1, R2 in partition 1, R3, R4 and R5 in partition 2. | [] | ... |
|4| Wait for sentinels to detect failure. Wait for failover. How long? | [] | ... |
|5| Write `key2` to M1 again and confirm.| [] | ... |
|6| Remove network partition. And wait for the three nodes to recover. | [] | ... |
|7| Read `key2` from M1. Now `key2` should be `nil`. | [] | ... |

```
Before partition:

+----+         +----+
| M1 |----+----| R2 |
| S1 |    |    | S2 |
+----+    |    +----+
          |
+------------+------------+
|            |            |
|            |            |
+----+        +----+      +----+
| R3 |        | R4 |      | R5 |
| S3 |        | S4 |      | S5 |
+----+        +----+      +----+

Configuration: quorum = 2

After partition:

+----+         +----+
| R2 |----+----| M1 |  <- C1 (writes will be lost)
| S2 |    |    | S1 |
+----+    |    +----+
        @@@@@
        @@@@@
          |
+------------+------------+
|            |            |
|            |            |
+----+        +----+      +----+
|[M2]|        | R4 |      | R5 |
| S3 |        | S4 |      | S5 |
+----+        +----+      +----+

```
