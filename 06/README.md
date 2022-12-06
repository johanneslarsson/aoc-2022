# Day 6

a) ([ifs](a/main.go))
```
pkg: aoc2022/06/a
BenchmarkPartOne-12    	  618633	      1916 ns/op
BenchmarkPartTwo-12    	   49843	     23438 ns/op
```

b) ([set](b/main.go))
```
pkg: aoc2022/06/b
BenchmarkPartOne-12    	   18159	     66114 ns/op
BenchmarkPartTwo-12    	    1747	    659222 ns/op
```

c) ([array](c/main.go))
```
pkg: aoc2022/06/c
BenchmarkPartOne-12    	   68722	     16646 ns/op
BenchmarkPartTwo-12    	   23130	     51808 ns/op
```

d) ([for loops](d/main.go))
```
pkg: aoc2022/06/d
BenchmarkPartOne-12    	  171638	      6658 ns/op
BenchmarkPartTwo-12    	   26151	     46604 ns/op
```

e) ([set with delete instead make new](e/main.go))
```
pkg: aoc2022/06/e
BenchmarkPartOne-12    	   16126	     74089 ns/op
BenchmarkPartTwo-12    	    2486	    479418 ns/op
```

f) ([set with reassign keys instead of make new](f/main.go))
```
pkg: aoc2022/06/f
BenchmarkPartOne-12    	    1592	    641736 ns/op
BenchmarkPartTwo-12    	     589	   1904948 ns/op
```