# process-inspection library
linux-only, at first. Then add Windows?

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


- get a list of all processes
- for each process, add data:
  - memory usage
  - cpu usage
  - open files
  - open network ports


## TODO
- write a darwin version?
- write a windows version?

- additional lookups: username
- use fmt.Sscanf() to scan data from /proc/[pid]/status?

- correct types for process data -- memory addresses, etc.
- get /proc/stat and expose as "host" metrics

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

## Dev

Build:
`GOOS=linux GOARCH=amd64 go build .`

Create and enter Docker environment:
```
docker run -d --rm -t --name gopsinspect -v "$(pwd):/root/" ubuntu:22.04
docker exec -it gopsinspect /bin/bash
```

Run:
`./root/gopsinspect`



# Design

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

