// Задача 6: Напишите программу, которая на вход принимает число и выдаёт, является ли число чётным (делится ли оно на два без остатка).

Console.Clear();
Console.Write("Введите A: ");
int numberA = int.Parse(Console.ReadLine());
if (numberA % 2 == 0)
{
    Console.Write("Число четное");
}
else
{
    Console.Write("Число не четное");
}