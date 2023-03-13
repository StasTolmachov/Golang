using System;

class Program
{
    static void Main(string[] args)
    {
        Console.Write("Введите число A: ");
        int a = int.Parse(Console.ReadLine());

        Console.Write("Введите число B: ");
        int b = int.Parse(Console.ReadLine());

        int result = 1;

        for (int i = 1; i <= b; i++)
        {
            result *= a;
        }

        Console.WriteLine(a + " в степени " + b + " равно " + result);

        Console.ReadKey();
    }
}
