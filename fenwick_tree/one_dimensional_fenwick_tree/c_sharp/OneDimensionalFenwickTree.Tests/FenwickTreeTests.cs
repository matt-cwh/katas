namespace OneDimensionalFenwickTree.Tests;

public class FenwickTreeTests
{
    [Theory]
    [InlineData("0,0,0,0,0,0,0,0", "5,5,0,5,0,0,0,5", 0, 5)]
    [InlineData("5,5,0,5,0,0,0,5", "5,5,2,7,0,0,0,7", 2, 2)]
    [InlineData("5,5,2,7,0,0,0,7", "5,5,2,7,4,4,0,11", 4, 4)]
    [InlineData("5,5,2,7,4,4,0,11", "5,5,2,7,7,7,0,14", 4, 7)]
    public void Update_Should_Update_Index_Value_And_Parents(string original, string expected, int i, int v)
    {
        var tree = original.Split(',').Select(x => int.Parse(x)).ToArray();
        var expectedTree = expected.Split(',').Select(x => int.Parse(x)).ToArray();
        FenwickTree.Update(tree, i, v);

        Assert.Equal(expectedTree, tree);
    }

    [Theory]
    [InlineData("5,5,2,7,4,4,0,11", 7, 11)]
    [InlineData("5,5,2,7,4,4,0,11", 4, 11)]
    [InlineData("5,5,2,7,4,4,0,11", 3, 7)]
    public void Get_Should_Return_Sum_From_Zero_To_Index(string treeValues, int i, int expected)
    {
          var tree = treeValues.Split(',').Select(x => int.Parse(x)).ToArray();

         var sum= FenwickTree.Get(tree, i);
           Assert.Equal(expected, sum);
    }
}