
def find_longest_palindromic_str(inputStr):

    inputLen = len(inputStr)
    start = 0
    end = 0
    maxLen = 0
    
    for i in range(inputLen):
        j = i
        k = i+1
        expand=0
        while(j-expand >0 and k+expand < inputLen-1 and  inputStr[j-expand] == inputStr[k+expand]):
            if(expand*2 > maxLen):
                maxLen = expand*2
                start = j-expand
                end = k+expand
            expand +=1
        
        m = i
        expand = 0
        while(m- expand> 0 and m+expand < inputLen-1 and inputStr[m-expand] == inputStr[m+expand]):
            if(2*expand+1 > maxLen):
                maxLen = (2*expand+1)
                start = m-expand
                end = m+expand
            expand+=1


    return (inputStr[start:end+1], end-start+1)
