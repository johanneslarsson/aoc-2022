# Day 6

a) ([ifs](a/main.go))
```
pkg: aoc2022/06/a
BenchmarkPartOne-8   	  646744	      1755 ns/op
BenchmarkPartTwo-8   	   56236	     20896 ns/op
```

b) ([set](b/main.go))
```
pkg: aoc2022/06/b
BenchmarkPartOne-8   	   18211	     63970 ns/op
BenchmarkPartTwo-8   	    1821	    623139 ns/op
```

c) ([array](c/main.go))
```
pkg: aoc2022/06/c
BenchmarkPartOne-8   	   52770	     21393 ns/op
BenchmarkPartTwo-8   	   19240	     63722 ns/op
```

d) ([for loops](d/main.go))
```
pkg: aoc2022/06/d
BenchmarkPartOne-8   	  213932	      4866 ns/op
BenchmarkPartTwo-8   	   29054	     40776 ns/op
```

e) ([set with delete instead make new](e/main.go))
```
pkg: aoc2022/06/e
BenchmarkPartOne-8   	    6686	    171676 ns/op
BenchmarkPartTwo-8   	    1028	   1146189 ns/op
```

f) ([bit operations](f/main.go))
```
pkg: aoc2022/06/f
BenchmarkPartOne-8   	  127954	      8191 ns/op
BenchmarkPartTwo-8   	   31873	     37686 ns/op
```