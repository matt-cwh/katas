using Bank;
using FluentAssertions;

namespace Bank.Tests;

public class TransactionRepositoryTests
{
    private readonly TransactionRepository _subject;

    public TransactionRepositoryTests()
    {
        _subject = new TransactionRepository();
    }

    [Fact]
    public void Add_Should_Insert_New_Value_To_Transation_List()
    {
        _subject.Add(100, new DateTime(2021, 9, 5));
        var allTransations = _subject.GetAll();

        allTransations.Should().HaveCount(1);
        allTransations.First().Amount.Should().Be(100);
        allTransations.First().Date.Should().BeSameDateAs(new DateTime(2021, 9, 5));
    }
}
