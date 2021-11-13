SELECT
    x.Col1,
    y.Col2::varchar
FROM
    crap.TableName1 x
LEFT HASH JOIN
    dbo.Table2 y ON x.ID = y.ID2


DROP TABLE dbo.Crap

WITH x AS (
    SELECT TOP 10 * FROM test.Crap WHERE X = 10
)
SELECT
    s.A
FROM
    x
CROSS JOIN
    tmp.Scores s


DROP table  IF EXISTS dbo.Table1, /* comment */ tmp.Crap,    test.Op
