using System;

class MainClass {
  static int Sum(int start, int end) {
    if (start > end) {
      return 0;
    }
    return start + Sum(start + 1, end);
  }

  public static void Main (string[] args) {
    int m = 1;
    int n = 10;
    int result = Sum(m, n);
    Console.WriteLine("Сумма натуральных элементов от {0} до {1}: {2}", m, n, result);
  }
}
