
#!/usr/bin/python3
# coding=UTF-8


import psycopg2

db = psycopg2.connect(host='pgm-7gonknvn7y0e3x54194080.pg.rds.aliyuncs.com',
                     port="5432",
                     user='postgres',
                     password='bgsPR5i38ZTdb6hK',
                     database='pregod11')


cursor = db.cursor()

cursor.execute("SELECT VERSION()")

data = cursor.fetchone()

for i in range(30):
        sql = """update transfers set timestamp = (select timestamp from transactions where transfers.transaction_hash = transactions.hash)
                WHERE transaction_hash IN (
                        SELECT transaction_hash FROM transfers where
                        timestamp < '1970-01-01 00:00:00.000000 +00:00' limit 1000
                ) """
        cursor.execute(sql)
        print("ok")

db.close()