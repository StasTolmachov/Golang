using System;

class Program
{
    static void Main(string[] args)
    {
        int[,] array = {
            {1, 4, 7, 2},
            {5, 9, 2, 3},
            {8, 4, 2, 4},
            {5, 2, 6, 7}
        };

        int smallestSumRow = 0;
        int smallestSum = int.MaxValue;

        for (int i = 0; i < array.GetLength(0); i++)
        {
            int sum = 0;

            for (int j = 0; j < array.GetLength(1); j++)
            {
                sum += array[i, j];
            }

            if (sum < smallestSum)
            {
                smallestSum = sum;
                smallestSumRow = i;
            }
        }

        Console.WriteLine("Номер строки с наименьшей суммой элементов: " + (smallestSumRow + 1));
    }
}
