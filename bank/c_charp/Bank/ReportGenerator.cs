using System.Text;

namespace Bank;

public class ReportGenerator : IReportGenerator
{
    public List<string> Create(List<Transaction> transactions)
    {
        var orderedTransactions = transactions.OrderByDescending(t => t.Date).ToList();

        var reportStrs = new List<string>();
        var balance = 0;

        foreach (var transaction in transactions)
        {
            Console.WriteLine(transaction.Amount);
            balance += transaction.Amount;
            reportStrs.Add(
                $"{transaction.Date:dd/MM/yyyy}||{transaction.Amount}||{balance}"
            );
        }

        reportStrs.Add("Date||Amount||Balance");
        reportStrs.Reverse();

        return reportStrs;
    }
}
