using Moq;

namespace Bank.Tests;

public class AccountServiceTests
{
    private readonly Mock<IClock> _mockClock;
    private readonly Mock<ITransactionRepository> _mockTransactionRepository;
    private readonly Mock<IReportGenerator> _mockReportGenerator;
    private readonly AccountService _subject;

    public AccountServiceTests()
    {
        _mockClock = new Mock<IClock>();
        _mockTransactionRepository = new Mock<ITransactionRepository>();
        _mockReportGenerator = new Mock<IReportGenerator>();

        _subject = new AccountService(
            _mockClock.Object,
            _mockTransactionRepository.Object,
            _mockReportGenerator.Object
        );
    }

    [Fact(Skip = "Skip until fully functional")]
    public void Should_Record_Transactions_And_Print_Statement_Correctly()
    {
        var expectedResult =
            @"
        Date        || Amount || Balance
        14/01/2012  || -500   || 2500
        13/01/2012  || 2000   || 3000
        10/01/2012  || 1000   || 1000";

        _subject.Deposit(1000);
        _subject.Deposit(2000);
        _subject.Withdraw(500);
        _subject.PrintStatement();
    }

    [Fact]
    public void Deposit_Should_Insert_Amount_To_Repository_With_Current_Date()
    {
        _mockClock.Setup(x => x.Now()).Returns(new DateTime(2017, 2, 3));
        _subject.Deposit(100);

        _mockTransactionRepository.Verify(x => x.Add(100, new DateTime(2017, 2, 3)), Times.Once);
    }

    [Fact]
    public void Withdrawn_Should_Insert_Negative_Amount_To_Repository_With_Current_Date()
    {
        _mockClock.Setup(x => x.Now()).Returns(new DateTime(2017, 2, 3));
        _subject.Withdraw(100);

        _mockTransactionRepository.Verify(x => x.Add(-100, new DateTime(2017, 2, 3)), Times.Once);
    }

    [Fact]
    public void PrintStatemt_Should_Pass_All_Transations_To_ReportGenerator()
    {
        var transactions = new List<Transaction> { new Transaction(100, new DateTime(2017, 2, 1)) };
        _mockTransactionRepository.Setup(x => x.GetAll()).Returns(transactions);

        _subject.PrintStatement();

        _mockReportGenerator.Verify(x => x.Create(transactions), Times.Once);
    }

    [Fact]
    public void PrintStatemt_Should_Pass_ReportGenerator_Result_To_StatementPrinter()
    {
        var transactions = new List<Transaction> { new Transaction(100, new DateTime(2017, 2, 1)) };
        _mockTransactionRepository.Setup(x => x.GetAll()).Returns(transactions);

        _subject.PrintStatement();

        _mockReportGenerator.Verify(x => x.Create(transactions), Times.Once);
    }
}
