using System.Data;

namespace OneDimensionalFenwickTree;
public static class FenwickTree
{
    public static int Get(int[] tree, int index)
    {
        var sum = 0;
        for (int i = index + 1; i > 0; i -= i & -i)
        {
            sum += tree[i - 1];
        }

        return sum;
    }

    public static void Update(int[] tree, int index, int value)
    {
        // We get the value of current sum and subtract the sum with index-1 to get the actual value
        var currentValue = Get(tree, index)-Get(tree, index-1);
        var diff = value - currentValue;

        for (int i = index + 1; i <= tree.Length; i += (i & -i))
        {
            tree[i - 1] += diff;
        }
    }

}
