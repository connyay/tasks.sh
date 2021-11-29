load("database", "db_query", "db_exec", "db_migrate")

def main(args):
    db_migrate([
'''
CREATE TABLE IF NOT EXISTS runs (
    id INTEGER PRIMARY KEY,
    timestamp TIMESTAMP
)
'''
    ])
    db_exec('INSERT INTO runs("timestamp") VALUES(CURRENT_TIMESTAMP)')
    rows = db_query("select * from runs")
    for row in rows:
        printf("id=%v timestamp=%q\n", row.get("id"), row.get("timestamp"))