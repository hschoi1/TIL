# Cycle detection problem: https://en.wikipedia.org/wiki/Cycle_detection
# complexity: O(cycleLength+startCycle) time and O(1) space



def floyd(f, x0):
    #if there is a cycle, x_i = x_{i+k*l} -> then i=k*l iff x_i = x_{2i}
    tortoise = f(x0)
    hare = f(f(x0))
    # find the multiple of the position where they intersect
    while tortoise != hare:
        tortoise = f(tortoise) # one step
        hare = f(f(hare)) # twice the speed
    #tortoise and hare now at the same spot, distance between them is v=m*l

    # x_{startCycle}  = x_{startCycle+m*l}
    startPos = 0
    tortoise = x0
    while tortoise != hare:
        tortoise = f(tortoise)
        hare = f(hare)
        startPos += 1

    #now beginning at the start of the cycle, find legnth of cycle by moving hare one by one
    cycleLength = 1
    hare = f(tortoise)
    while tortoise != hare:
        hare = f(hare)
        cycleLength += 1

    return pos, cycleLength
