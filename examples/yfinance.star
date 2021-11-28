load("yfinance", "load_ticker_data")

def run():
    data = load_ticker_data(parameters.get("ticker", "GME"))
    if parameters.get("dump"):
        dump(data)
    for item in data.Quotes:
        printf("OpensAt %v Open %v Close %v\n", item.OpensAt, item.Open, item.Close)

run()