from findMissingNumber import findMissingNumber


def test_findMissingNumber():
    inputArray = {1, 2, 4, 6, 3, 7, 8} 
    n = 8

    missingNumber = findMissingNumber(inputArray,n)

    assert missingNumber == 5
