#https://www.tutorialspoint.com/python/python_multithreading.htm
"""
1. multiple threads within a process share the same data space with the main thread
2. threads do not require much memory overhead

A thread has a beginning, an execution sequence, and a conclusion.
It can be pre-empted
It can sleep while other threads are running - yielding
"""

import thread # deprecated in python3. so could not run this program..
import time

def print_time(threadName, delay):
    count = 0
    while count < 5:
        time.sleep(delay)
        count += 1
        print("%s: %s" % (threadName, time.ctime(time.time())))

try:
    thread.start_new_thread(print_time, ("Thread-1", 2, ))
    thread.start_new_thread(print_time, ("Thread-2", 4, ))
except:
    print("Error: unable to start thread")

while 1:
    pass
