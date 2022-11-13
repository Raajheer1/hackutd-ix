export default function getStockPrice(stocks) {
  for (var i = 0; i < tickers.length; i++) {
    if (tickers[stocks[i].ticker] != undefined)
      stocks[i].price = tickers[stocks[i].ticker];
    else stocks[i].price = 0;
  }
}

const tickers = {
  TSLA: 195.97,
  TQQQ: 22.67,
  AMZN: 100.79,
  AMD: 72.37,
  AAPL: 149.7,
  SPDR: 398.51,
  NVDA: 163.27,
  META: 113.02,
  MSFT: 247.11,
  SHOP: 39.44,
  QQQ: 287.96,
  GOOGL: 96.41,
  INTC: 30.43,
};
