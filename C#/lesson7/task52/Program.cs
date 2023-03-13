using System;

class Program {
    static void Main(string[] args) {
        int[,] array = new int[,] {
            {1, 4, 7, 2},
            {5, 9, 2, 3},
            {8, 4, 2, 4}
        };

        int rows = array.GetLength(0);
        int cols = array.GetLength(1);

        for (int j = 0; j < cols; j++) {
            double sum = 0;
            for (int i = 0; i < rows; i++) {
                sum += array[i, j];
            }
            double average = sum / rows;
            Console.WriteLine("Среднее арифметическое столбца " + (j + 1) + ": " + average);
        }
    }
}
