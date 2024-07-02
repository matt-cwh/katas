namespace Bank;

public class AccountService : IAccountService
{
    private readonly IClock _clock;
    private readonly ITransactionRepository _transactionRepository;

    private readonly IReportGenerator _reportGenerator;

    public AccountService(
        IClock clock,
        ITransactionRepository transactionRepository,
        IReportGenerator reportGenerator
    )
    {
        _clock = clock;
        _transactionRepository = transactionRepository;
        _reportGenerator = reportGenerator;
    }

    public void Deposit(int amount)
    {
        _transactionRepository.Add(amount, _clock.Now());
    }

    public void PrintStatement()
    {
        _reportGenerator.Create(_transactionRepository.GetAll());
    }

    public void Withdraw(int amount)
    {
        _transactionRepository.Add(-amount, _clock.Now());
    }
}
