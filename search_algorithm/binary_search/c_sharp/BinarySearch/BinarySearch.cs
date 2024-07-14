namespace BinarySearch;
public static class BinarySearch
{
    public static int Find(int[] sortedArray, int low, int high, int valueToSearch)
    {
        if (low >= high)
        {
            return -1;
        }

        var midPoint = high - low / 2;

        if (sortedArray[low + midPoint] == valueToSearch)
        {
            return low + midPoint;
        }

        var newLow = low;
        var newHigh = high;

        if (sortedArray[low + midPoint] < valueToSearch)
        {
            newLow = low + midPoint + 1;
        }
        else
        {
            newHigh = low + midPoint - 1;
        }

        return Find(sortedArray, newLow, newHigh, valueToSearch);
    }

}
