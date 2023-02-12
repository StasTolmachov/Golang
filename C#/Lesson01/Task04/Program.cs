// Задача 4: Напишите программу, которая принимает на вход три числа и выдаёт максимальное из этих чисел.

Console.Clear();
Console.Write("Введите A: ");
int numberA = int.Parse(Console.ReadLine());
Console.Write("Введите B: ");
int numberB = int.Parse(Console.ReadLine());
Console.Write("Введите C: ");
int numberC = int.Parse(Console.ReadLine());
int numberMax = numberA;
if (numberMax < numberB)
{
    numberMax = numberB;
}
if (numberMax < numberC)
{
    numberMax = numberC;
}


Console.Write("Максимальное число:");
Console.Write(numberMax);