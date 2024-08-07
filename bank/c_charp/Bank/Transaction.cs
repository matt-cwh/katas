namespace Bank;

public class Transaction
{
    public int Amount { get; private set; }
    public DateTime Date { get; private set; }

    public Transaction(int amount, DateTime date)
    {
        Amount = amount;
        Date = date;
    }
}
