
using System.ComponentModel.DataAnnotations;

namespace TrappingRainWater;
public class WaterSumCalculator
{
    public static (int[] LeftMaxArray, int[] RightMaxArray) FindMaxArrays(int[] inputArray)
    {
        var rightMaxArray = new int[inputArray.Length];
        var leftMaxArray = new int[inputArray.Length];

        var curLeftMax = 0;
        var curRightMax = 0;
        for (int i = 0; i < inputArray.Length; i++)
        {
            if (curLeftMax <= inputArray[i])
            {
                curLeftMax = inputArray[i];
            }

            var rightIndex = inputArray.Length - 1 - i;

            if (curRightMax <= inputArray[rightIndex])
            {
                curRightMax = inputArray[rightIndex];
            }

            leftMaxArray[i] = curLeftMax;
            rightMaxArray[rightIndex] = curRightMax;
        }

        return (leftMaxArray, rightMaxArray);
    }

    public static int FindSumWithA2Pointers(int[] inputArray)
    {
        int left = 0;
        int right = inputArray.Length - 1;
        int leftMax = 0;
        int rightMax = 0;
        int sum = 0;

        while(left < right)
        {
            leftMax = Math.Max(inputArray[left], leftMax);
            rightMax = Math.Max(inputArray[right], rightMax);

            if(leftMax > rightMax){
                sum += rightMax - inputArray[right];
                right --;
            }
            else{
                sum+= leftMax - inputArray[left];
                left ++;
            }
        }

        return sum;
    }

    public static int FindSumWithArrays(int[] inputArray)
    {
        var (leftMaxArray, rightMaxArray) = FindMaxArrays(inputArray);
        var sumOfWater = 0;

        for (int i = 0; i < inputArray.Length; i++)
        {
            var maxLevel = Math.Min(leftMaxArray[i], rightMaxArray[i]);
            sumOfWater += maxLevel - inputArray[i];
        }

        return sumOfWater;
    }
}