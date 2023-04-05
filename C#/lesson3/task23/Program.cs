using System;

class Program
{
    static void Main(string[] args)
    {
        Console.Write("Введите число N: ");
        int n = int.Parse(Console.ReadLine());

        Console.WriteLine("Таблица кубов чисел от 1 до " + n + ":");

        for (int i = 1; i <= n; i++)
        {
            int cube = i * i * i;
            Console.Write(cube + " ");
        }

        Console.ReadKey();
    }
}
