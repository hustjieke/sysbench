Table of Contents
=================

   * [RadonDB/Mysql sysbench 测试数据记录](#radondbmysql-sysbench-测试数据记录)
   * [overview(整体结果汇总)](#overview整体结果汇总)
   * [RadonDB(版本5.7-Radon-1.0)](#radondb版本57-radon-10)
      * [1 net_optane_neonsan](#1-net_optane_neonsan)
         * [1.1 RadonDB 3节点写2亿行](#11-radondb-3节点写2亿行)
         * [1.2 RadonDB 3节点只读, 读取刚写入的2亿行数据，2 hours, 读取约两亿3千万行](#12-radondb-3节点只读-读取刚写入的2亿行数据2-hours-读取约两亿3千万行)
         * [1.3 RadonDB 3节点混合读写（比例7:2), 查询次数：两亿](#13-radondb-3节点混合读写比例72-查询次数两亿)
         * [1.4 RadonDB 2节点写2亿行](#14-radondb-2节点写2亿行)
         * [1.5 RadonDB 2节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据](#15-radondb-2节点只读-读取刚写入的2亿行数据2-hours-读取约1亿4千万条数据)
         * [1.6 RadonDB 2节点混合读写（比例7:2), 查询次数:两亿行](#16-radondb-2节点混合读写比例72-查询次数两亿行)
         * [1.7 RadonDB 1节点写2亿行](#17-radondb-1节点写2亿行)
         * [1.8 RadonDB 1节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据](#18-radondb-1节点只读-读取刚写入的2亿行数据2-hours-读取约1亿4千万条数据)
         * [1.9 RadonDB 1节点混合读写(比例7:2), 查询次数:两亿行](#19-radondb-1节点混合读写比例72-查询次数两亿行)
      * [2 local optane](#2-local-optane)
         * [2.1 RadonDB 3节点写2亿行](#21-radondb-3节点写2亿行)
         * [2.2 RadonDB 3节点只读, 读取刚写入的2亿行数据，2 hours](#22-radondb-3节点只读-读取刚写入的2亿行数据2-hours)
         * [2.3 RadonDB 3节点混合读写（比例7:2), 查询次数：两亿](#23-radondb-3节点混合读写比例72-查询次数两亿)
         * [2.4 RadonDB 2节点写2亿行](#24-radondb-2节点写2亿行)
         * [2.5 RadonDB 2节点只读, 读取刚写入的2亿行数据，2 hours](#25-radondb-2节点只读-读取刚写入的2亿行数据2-hours)
         * [2.6 RadonDB 2节点混合读写（比例7:2), 查询次数:两亿行](#26-radondb-2节点混合读写比例72-查询次数两亿行)
         * [2.7 RadonDB 1节点写2亿行](#27-radondb-1节点写2亿行)
         * [2.8 RadonDB 1节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据](#28-radondb-1节点只读-读取刚写入的2亿行数据2-hours-读取约1亿4千万条数据)
         * [2.9 RadonDB 1节点混合读写(比例7:2), 查询次数:两亿行](#29-radondb-1节点混合读写比例72-查询次数两亿行)
   * [Mysql(版本5.7.20)](#mysql版本5720)
      * [1 net_optane_neonsan](#1-net_optane_neonsan-1)
         * [1.1 写2亿行](#11-写2亿行)
         * [1.2 读取刚写入的2亿行数据，2 hours](#12-读取刚写入的2亿行数据2-hours)
         * [1.3 混合读写(比例7:2), queries:两亿](#13-混合读写比例72-queries两亿)
      * [local optane](#local-optane)
         * [2.1 写2亿行](#21-写2亿行)
         * [2.2 读取刚写入的2亿行数据，2 hours](#22-读取刚写入的2亿行数据2-hours)
         * [2.3 混合读写(比例7:2), queries:两亿](#23-混合读写比例72-queries两亿)
----------------------------------------------------------

# RadonDB/Mysql sysbench 测试数据记录

# overview(整体结果汇总)

（物理环境、RadonDB、Mysql、脚本配置等相关参数待补充）

`optane盘 文件io测试记录`

`prepare`:

```
sysbench 1.0.11 (using system LuaJIT 2.1.0-beta3)

3 files, 1048576Kb each, 3072Mb total
Creating files for the test...
Extra file open flags: 0
Creating file test_file.0
Creating file test_file.1
Creating file test_file.2
3221225472 bytes written in 3.65 seconds (840.51 MiB/sec).
```

`run`:

```
sysbench 1.0.11 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 32
Initializing random number generator from current time


Extra file open flags: 0
3 files, 1GiB each
3GiB total file size
Block size 16KiB
Number of IO requests: 0
Read/Write ratio for combined random IO test: 0.25
Periodic FSYNC enabled, calling fsync() each 100 requests.
Calling fsync() at the end of test, Enabled.
Using synchronous I/O mode
Doing random r/w test
Initializing worker threads...

Threads started!


File operations:
    reads/s:                      33296.86
    writes/s:                     133188.16
    fsyncs/s:                     4994.48

Throughput:
    read, MiB/s:                  520.26
    written, MiB/s:               2081.06

General statistics:
    total time:                          10.0013s
    total number of events:              1715384

Latency (ms):
         min:                                  0.00
         avg:                                  0.18
         max:                                 22.58
         95th percentile:                      0.12
         sum:                             316965.70

Threads fairness:
    events (avg/stddev):           53605.7500/1191.36
    execution time (avg/stddev):   9.9052/0.01
```

`线程数`：512
`表数量`：16
`测试数据`: 2亿行约81G

`net_optane_neonsan`: 为数据通过网络写入NeonSAN集群,并且直接NeonSAN集群将optane盘直接作为存储盘

`local_optane`: 为数据直接写入本地的optane盘,不走网络写入NeonSAN, 以下数据记录`qps`

| 写2亿条数据 | RadonDB3节点 | RadonDB2节点 | RadonDB1节点 | 单机mysql
| ------ | ------ | ------ |----|----|
| net_optane_neonsan |6820.36| 3995.52|2495.48|5136.86|
| local_optane | 56075.02 |24730.52 |11889.50 |22547.43

| 读2 hours | RadonDB3节点 | RadonDB2节点 | RadonDB1节点 | 单机mysql
| ------ | ------ | ------ |----|----|
| net_optane_neonsan |33078.69| 20047.72 |6979.96 |49955.17
| local_optane | 27492.23| 20737.91 |7157.75| 51740.06

| 混合读写(7:2)2亿条数据 | RadonDB3节点 | RadonDB2节点 | RadonDB1节点 | 单机mysql
| ------ | ------ | ------ |----|----|
| net_optane_neonsan |40320.60 |22970.97 | 8473.46 |38799.02
| local_optane | 7157.75 |25313.72 |(测试中) |51886.00

# RadonDB(版本5.7-Radon-1.0)

## 1 net_optane_neonsan

### 1.1 RadonDB 3节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0    
        write:                           200000000
        other:                           0    
        total:                           200000000
    transactions:                        200000000 (6820.36 per sec.)
    queries:                             200000000 (6820.36 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          29323.9469s
    total number of events:              200000000

Latency (ms):
         min:                                  0.63 
         avg:                                 75.07
         max:                              30778.41
         95th percentile:                    458.96
         sum:                            15013311492.00

Threads fairness:
    events (avg/stddev):           390625.0000/1284.88
    execution time (avg/stddev):   29322.8740/0.05


```

### 1.2 RadonDB 3节点只读, 读取刚写入的2亿行数据，2 hours, 读取约两亿3千万行

```
SQL statistics:
    queries performed:
        read:                            238171402
        write:                           0    
        other:                           0    
        total:                           238171402
    transactions:                        17012243 (2362.76 per sec.)
    queries:                             238171402 (33078.69 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.1425s
    total number of events:              17012243

Latency (ms):
         min:                                 53.50
         avg:                                216.69
         max:                               1145.61
         95th percentile:                    287.38
         sum:                            3686403703.37

Threads fairness:
    events (avg/stddev):           33227.0371/33.69
    execution time (avg/stddev):   7200.0072/0.04
```

### 1.3 RadonDB 3节点混合读写（比例7:2), 查询次数：两亿

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           0    
        total:                           200000016
    transactions:                        11111112 (2240.03 per sec.)
    queries:                             200000016 (40320.60 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          4960.2418s
    total number of events:              11111112

Latency (ms):
         min:                                 18.07
         avg:                                228.56
         max:                               1907.91
         95th percentile:                    369.77
         sum:                            2539583222.78

Threads fairness:
    events (avg/stddev):           21701.3906/49.94
    execution time (avg/stddev):   4960.1235/0.04
```

### 1.4 RadonDB 2节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0
        write:                           200000000
        other:                           0
        total:                           200000000
    transactions:                        200000000 (3995.52 per sec.)
    queries:                             200000000 (3995.52 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          50056.0945s
    total number of events:              200000000

Latency (ms):
         min:                                  0.63
         avg:                                128.14
         max:                              31052.41
         95th percentile:                    569.67
         sum:                            25628227171.64

Threads fairness:
    events (avg/stddev):           390625.0000/864.06
    execution time (avg/stddev):   50055.1312/0.02
```

### 1.5 RadonDB 2节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据

```
SQL statistics:
    queries performed:
        read:                            144347770
        write:                           0
        other:                           0
        total:                           144347770
    transactions:                        10310555 (1431.98 per sec.)
    queries:                             144347770 (20047.72 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.2053s
    total number of events:              10310555

Latency (ms):
         min:                                 78.54
         avg:                                357.54
         max:                                692.24
         95th percentile:                    427.07
         sum:                            3686438527.84

Threads fairness:
    events (avg/stddev):           20137.8027/17.28
    execution time (avg/stddev):   7200.0752/0.06
```

### 1.6 RadonDB 2节点混合读写（比例7:2), 查询次数:两亿行

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           0    
        total:                           200000016
    transactions:                        11111112 (1276.16 per sec.)
    queries:                             200000016 (22970.97 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          8706.6406s
    total number of events:              11111112

Latency (ms):
         min:                                 30.94
         avg:                                401.20
         max:                               1939.10
         95th percentile:                    502.20
         sum:                            4457724349.38

Threads fairness:
    events (avg/stddev):           21701.3906/20.35
    execution time (avg/stddev):   8706.4929/0.07
```

### 1.7 RadonDB 1节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0     
        write:                           200000000
        other:                           0     
        total:                           200000000
    transactions:                        200000000 (2495.48 per sec.) 
    queries:                             200000000 (2495.48 per sec.) 
    ignored errors:                      0      (0.00 per sec.) 
    reconnects:                          0      (0.00 per sec.) 

General statistics:
    total time:                          80144.7502s
    total number of events:              200000000

Latency (ms): 
         min:                                  0.67  
         avg:                                205.17
         max:                              75608.44
         95th percentile:                     44.98 
         sum:                            41033658803.98

Threads fairness:
    events (avg/stddev):           390625.0000/367.71
    execution time (avg/stddev):   80143.8649/0.04
```

### 1.8 RadonDB 1节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据

```
SQL statistics:
    queries performed:
        read:                            50259538
        write:                           0
        other:                           0
        total:                           50259538
    transactions:                        3589967 (498.57 per sec.)
    queries:                             50259538 (6979.96 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.5440s
    total number of events:              3589967

Latency (ms):
         min:                                495.99
         avg:                               1026.91
         max:                               1867.63
         95th percentile:                   1129.24
         sum:                            3686566955.22

Threads fairness:
    events (avg/stddev):           7011.6543/4.67
    execution time (avg/stddev):   7200.3261/0.16
```

### 1.9 RadonDB 1节点混合读写(比例7:2), 查询次数:两亿行

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           0
        total:                           200000016
    transactions:                        11111112 (470.75 per sec.)
    queries:                             200000016 (8473.46 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          23603.1096s
    total number of events:              11111112

Latency (ms):
         min:                                392.61
         avg:                               1087.62
         max:                               2172.95
         95th percentile:                   1213.57
         sum:                            12084657746.82

Threads fairness:
    events (avg/stddev):           21701.3906/9.19
    execution time (avg/stddev):   23602.8472/0.16
```

## 2 local optane

### 2.1 RadonDB 3节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0
        write:                           200000000
        other:                           0
        total:                           200000000
    transactions:                        200000000 (56075.02 per sec.)
    queries:                             200000000 (56075.02 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          3566.6487s
    total number of events:              200000000

Latency (ms):
         min:                                  0.33
         avg:                                  9.13
         max:                                638.45
         95th percentile:                     51.02
         sum:                            1825758962.32

Threads fairness:
    events (avg/stddev):           390625.0000/1600.96
    execution time (avg/stddev):   3565.9355/0.12
```

### 2.2 RadonDB 3节点只读, 读取刚写入的2亿行数据，2 hours

```
SQL statistics:
    queries performed:
        read:                            197948660
        write:                           0
        other:                           0
        total:                           197948660
    transactions:                        14139190 (1963.73 per sec.)
    queries:                             197948660 (27492.23 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.1639s
    total number of events:              14139190

Latency (ms):
         min:                                 91.69
         avg:                                260.72
         max:                                586.02
         95th percentile:                    331.91
         sum:                            3686417273.31

Threads fairness:
    events (avg/stddev):           27615.6055/24.97
    execution time (avg/stddev):   7200.0337/0.04
```

### 2.3 RadonDB 3节点混合读写（比例7:2), 查询次数：两亿

```
SQL statistics:
    queries performed:
        read:                            51539418
        write:                           0
        other:                           0
        total:                           51539418
    transactions:                        3681387 (511.27 per sec.)
    queries:                             51539418 (7157.75 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.5012s
    total number of events:              3681387

Latency (ms):
         min:                                470.68
         avg:                               1001.40
         max:                               1609.69
         95th percentile:                   1089.30
         sum:                            3686552069.68

Threads fairness:
    events (avg/stddev):           7190.2090/4.63
    execution time (avg/stddev):   7200.2970/0.15
```

### 2.4 RadonDB 2节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0
        write:                           200000000
        other:                           0
        total:                           200000000
    transactions:                        200000000 (24730.52 per sec.)
    queries:                             200000000 (24730.52 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          8087.1704s
    total number of events:              200000000

Latency (ms):
         min:                                  0.34
         avg:                                 20.70
         max:                               2976.54
         95th percentile:                    121.08
         sum:                            4140023661.40

Threads fairness:
    events (avg/stddev):           390625.0000/1177.76
    execution time (avg/stddev):   8085.9837/0.20
```

### 2.5 RadonDB 2节点只读, 读取刚写入的2亿行数据，2 hours

```
SQL statistics:
    queries performed:
        read:                            149317070
        write:                           0
        other:                           0
        total:                           149317070
    transactions:                        10665505 (1481.28 per sec.)
    queries:                             149317070 (20737.91 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.1940s
    total number of events:              10665505

Latency (ms):
         min:                                143.18
         avg:                                345.64
         max:                                672.06
         95th percentile:                    411.96
         sum:                            3686430284.64

Threads fairness:
    events (avg/stddev):           20831.0645/16.62
    execution time (avg/stddev):   7200.0591/0.05
```

### 2.6 RadonDB 2节点混合读写（比例7:2), 查询次数:两亿行

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           0
        total:                           200000016
    transactions:                        11111112 (1406.32 per sec.)
    queries:                             200000016 (25313.72 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7900.8538s
    total number of events:              11111112

Latency (ms):
         min:                                 21.44
         avg:                                364.07
         max:                               5382.09
         95th percentile:                    434.83
         sum:                            4045170030.32

Threads fairness:
    events (avg/stddev):           21701.3906/16.96
    execution time (avg/stddev):   7900.7227/0.05
```

### 2.7 RadonDB 1节点写2亿行

```
SQL statistics:
    queries performed:
        read:                            0
        write:                           200000000
        other:                           0
        total:                           200000000
    transactions:                        200000000 (11889.50 per sec.)
    queries:                             200000000 (11889.50 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          16821.5592s
    total number of events:              200000000

Latency (ms):
         min:                                  0.39
         avg:                                 43.06
         max:                              14673.89
         95th percentile:                     20.37
         sum:                            8612202713.71

Threads fairness:
    events (avg/stddev):           390625.0000/544.78
    execution time (avg/stddev):   16820.7084/0.06
```

### 2.8 RadonDB 1节点只读, 读取刚写入的2亿行数据，2 hours, 读取约1亿4千万条数据

```
SQL statistics:
    queries performed:
        read:                            51539418
        write:                           0
        other:                           0
        total:                           51539418
    transactions:                        3681387 (511.27 per sec.)
    queries:                             51539418 (7157.75 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.5012s
    total number of events:              3681387

Latency (ms):
         min:                                470.68
         avg:                               1001.40
         max:                               1609.69
         95th percentile:                   1089.30
         sum:                            3686552069.68

Threads fairness:
    events (avg/stddev):           7190.2090/4.63
    execution time (avg/stddev):   7200.2970/0.15
```

### 2.9 RadonDB 1节点混合读写(比例7:2), 查询次数:两亿行


# Mysql(版本5.7.20)

## 1 net_optane_neonsan

### 1.1 写2亿行

```
SQL statistics:
    queries performed:
        read:                            0    
        write:                           200000000
        other:                           0    
        total:                           200000000
    transactions:                        200000000 (5136.86 per sec.)
    queries:                             200000000 (5136.86 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          38934.2999s
    total number of events:              200000000

Latency (ms):
         min:                                  0.38 
         avg:                                 99.67
         max:                               2737.46
         95th percentile:                    292.60
         sum:                            19933513680.89

Threads fairness:
    events (avg/stddev):           390625.0000/1576.80
    execution time (avg/stddev):   38932.6439/0.09
```

### 1.2 读取刚写入的2亿行数据，2 hours

```
SQL statistics:
    queries performed:
        read:                            314722758
        write:                           0
        other:                           44960394
        total:                           359683152
    transactions:                        22480197 (3122.20 per sec.)
    queries:                             359683152 (49955.17 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.1169s
    total number of events:              22480197

Latency (ms):
         min:                                 18.01
         avg:                                163.98
         max:                                420.30
         95th percentile:                    244.38
         sum:                            3686379631.98

Threads fairness:
    events (avg/stddev):           43906.6348/8852.88
    execution time (avg/stddev):   7199.9602/0.04
```

### 1.3 混合读写(比例7:2), queries:两亿

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           44444448
        total:                           244444464
    transactions:                        11111112 (1763.59 per sec.)
    queries:                             244444464 (38799.02 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          6300.2710s
    total number of events:              11111112

Latency (ms):
         min:                                  5.00
         avg:                                290.31
         max:                               2158.16
         95th percentile:                    539.71
         sum:                            3225644758.86

Threads fairness:
    events (avg/stddev):           21701.3906/464.48
    execution time (avg/stddev):   6300.0874/0.09
```

## local optane

### 2.1 写2亿行

```
SQL statistics:
    queries performed:
        read:                            0
        write:                           200000000
        other:                           0
        total:                           200000000
    transactions:                        200000000 (22547.43 per sec.)
    queries:                             200000000 (22547.43 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          8870.1877s
    total number of events:              200000000

Latency (ms):
         min:                                  0.17
         avg:                                 22.70
         max:                                581.16
         95th percentile:                     70.55
         sum:                            4540829536.25

Threads fairness:
    events (avg/stddev):           390625.0000/5085.83
    execution time (avg/stddev):   8868.8077/0.04
```

### 2.2 读取刚写入的2亿行数据，2 hours

```
SQL statistics:
    queries performed:
        read:                            372534190
        write:                           0
        other:                           0
        total:                           372534190
    transactions:                        26609585 (3695.72 per sec.)
    queries:                             372534190 (51740.06 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          7200.1086s
    total number of events:              26609585

Latency (ms):
         min:                                 10.11
         avg:                                138.54
         max:                                374.15
         95th percentile:                    211.60
         sum:                            3686365944.02

Threads fairness:
    events (avg/stddev):           51971.8457/11367.52
    execution time (avg/stddev):   7199.9335/0.04
```

### 2.3 混合读写(比例7:2), queries:两亿

```
SQL statistics:
    queries performed:
        read:                            155555568
        write:                           44444448
        other:                           0
        total:                           200000016
    transactions:                        11111112 (2882.56 per sec.)
    queries:                             200000016 (51886.00 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          3854.6025s
    total number of events:              11111112

Latency (ms):
         min:                                 26.25
         avg:                                177.61
         max:                                475.00
         95th percentile:                    262.64
         sum:                            1973498070.74

Threads fairness:
    events (avg/stddev):           21701.3906/4366.46
    execution time (avg/stddev):   3854.4884/0.04
```



