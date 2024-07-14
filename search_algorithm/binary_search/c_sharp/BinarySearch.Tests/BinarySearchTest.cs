namespace BinarySearch.Tests;

public class BinarySearchTest
{
    [Theory]
    [InlineData("-5,-2,0,1,2,4,5,6,7,10", 4, 5)]
    [InlineData("-5,-2,0,1,2,4,5,6,7,10", -2, 1)]
    [InlineData("-5,-2,0,1,2,4,5,6,7,10", 10, 9)]
    [InlineData("-5,-2,0,1,2,4,5,6,7,10", 3, -1)]
    public void Search_Should_Return_Index_Of_The_Query_Value(string inputValues, int valueToSearch, int expectedIndex)
    {
        var inputArray = inputValues.Split(",").Select(x => int.Parse(x)).ToArray();

        var index = BinarySearch.Find(inputArray, 0, inputArray.Length-1, valueToSearch);

        Assert.Equal(expectedIndex, index);

    }
}