using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Введите пятизначное число:");
        int number = Convert.ToInt32(Console.ReadLine());
        int temp = number;
        int reverse = 0;
        while (temp != 0)
        {
            int remainder = temp % 10;
            reverse = reverse * 10 + remainder;
            temp /= 10;
        }
        if (number == reverse)
        {
            Console.WriteLine("Число является палиндромом");
        }
        else
        {
            Console.WriteLine("Число не является палиндромом");
        }
    }
}
