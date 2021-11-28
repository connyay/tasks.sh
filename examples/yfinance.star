load("yfinance", "load_ticker_data")

def run():
    ticker = parameters.get("ticker", "AAPL")
    data = load_ticker_data(ticker)
    if parameters.get("dump"):
        dump(data)
    for item in data.Quotes:
        printf("[%s]: %v (open=%v, close=%v)\n", ticker, item.OpensAt, item.Open, item.Close)

run()