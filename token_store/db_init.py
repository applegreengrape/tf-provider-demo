import sqlite3

db = sqlite3.connect('demo.db')

cursor = db.cursor()
cursor.execute('''CREATE TABLE demo(
    user CHAR(100) UNIQUE,
    id NVARCHAR(100),
    stats CHAR(100)
    );
               ''')
db.commit()
db.close