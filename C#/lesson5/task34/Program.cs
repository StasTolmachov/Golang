using System;

class Program
{
    static void Main(string[] args)
    {
        int[] array = new int[4];
        Random random = new Random();

        Console.Write("Сгенерированный массив: [");
        for (int i = 0; i < array.Length; i++)
        {
            array[i] = random.Next(100, 1000);
            Console.Write(array[i]);
            if (i != array.Length - 1)
            {
                Console.Write(", ");
            }
        }
        Console.WriteLine("]");

        int count = 0;
        for (int i = 0; i < array.Length; i++)
        {
            if (array[i] % 2 == 0)
            {
                count++;
            }
        }

        Console.WriteLine("Количество чётных чисел в массиве: " + count);

        Console.ReadKey();
    }
}
