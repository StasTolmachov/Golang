using System;

class Program
{
    static void Main(string[] args)
    {
        double[] array = { 3.5, 7.2, 22.1, 2.8, 78.4 };

        double min = array[0];
        double max = array[0];

        for (int i = 1; i < array.Length; i++)
        {
            if (array[i] < min)
            {
                min = array[i];
            }

            if (array[i] > max)
            {
                max = array[i];
            }
        }

        double difference = max - min;

        Console.WriteLine("Разница между максимальным и минимальным элементами массива: " + difference);

        Console.ReadKey();
    }
}
