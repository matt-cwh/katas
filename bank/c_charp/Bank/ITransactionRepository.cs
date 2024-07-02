namespace Bank;

public interface ITransactionRepository
{
    public void Add(int amount, DateTime date);
    public List<Transaction> GetAll();
}
