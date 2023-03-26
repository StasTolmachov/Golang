using System;

class Program
{
    static void Main()
    {
        // Исходный массив строк
        string[] source = { "one", "two", "three", "four", "five", "six" };

        // Создаем массив для хранения строк длиной <= 3 символов
        string[] result = new string[source.Length];
        int count = 0;

        // Проходим по всем элементам исходного массива
        for (int i = 0; i < source.Length; i++)
        {
            // Если длина строки <= 3 символов, то добавляем ее в результирующий массив
            if (source[i].Length <= 3)
            {
                result[count] = source[i];
                count++;
            }
        }
    }
}
