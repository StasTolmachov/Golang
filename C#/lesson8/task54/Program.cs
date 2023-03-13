int[,] array = new int[,] {
    { 1, 4, 7, 2 },
    { 5, 9, 2, 3 },
    { 8, 4, 2, 4 }
};

// Проходимся по каждой строке массива и сортируем ее элементы по убыванию
for (int i = 0; i < array.GetLength(0); i++)
{
    // Преобразуем строку в массив, чтобы отсортировать его элементы
    int[] row = new int[array.GetLength(1)];
    for (int j = 0; j < array.GetLength(1); j++)
    {
        row[j] = array[i, j];
    }

    // Сортируем элементы по убыванию
    Array.Sort(row);
    Array.Reverse(row);

    // Записываем отсортированные элементы обратно в строку
    for (int j = 0; j < array.GetLength(1); j++)
    {
        array[i, j] = row[j];
    }
}

// Выводим отсортированный массив на экран
for (int i = 0; i < array.GetLength(0); i++)
{
    for (int j = 0; j < array.GetLength(1); j++)
    {
        Console.Write("{0} ", array[i, j]);
    }
    Console.WriteLine();
}
