using System;

class Program
{
    static void Main(string[] args)
    {
        int[] array = { 1, 2, 5, 7, 19 };

        Console.Write("[");
        for (int i = 0; i < array.Length; i++)
        {
            Console.Write(array[i]);
            if (i != array.Length - 1)
            {
                Console.Write(", ");
            }
        }
        Console.WriteLine("]");

        Console.ReadKey();
    }
}
