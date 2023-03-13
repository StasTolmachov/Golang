using System;

class Program
{
    static void Main(string[] args)
    {
        Console.Write("Введите количество чисел: ");
        int m = int.Parse(Console.ReadLine());

        int count = 0;

        Console.WriteLine("Введите числа:");
        for (int i = 0; i < m; i++)
        {
            int num = int.Parse(Console.ReadLine());

            if (num > 0)
            {
                count++;
            }
        }

        Console.WriteLine("Количество чисел, больших 0: " + count);

        Console.ReadKey();
    }
}
