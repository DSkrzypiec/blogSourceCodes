using System.Collections.Generic;

public class Program {
    public void Main() {
    }
}

public static class Generator
{
    public static IEnumerable<int> GenerateNumbers()
    {
        var secretLocalNumber = 42;
        yield return secretLocalNumber + 13;

        for (var tmpId = 0; tmpId < 3; tmpId++)
            yield return tmpId + 10;

        yield return SomeStaticMethod(-99);
    }

    public static int SomeStaticMethod(int x) => x - 1;
}
