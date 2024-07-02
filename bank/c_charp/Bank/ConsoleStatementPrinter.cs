namespace Bank;

public class ConsoleStatementPrinter : IStatementPrinter
{
    public void Print(List<string> statement)
    {
        foreach (var line in statement)
        {
            Console.WriteLine(line);
        }
    }
}
