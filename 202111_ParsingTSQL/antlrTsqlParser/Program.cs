namespace TSqlParserTest;

using System;
using System.Linq;
using Antlr4.Runtime;
using Antlr4.Runtime.Tree;

public static class Program
{
    public static void Main(string[] args)
    {
        const string sqlFile = "t.sql";

        using (TextReader str = new StreamReader(sqlFile))
        {
            var lexer = new TSqlLexer(new AntlrInputStream(str));
            var tokenStream = new CommonTokenStream(lexer);
            var parser = new TSqlParser(tokenStream);
            var tree = parser.tsql_file();

            var ml = new MyParserListener();
            ParseTreeWalker.Default.Walk(ml, tree);
        }
    }

    public class MyParserListener : TSqlParserBaseListener
    {
        public override void ExitSelect_statement(TSqlParser.Select_statementContext context)
        {
            var expressions = context.query_expression();
            var querySpecs = expressions.query_specification();
            var sourceTable = querySpecs.table_sources().table_source().First().GetText();
            Console.WriteLine(sourceTable);
        }

        public override void ExitCase_expression(TSqlParser.Case_expressionContext context)
        {
            var elseMsg = context.elseExpr.GetText();
            Console.WriteLine($"Case ELSE expression: {elseMsg}");
        }

        public override void ExitDrop_table(TSqlParser.Drop_tableContext context)
        {
            var idNames = context.table_name().id_().Select(e => e.GetText());
            Console.WriteLine($"You're trying to delete {string.Join(".", idNames)}");
        }
    }
}
