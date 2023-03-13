using System;

class Program {
    static void Main(string[] args) {
        int[,] array = new int[,] {
            {1, 4, 7, 2},
            {5, 9, 2, 3},
            {8, 4, 2, 4}
        };

        Console.WriteLine("Введите индексы элемента через запятую:");
        string[] input = Console.ReadLine().Split(',');
        int row = int.Parse(input[0]) - 1;
        int col = int.Parse(input[1]) - 1;

        if (row < 0 || row >= array.GetLength(0) || col < 0 || col >= array.GetLength(1)) {
            Console.WriteLine("Элемента с данными индексами в массиве нет.");
        } else {
            Console.WriteLine("Значение элемента: " + array[row, col]);
        }
    }
}
