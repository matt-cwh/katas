
namespace MaxSubarraySum;
public static class MaxSubarraySumCalculator
{
    public static int Calculate(int[] inputArray)
    {
        var curMax = 0;
        var curEnd = 0;

        foreach(var item in inputArray){

            curEnd +=  item;

            if(curEnd > curMax){
                curMax = curEnd;
            }

            if(curEnd<0){
                curEnd = 0;
            }
        }

        return curMax;
    }
}
