namespace Bank;

public interface IReportGenerator
{
    List<string> Create(List<Transaction> transactions);
}
