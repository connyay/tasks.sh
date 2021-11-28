load("yfinance", "load_ticker_data")

def run():
    data = load_ticker_data("GME")
    for item in data.Quotes:
        printf("OpensAt %v Open %v Close %v\n", item.OpensAt, item.Open, item.Close)

run()