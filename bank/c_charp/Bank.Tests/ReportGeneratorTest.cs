using FluentAssertions;

namespace Bank.Tests;

public class ReportGeneratorTest
{
    public ReportGenerator _subject;

    public ReportGeneratorTest()
    {
        _subject = new ReportGenerator();
    }

    [Fact]
    public void Format_Should_Return_Formtted_String_When_Transactions_Are_Provided()
    {
        var header = "Date||Amount||Balance";
        var line1 = "11/05/2012||1000||1000";
        var line2 = "10/05/2012||-100||0";
        var line3 = "09/05/2012||100||100";

        var transactions = new List<Transaction>()
        {
            new(100, new DateTime(2012, 5, 9)),
            new(-100, new DateTime(2012, 5, 10)),
            new(1000, new DateTime(2012, 5, 11))
        };

        var report = _subject.Create(transactions);

        report[0].Should().Be(header);
        report[1].Should().Be(line1);
        report[2].Should().Be(line2);
        report[3].Should().Be(line3);
    }
}
