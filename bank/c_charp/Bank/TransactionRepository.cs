namespace Bank;

public class TransactionRepository : ITransactionRepository
{
    private readonly List<Transaction> _transactions;

    public TransactionRepository()
    {
        _transactions = new List<Transaction>();
    }

    public void Add(int amount, DateTime date)
    {
        _transactions.Add(new Transaction(amount, date));
    }

    public List<Transaction> GetAll()
    {
        return _transactions;
    }
}
