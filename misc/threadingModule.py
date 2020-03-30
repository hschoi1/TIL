#https://www.tutorialspoint.com/python/python_multithreading.htm
"""
The threading module provides much more powerful, high-level support for threads

threading.activeCount() - returns the # of thread objects that are active
threading.currentThread() - returns the # of thread objects in the caller's thread control
threading.enumerate() = returns a list of all thread objects that are currently active

methods provided by the Thread class
run() - entry point for a thread
start() - starts a thread by calling the run method
join([time]) - waits for threads to terminate
isAlive() - checks whether a thread is still executing
getName() - returns the name of a thread
setName() - sets the name of a thread

To implement a new thread,
1.define a new subclass of the Thread class
2. override the __init__(self, [,args]) to add additional arguments
3. override the run(self, [,args])

Then create an instance of it and then start a new thread by invoking start()
"""

import threading
import time

exitFlag = 0
class myThread(threading.Thread):
    def __init__ (self, threadID, name, counter):
        threading.Thread.__init__(self)
        self.threadID = threadID
        self.name = name
        self.counter = counter

    def run(self):
        print("starting", self.name)
        print_time(self.name, 5, self.counter)
        print("exiting", self.name)

def print_time(threadName, counter, delay):
    while counter:
        if exitFlag:
            threadName.exit()
        time.sleep(delay)
        print("%s: %s" % (threadName, time.ctime(time.time())))
        counter -= 1


thread1 = myThread(1, "Thread-1", 1)
thread2 = myThread(2, "Thread-2", 2)

thread1.start()
thread2.start()

print("Exiting Main Thread")

"""
starting Thread-1
starting Thread-2
Exiting Main Thread
Thread-1: Mon Mar 30 22:10:25 2020
Thread-2: Mon Mar 30 22:10:26 2020
Thread-1: Mon Mar 30 22:10:26 2020
Thread-1: Mon Mar 30 22:10:27 2020
Thread-2: Mon Mar 30 22:10:28 2020
Thread-1: Mon Mar 30 22:10:28 2020
Thread-1: Mon Mar 30 22:10:29 2020
exiting Thread-1
Thread-2: Mon Mar 30 22:10:30 2020
Thread-2: Mon Mar 30 22:10:32 2020
Thread-2: Mon Mar 30 22:10:34 2020
exiting Thread-2
"""


class myThread2(threading.Thread):
    def __init__ (self, threadID, name, counter):
        threading.Thread.__init__(self)
        self.threadID = threadID
        self.name = name
        self.counter = counter

    def run(self):
        print("starting", self.name)
        threadLock.acquire()
        print_time2(self.name, 3, self.counter)
        threadLock.release()
        print("exiting", self.name)

def print_time2(threadName, counter, delay):
    while counter:
        time.sleep(delay)
        print("%s: %s" % (threadName, time.ctime(time.time())))
        counter -= 1


threadLock = threading.Lock()
threads = []
thread1 = myThread2(1, "Thread-1", 1)
thread2 = myThread2(2, "Thread-2", 2)

print("Starting Main Thread2")

thread1.start()
thread2.start()

threads.append(thread1)
threads.append(thread2)

for t in threads:
    t.join()
print("Exiting Main Thread2")

"""
Starting Main Thread2
starting Thread-1
starting Thread-2
Thread-1: Mon Mar 30 22:51:52 2020
Thread-1: Mon Mar 30 22:51:53 2020
Thread-1: Mon Mar 30 22:51:54 2020
exiting Thread-1
Thread-2: Mon Mar 30 22:51:57 2020
Thread-2: Mon Mar 30 22:51:59 2020
Thread-2: Mon Mar 30 22:52:01 2020
exiting Thread-2
Exiting Main Thread2
"""
