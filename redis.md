# 10
redis-benchmark -t set -d 10
====== SET ======
100000 requests completed in 1.72 seconds
50 parallel clients
10 bytes payload
keep alive: 1

99.03% <= 1 milliseconds
99.99% <= 2 milliseconds
100.00% <= 2 milliseconds
58173.36 requests per second

redis-benchmark -t get -d 10
====== GET ======
100000 requests completed in 1.68 seconds
50 parallel clients
10 bytes payload
keep alive: 1

99.61% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 2 milliseconds
59488.40 requests per second

# 20
redis-benchmark -t get -d 20
====== GET ======
100000 requests completed in 1.73 seconds
50 parallel clients
20 bytes payload
keep alive: 1

99.74% <= 1 milliseconds
100.00% <= 1 milliseconds
57870.37 requests per second



redis-benchmark -t set -d 20
====== SET ======
100000 requests completed in 1.73 seconds
50 parallel clients
20 bytes payload
keep alive: 1

99.11% <= 1 milliseconds
99.98% <= 2 milliseconds
100.00% <= 2 milliseconds
57937.43 requests per second

# 50
redis-benchmark -t set -d 50
====== SET ======
100000 requests completed in 1.75 seconds
50 parallel clients
50 bytes payload
keep alive: 1

99.83% <= 1 milliseconds
100.00% <= 1 milliseconds
57077.62 requests per second


redis-benchmark -t get -d 50
====== GET ======
100000 requests completed in 1.71 seconds
50 parallel clients
50 bytes payload
keep alive: 1

99.57% <= 1 milliseconds
100.00% <= 1 milliseconds
58411.21 requests per second

# 100
redis-benchmark -t get -d 100
====== GET ======
100000 requests completed in 1.77 seconds
50 parallel clients
100 bytes payload
keep alive: 1

98.71% <= 1 milliseconds
100.00% <= 1 milliseconds
56465.27 requests per second


redis-benchmark -t set -d 100
====== SET ======
100000 requests completed in 1.73 seconds
50 parallel clients
100 bytes payload
keep alive: 1

98.86% <= 1 milliseconds
99.97% <= 2 milliseconds
99.98% <= 3 milliseconds
100.00% <= 3 milliseconds
57803.47 requests per second

# 200
redis-benchmark -t set -d 200
====== SET ======
100000 requests completed in 1.69 seconds
50 parallel clients
200 bytes payload
keep alive: 1

99.77% <= 1 milliseconds
100.00% <= 1 milliseconds
59206.63 requests per second


redis-benchmark -t get -d 200
====== GET ======
100000 requests completed in 1.73 seconds
50 parallel clients
200 bytes payload
keep alive: 1

99.82% <= 1 milliseconds
100.00% <= 1 milliseconds
57670.13 requests per second

# 1000
redis-benchmark -t get -d 1000
====== GET ======
100000 requests completed in 1.81 seconds
50 parallel clients
1000 bytes payload
keep alive: 1

99.75% <= 1 milliseconds
100.00% <= 2 milliseconds
55279.16 requests per second



redis-benchmark -t set -d 1000
====== SET ======
100000 requests completed in 1.76 seconds
50 parallel clients
1000 bytes payload
keep alive: 1

99.73% <= 1 milliseconds
100.00% <= 1 milliseconds
56721.50 requests per second

# 5000
redis-benchmark -t set -d 5000
====== SET ======
100000 requests completed in 1.88 seconds
50 parallel clients
5000 bytes payload
keep alive: 1

99.61% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 2 milliseconds
53333.33 requests per second



redis-benchmark -t get  -d 5000
====== GET ======
100000 requests completed in 1.94 seconds
50 parallel clients
5000 bytes payload
keep alive: 1

98.93% <= 1 milliseconds
99.99% <= 2 milliseconds
100.00% <= 2 milliseconds
51626.23 requests per second

第二项作业没有看懂