

def merge(arr1, arr2):
    arr1Len = len(arr1)
    arr2Len = len(arr2)
    rLen = arr1Len + arr2Len
    r = [0] * rLen

    i = 0
    j = 0

    for k in range(rLen):
        if(j>= arr2Len):
            r[k] = arr1[i]
            i +=1
        elif(i>= arr1Len):
            r[k] = arr2[j]
            j+=1
        elif(arr1[i] <= arr2[j]):
            r[k] = arr1[i]
            i +=1
        else:
            r[k] = arr2[j]
            j +=1
     

    return r
            
def mergeSort(arr):
    arrLen = len(arr)
    if(arrLen<2):
        return arr
    
    mid = arrLen//2

    lowerHalf = mergeSort(arr[0:mid])
    upperHalf = mergeSort(arr[mid:arrLen])

    return merge(lowerHalf,upperHalf)
