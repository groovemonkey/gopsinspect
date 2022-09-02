# process-inspection library

## TODO
- actually parse new windows command
- test windows version

- add logging - try a cool logging library

- correct types for process data -- memory addresses, etc.

- additional lookups: friendly username, mem%
  - https://unix.stackexchange.com/questions/496868/how-to-get-users-name-from-uid

- use fmt.Sscanf() to scan data from /proc/[pid]/status?
- get /proc/stat and expose as "host" metrics

- Additional data gathering:
  - open files
  - open network ports

- starttime, stime, utime, cstime, cutime: get sysconf(_SC_CLK_TCK) and divide by that
- somehow mark deprecated/old values?

- add useful "String()" methods so that e.g. a processlist prints out similar to `ps aux`
- Actually implement Sort() for a PsListing

- benchmark -- is all of this faster than just running and parsing `top` and `ps aux`?
  - is all of this code actually worth it?
  - aug 31, 2022 - calling `ps auxh` vs looking it up ourselves:
    - passthrough (maybe because spawning a shell is slow?)
      real	0m0.034s
      user	0m0.000s
      sys	0m0.014s

    - looking it up ourselves:
      real	0m0.004s
      user	0m0.000s
      sys	0m0.004s

    - raw ps auxh
      real	0m0.004s
      user	0m0.000s
      sys	0m0.004s


- input-parsing map function
  - how do I design this to be consumed easily too?
  - map["Key"]"value"
  - then each tool-result-representation struct can somehow have a string lookup to its attrs? (no I don't think Go works that way) psInfo.Set("Key", "value") -> does it or logs an error and continues?

## Dev

Build:
`GOOS=linux GOARCH=amd64 go build -o bin/ .`

Create and enter Docker environment:
```
docker run -d --rm -t --name gopsinspect -v "$(pwd):/tmp/gopsinspect" ubuntu:22.04
docker exec -it gopsinspect /bin/bash
```

Run:
`./root/gopsinspect`


## Use as a library

See `cmd/gopsinspect/main.go`

# Original Design Notes

v0.1: Mimic Information in `ps`:
```
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.0   4492  3556 pts/0    Ss+  05:34   0:00 bash
```

ALSO: mimic information in "top":
  - uptime (days, hours)
  - users
  - load averages
  - tasks (total, running, sleeping, stopped, zombie)
  - CPU: user, systems, ni, id, wa, hi, si, st
  - Mem: total, free, used, buff/cache
  - Swap: total, free, used, avail

Some combined version of this?
## system usage
/proc/stat

## per process
`/proc/1/status`

### Basics
name, umask, state (refresh), pid, ppid, uid, gid, groups, seccomp, fdsize

### Memory
vmpeak/vmsize/vm*

  - more memory data? Avgs?


### CPU
CPU: thp, threads, voluntary/nonvoluntary_ctxt_switches




  - time since start
  - total cpu time
  - cpu usage (point-in-time, avg)


### File handles
/proc/9/fd

- resolve pathnames for file descriptors
- perms and other ls -l data for each one?

