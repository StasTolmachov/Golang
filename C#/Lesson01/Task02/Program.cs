// Задача 2: Напишите программу, которая на вход принимает два числа и выдаёт, какое число большее, а какое меньшее.


Console.Clear();
Console.Write("Введите A: ");
int number1 = int.Parse(Console.ReadLine());
Console.Write("Введите B: ");
int number2 = int.Parse(Console.ReadLine());

if (number1 > number2)
{
    Console.WriteLine("Число А больше");
}
else
{
    Console.WriteLine("Число В больше");
}