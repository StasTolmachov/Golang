using System;

class Program {
    static void Main(string[] args) {
        double b1, k1, b2, k2;
        Console.WriteLine("Введите значения b1, k1, b2, k2 через пробел:");
        string[] input = Console.ReadLine().Split(' ');
        b1 = double.Parse(input[0]);
        k1 = double.Parse(input[1]);
        b2 = double.Parse(input[2]);
        k2 = double.Parse(input[3]);

        if (k1 == k2) {
            Console.WriteLine("Прямые параллельны, нет точки пересечения.");
        } else {
            double x = (b2 - b1) / (k1 - k2);
            double y = k1 * x + b1;
            Console.WriteLine("Точка пересечения: ({0}; {1})", x, y);
        }
    }
}
