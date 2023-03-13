using System;

class Program {
    static void Main(string[] args) {
        int m = 3, n = 4;
        double[,] array = new double[m, n];
        Random rand = new Random();

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                array[i, j] = rand.NextDouble() * 20 - 10; // случайное число в диапазоне [-10, 10)
                Console.Write(array[i, j] + " ");
            }
            Console.WriteLine();
        }
    }
}
