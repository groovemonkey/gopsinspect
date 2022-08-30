# process-inspection library
linux-only, at first. Then add Windows?

- get a list of all processes
- for each process, add data:
  - memory usage
  - cpu usage
  - open files
  - open network ports


## TODO
- get /proc/stat and add it as "host" metrics or something
  - maybe inspector.HostStats and inspector.ProcStats?

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

