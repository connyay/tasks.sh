load("yfinance", "load_ticker_data")
load("varz", "set_varz", "get_varz")
load("convert", "to_string_map")
load("database", "db_query", "db_exec", "db_migrate")


_in_format = "2006-01-02 15:04:05 -0700 MST"
_out_format = "2006-01-02 15:04:05 MST"
_location = "America/New_York"

def ny_time(timestamp):
    return time \
        .parse_time(timestamp.String(), _in_format) \
        .in_location(_location) \
        .format(_out_format)

def main(args):
    db_migrate([
'''
CREATE TABLE IF NOT EXISTS watch_list (
    ticker TEXT,
    open TEXT,
    close TEXT,
    timestamp TIMESTAMP,
    UNIQUE(ticker,open,close,timestamp)
)
'''
    ])
    tickers = args.get("tickers", "AAPL")
    for ticker in tickers.split(','):
        ticker = ticker.upper()
        data = load_ticker_data(ticker)
        changed = False
        for item in data.Quotes:
            if str(item.Open).endswith(".0"):
                printf("Watching %s for open=%v at %s\n", ticker, item.Open, ny_time(item.OpensAt))
                db_exec('INSERT INTO watch_list("ticker", "open", "close", "timestamp") VALUES (?, ?, ?, ?)', ticker, item.Open, item.Close, ny_time(item.OpensAt))
                changed = True
            if str(item.Close).endswith(".0"):
                printf("Watching %s for close=%v at %s\n", ticker, item.Close, ny_time(item.OpensAt))
                db_exec('INSERT INTO watch_list("ticker", "open", "close", "timestamp") VALUES (?, ?, ?, ?)', ticker, item.Open, item.Close, ny_time(item.OpensAt))
                changed = True
        printf("\n")
    double_zeros = {}
    for row in db_query('SELECT ticker, count(*) as count FROM watch_list GROUP BY ticker'):
        double_zeros[str(row.get("ticker"))] = row.get("count")
    set_varz("double_zeros", to_string_map(double_zeros))
