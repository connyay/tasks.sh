load("yfinance", "load_ticker_data")

_in_format = "2006-01-02 15:04:05 -0700 MST"
_out_format = "2006-01-02 15:04:05 MST"
_location = "America/New_York"

def ny_time(timestamp):
    return time \
        .parse_time(timestamp.String(), _in_format) \
        .in_location(_location) \
        .format(_out_format)

def main(args):
    ticker = args.get("ticker", "AAPL")
    data = load_ticker_data(ticker)
    if args.get("dump"):
        dump(data)
    for item in data.Quotes:
        printf("[%s]: %v (open=%v, close=%v)\n", ticker, ny_time(item.OpensAt), item.Open, item.Close)
