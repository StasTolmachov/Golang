using System;

class Program
{
    static void Main(string[] args)
    {
        Console.Write("Введите номер дня недели (1-7): ");
        int dayOfWeek = int.Parse(Console.ReadLine());

        if (dayOfWeek == 6 || dayOfWeek == 7)
        {
            Console.WriteLine("Это выходной день");
        }
        else
        {
            Console.WriteLine("Это рабочий день");
        }
    }
}
