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
            array[i] = random.Next(-10, 10);
            Console.Write(array[i]);
            if (i != array.Length - 1)
            {
                Console.Write(", ");
            }
        }
        Console.WriteLine("]");

        int sum = 0;
        for (int i = 0; i < array.Length; i += 2)
        {
            sum += array[i];
        }

        Console.WriteLine("Сумма элементов, стоящих на нечётных позициях: " + sum);

        Console.ReadKey();
    }
}
