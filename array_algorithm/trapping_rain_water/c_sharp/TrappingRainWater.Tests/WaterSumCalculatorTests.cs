namespace TrappingRainWater.Tests;

public class WaterSumCalculatorTests
{
    [Fact]
    public void FindMaxArrays_Should_Return_Max_Left_And_Right_Arrays()
    {
        var inputArray = new int[] { 3, 0, 1, 0, 4, 0, 2 };
        var expectedRightMaxArray = new int[] { 4, 4, 4, 4, 4, 2, 2 };
        var expectedLeftMaxArray = new List<int> { 3, 3, 3, 3, 4, 4, 4 };

        var (leftMaxArray, rightMaxArray) = WaterSumCalculator.FindMaxArrays(inputArray);

        Assert.Equal(expectedLeftMaxArray, leftMaxArray);
        Assert.Equal(expectedRightMaxArray, rightMaxArray);
    }

    [Fact]
    public void FindSumWithArrays_Should_Return_Sum_Of_Water()
    {
        var inputArray = new int[] { 3, 0, 1, 0, 4, 0, 2 };
        var result = WaterSumCalculator.FindSumWithArrays(inputArray);
        Assert.Equal(10, result);
    }

    [Fact]
    public void FindSumWith2Pointers_Should_Return_Sum_Of_Water()
    {
        var inputArray = new int[] { 3, 0, 1, 0, 4, 0, 2 };
        var result = WaterSumCalculator.FindSumWithA2Pointers(inputArray);
        Assert.Equal(10, result);
    }
}