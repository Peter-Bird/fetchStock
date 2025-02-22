# Stock Data Downloader

This Go application provides a simple graphical user interface (GUI) to download historical stock data from Yahoo Finance, save the data as a CSV file, and display a plot of the stock's closing price. The program is built using the following libraries:

- **Fyne**: For building the GUI.
- **piquette/finance-go**: To fetch stock data from Yahoo Finance.
- **Gonum/plot**: To generate a line plot of the stock's closing prices.
- **shopspring/decimal**: To handle decimal arithmetic when converting financial data.

## Features

- **Ticker Selection**: Choose from a list of major tech tickers (e.g., AAPL, MSFT, GOOGL, etc.).
- **Data Download**: Fetches historical daily data for the selected stock.
- **CSV Export**: Saves the downloaded data into a CSV file.
- **Plotting**: Displays a plot of the closing prices over time.
- **GUI Interface**: A simple window with a dropdown to select the ticker, a button to download data, and an image widget to display the plot.

## Prerequisites

- **Go**: Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
- **Git** (optional): To clone the repository.

## Installation

1. **Clone the Repository** (if applicable):

   ```bash
   git clone https://github.com/yourusername/stock-data-downloader.git
   cd stock-data-downloader
   ```

2. **Install Dependencies**:

   This project uses Go modules. To install the required packages, run:

   ```bash
   go mod tidy
   ```

   This command will download and install the following dependencies:
   - [fyne.io/fyne/v2](https://github.com/fyne-io/fyne)
   - [github.com/piquette/finance-go](https://github.com/piquette/finance-go)
   - [gonum.org/v1/plot](https://gonum.org/v1/plot/)
   - [github.com/shopspring/decimal](https://github.com/shopspring/decimal)

## Usage

To run the application, execute:

```bash
go run .
```

The application window will appear with a dropdown menu to select a tech stock. Click the **Download Data** button to fetch the data, save it as a CSV file (e.g., `AAPL_stock_data.csv`), and update the plot displayed in the window.

## Troubleshooting

- **Data Fetch Errors**:  
  If you receive an error message like `Failed to fetch data: code: remote-error: error response received from upstream api`, verify that:
  - The ticker symbol is valid.
  - There are no network issues or API rate limiting problems.
  - The underlying Yahoo Finance API is working correctly.

- **Fyne Image Errors**:  
  If you see errors related to image loading (e.g., "Failed to load image: image: unknown format"), ensure that the initial image resource is a valid PNG. The provided code initializes a blank PNG image for display until the plot is updated.

## Customization

- **Tickers**:  
  You can update the list of tech tickers by modifying the `techTickers` slice in the source code.

- **Plot Appearance**:  
  To customize the plot's appearance (e.g., title, axis labels, dimensions), modify the `plotStockData` function in the source code.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fyne](https://fyne.io/) for providing a modern GUI toolkit.
- [piquette/finance-go](https://github.com/piquette/finance-go) for easy access to Yahoo Finance data.
- [Gonum/plot](https://github.com/gonum/plot) for creating beautiful plots.

---
