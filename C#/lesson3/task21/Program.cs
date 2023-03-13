using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Введите координаты точки A (x,y,z):");
        double x1 = Convert.ToDouble(Console.ReadLine());
        double y1 = Convert.ToDouble(Console.ReadLine());
        double z1 = Convert.ToDouble(Console.ReadLine());

        Console.WriteLine("Введите координаты точки B (x,y,z):");
        double x2 = Convert.ToDouble(Console.ReadLine());
        double y2 = Convert.ToDouble(Console.ReadLine());
        double z2 = Convert.ToDouble(Console.ReadLine());

        double distance = Math.Sqrt(Math.Pow(x2 - x1, 2) + Math.Pow(y2 - y1, 2) + Math.Pow(z2 - z1, 2));
        Console.WriteLine("Расстояние между точками A и B: {0:F2}", distance);
    }
}
