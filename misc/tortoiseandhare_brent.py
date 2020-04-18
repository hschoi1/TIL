#https://en.wikipedia.org/wiki/Cycle_detection
# find the smallest power of two that is larget than both cycleLength and startPosOfCycle
# this algorithm finds the cycleLength directly
def brent(f, x0):
    power, cycleLength = 1 
    tortoise = x0 
    hare = f(x0) 
    # find the smallest power of two greater than cycleLength and startPos
    while tortoise != hare: 
        if power == cycleLength: 
            tortoise = hare  # before the tortoiise and hare meet, tortoise is inside the loop
            power *= 2 
            cycleLength = 0
        hare  = f(hare)
        cycleLength + = 1

    tortoise, hare = x0 
    for i in range(cycleLength):
        hare = f(hare)
    
    # now the distnace between them is the cycleLength
    startPos = 0
    while tortoise != hare: # when they meet they meet at the start of the cycle
        tortoise = f(tortoise)
        hare = f(hare)
        startPos += 1 
    
    return cycleLength, startPos

# complexity: O(cycleLength+startPosOfCycle) time and O(1) space