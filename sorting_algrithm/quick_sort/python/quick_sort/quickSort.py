

def quickSort(inputArray, low, high):
    print(low)
    print(high)
    if(low < high):
        pivot = partition(inputArray,low,high)
        quickSort(inputArray, low,pivot-1)
        quickSort(inputArray, pivot+1, high)

def partition(inputArray, low, high):
    pivot = inputArray[high]
    i = low -1
    
    for j in range(high-low):
        if(inputArray[j+low] <  pivot):
            i = i+1
            swap(inputArray, i,j+low)
    
    swap(inputArray, i+1, high)

    return i+1
        


def swap(inputArray, i, j):
     tmp = inputArray[i]
     inputArray[i] = inputArray[j]
     inputArray[j] = tmp
       
   
    
