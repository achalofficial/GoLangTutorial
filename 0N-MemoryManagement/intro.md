# Memory Management 

## new()
    1.  Allocate memory but no int
    2.  will get a memory address
    3.  zeroed stroage 

## make()
    1.  Allocated memory and INIT
    2.  will get a memory address
    3.  non-zeroed storage

## Garbage Collection 
    1.  happens automatically 
    2.  Anything which is nil or Out of Scope is eligible for Garbage collection 
    3.  GC happens when a specific percentage of memory is met. this percentage is stored in GOGC variable 
    4.  GOGC=100 by default and setting GOGC=off disables the garbage collector entirely
    