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
        for (int i = index + 1; i <= tree.Length; i += (i & -i))
        {
            tree[i - 1] += value;
        }
    }

}
