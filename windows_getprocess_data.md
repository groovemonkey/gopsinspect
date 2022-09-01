# Sample data for windows
Sample data returned by (unprivileged) command: `powershell "Get-Process | Format-List *"`

The returned data starts with two empty lines; processes are delimited by one empty line between them (Container\n\nName)

```
Name                       : Spectrum
Id                         : 14328
PriorityClass              :
FileVersion                :
HandleCount                : 276
WorkingSet                 : 15527936
PagedMemorySize            : 6856704
PrivateMemorySize          : 6856704
VirtualMemorySize          : 105459712
TotalProcessorTime         :
SI                         : 0
Handles                    : 276
VM                         : 2203423682560
WS                         : 15527936
PM                         : 6856704
NPM                        : 14576
Path                       :
Company                    :
CPU                        :
ProductVersion             :
Description                :
Product                    :
__NounName                 : Process
BasePriority               : 8
ExitCode                   :
HasExited                  :
ExitTime                   :
Handle                     :
SafeHandle                 :
MachineName                : .
MainWindowHandle           : 0
MainWindowTitle            :
MainModule                 :
MaxWorkingSet              :
MinWorkingSet              :
Modules                    :
NonpagedSystemMemorySize   : 14576
NonpagedSystemMemorySize64 : 14576
PagedMemorySize64          : 6856704
PagedSystemMemorySize      : 148048
PagedSystemMemorySize64    : 148048
PeakPagedMemorySize        : 7217152
PeakPagedMemorySize64      : 7217152
PeakWorkingSet             : 17928192
PeakWorkingSet64           : 17928192
PeakVirtualMemorySize      : 112152576
PeakVirtualMemorySize64    : 2203430375424
PriorityBoostEnabled       :
PrivateMemorySize64        : 6856704
PrivilegedProcessorTime    :
ProcessName                : Spectrum
ProcessorAffinity          :
Responding                 : True
SessionId                  : 0
StartInfo                  : System.Diagnostics.ProcessStartInfo
StartTime                  :
SynchronizingObject        :
Threads                    : {9292, 9656, 2632}
UserProcessorTime          :
VirtualMemorySize64        : 2203423682560
EnableRaisingEvents        : False
StandardInput              :
StandardOutput             :
StandardError              :
WorkingSet64               : 15527936
Site                       :
Container                  :
```