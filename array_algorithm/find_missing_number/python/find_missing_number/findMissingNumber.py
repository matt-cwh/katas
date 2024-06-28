
def findMissingNumber(inputArray, n):

    tempArray = [0]*n

    print(tempArray)

    for i in inputArray:
        tempArray[i-1] = 1

    for j in range(n):
        if(tempArray[j] == 0):
            return j +1
        
    return -1
