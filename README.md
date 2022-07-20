## go-geofence is a library to perform point-in-polygon searches in Golang

_Forked from weilunwu/go-geofence to add go.mod and remove dependences_

### Advantages compared with golang-geo

1. go-geofence uses a tiled cache to store pre-computed search results so it can determine inclusion very efficiently. Therefore the library is tailored for create once, query many times uses.

2. go-geofence is 4 times faster than kellydunn's golang geo for checking whether a point is inside a polygon.

Note: The _holes_ feature has been removed as I just didn't need it!

### Benchmark results:

* BenchmarkGeofence	10000000	       109 ns/op
* Benchmark kellydunn/golang-geo's GeoContains	 3000000	       475 ns/op

Detailed benchmark tests can be found in geofence_test.go
