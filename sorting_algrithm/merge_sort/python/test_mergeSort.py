from mergeSort import *


def test_merge():
    a1 = [8,25]
    a2 = [4,32]

    result = merge(a1, a2)

    assert result == [4,8,25,32]

def test_mergeSort():
    a1 = [7,3,2,6,8,4,5,1,2]
    
    result = mergeSort(a1)
    assert result == [1,2,2,3,4,5,6,7,8]
