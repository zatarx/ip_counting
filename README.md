### Solution details
One of the most efficient ways all the IPv4 addresses can be stored in memory is by using a bitmap of all the potential addresses.
```
1) Each octet is a byte -> 256 values in each octet
2) We have 4 octets total -> 256**4 = 4_294_967_296 - total amount of bits needed to represent an entry for all the ipv4 ips
3) 4_294_967_296 / 8 = 536_870_912 (512MB) - byte array size to accomodate all the entries 
```

### Performance metrics
I used goroutines to split up ip counting.

```
File Size: 5.69GB
CPU Cores : 8
Peak Memory Usage: 1.6GB
Execution Time: 190-220s (~3-4 minutes)
Goroutine Execution Memory Overhead: 1.1GB
```

### Potential Improvements
* Use channels to restrict access on a byte level of the byte array. 
However, this could introduce additional memory overhead on storing a 
data structure to keep track of which bytes are being accessed and 
which are not.
