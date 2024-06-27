from quickSort import *

def test_quickSort():
    inputArray = [2,8,7,1,3,5,6,4]
    expectedArray = [1,2,3,4,5,6,7,8]
    high = len(inputArray)-1
    low = 0
    quickSort(inputArray, low, high)
    assert inputArray == expectedArray


def test_partition():
    inputArray = [2,8,7,1,3,5,6,4]
    expectedArray = [2,1,3,4,7,5,6,8]
    high = len(inputArray)-1
    low = 0
    pivotPosition = partition(inputArray, low, high)
    assert pivotPosition == 3
    assert inputArray == expectedArray

def test_partition_low_subarray():
    inputArray = [3,2,1,4,7,5,6,8]
    expectedArray = [1,2,3,4,7,5,6,8]
    high = 2
    low = 0
    partition(inputArray, low, high)
    assert inputArray == expectedArray
    
def test_swap():
    inputArray = [2,8,7,1,3,5,6,4]
    expectedArray = [2,6,7,1,3,5,8,4]

    swap(inputArray, 1, 6)

    assert inputArray == expectedArray



