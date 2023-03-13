using System;

class Program
{
    static void Main(string[] args)
    {
        int m = 60;
        int n = 48;

        int gcd = Gcd(m, n);

        Console.WriteLine($"НОД({m}, {n}) = {gcd}");
    }

    static int Gcd(int a, int b)
    {
        if (b == 0)
        {
            return a;
        }
        else
        {
            return Gcd(b, a % b);
        }
    }
}
