using System;

class Program
{
    static void Main(string[] args)
    {
        int m = 1;
        int n = 10;

        Console.WriteLine("Четные числа в промежутке от {0} до {1}:", m, n);
        PrintEvenNumbers(m, n);
    }

    static void PrintEvenNumbers(int current, int end)
    {
        if (current > end)
        {
            return;
        }

        if (current % 2 == 0)
        {
            Console.WriteLine(current);
        }

        PrintEvenNumbers(current + 1, end);
    }
}
