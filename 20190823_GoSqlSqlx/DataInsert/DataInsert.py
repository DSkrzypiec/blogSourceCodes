import psycopg2

conn = psycopg2.connect(user="postgres",
                      password="**********",
                      host="127.0.0.1",
                      port="5432",
                      database="Random")

print("Start...")
curs = conn.cursor()
insQuery = """INSERT INTO dbo.forbench(x1, x2, x3, x4, x5, x6, x7, x8, x9, x10) 
              VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""

for e in range(0, 10000000):
    x1 = e
    x2 = (e + 10) / 1312.5
    x3 = "CRAP"
    x4 = e + 142
    x5 = x2 * 1.123
    x6 = "DSkrzypiec"
    x7 = x4 - 321
    x8 = x5 * 0.98712
    x9 = "LOOOONGER VARCHAR"
    x10 = e - 11
    r = (x1, x2, x3, x4, x5, x6, x7, x8, x9, x10)
    curs.execute(insQuery, r)
    if e % 1281 == 0:
        print(e)

conn.commit()
print("Done!")
conn.close()
