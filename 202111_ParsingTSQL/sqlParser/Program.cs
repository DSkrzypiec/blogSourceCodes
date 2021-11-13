namespace SqlParser;

using System;
using Microsoft.SqlServer.TransactSql.ScriptDom;

public static class Program
{
    public static void Main(string[] args)
    {
        const string sqlFile = "t2.sql";

        using (TextReader str = new StreamReader(sqlFile))
        {
            var parser = new TSql150Parser(true, SqlEngineType.All);
            var tree = parser.Parse(str, out var errors);

            if (errors.Count > 0)
            {
                Console.WriteLine($"Errors in {sqlFile}:");
                foreach (var err in errors)
                {
                    Console.WriteLine(err.Message);
                }
            }

            Console.WriteLine($"Visiting tables and [DROP TABLE]s in {sqlFile}...");
            tree.Accept(new MyVisitor());

            /*
                var srcGen = new Sql150ScriptGenerator();
                srcGen.GenerateScript(tree, out var newSql);
                Console.WriteLine("Generated from parserd:");
                Console.WriteLine(newSql);
            */
        }
    }

    public class MyVisitor : TSqlFragmentVisitor
    {
        public override void ExplicitVisit(NamedTableReference table)
        {
            var schema = table.SchemaObject.SchemaIdentifier;
            var schemaName = schema?.Value ?? "";
            var tableName = table.SchemaObject.BaseIdentifier.Value;
            Console.WriteLine($"schema: {schemaName}, table: {tableName}");

            base.ExplicitVisit(table);
        }

        public override void ExplicitVisit(DropTableStatement stm)
        {
            var objects = stm.Objects;
            if (objects.Count > 0)
            {
                foreach (var obj in objects)
                {
                    Console.WriteLine($"You try to DELETE {obj.BaseIdentifier.Value} at line: {obj.StartLine} col: {obj.StartColumn}");
                }
            }

            base.ExplicitVisit(stm);
        }
    }
}


