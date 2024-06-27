using MaxSubarraySum;

namespace MaxSubarraySum.Tests;

public class MaxSubarraySumCalculatorTest
{
    [Fact]
    public void Calculate_Should_Return_Max_Sum_Of_Subarray()
    {
        var inputArray = new []{-2,-3,4,-1,-2,1,5,-3};

        var sum = MaxSubarraySumCalculator.Calculate(inputArray);

        Assert.Equal(7,sum);

    }
}