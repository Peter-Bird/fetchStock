import tkinter as tk
from tkinter import ttk, messagebox
import yfinance as yf
import pandas as pd
import matplotlib.pyplot as plt
from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg

# List of major technology sector tickers
tech_tickers = [
    "AAPL", "MSFT", "GOOGL", "AMZN", "TSLA", "NVDA", "META",
    "ORCL", "IBM", "ADBE", "INTC", "CSCO", "AMD", "PYPL", "CRM"
]

def fetch_stock_data():
    """Fetch stock data from Yahoo Finance."""
    symbol = ticker_combobox.get().strip().upper()

    if not symbol:
        messagebox.showerror("Input Error", "Please select or enter a stock ticker.")
        return

    try:
        stock = yf.Ticker(symbol)
        stock_data = stock.history(period="max")  # Get full available data
        
        if stock_data.empty:
            messagebox.showerror("Error", f"No data found for {symbol}")
            return

        # Save data to CSV
        csv_filename = f"{symbol}_stock_data.csv"
        stock_data.to_csv(csv_filename)
        
        # Display the graph
        plot_stock_data(stock_data, symbol)
        
        messagebox.showinfo("Success", f"Data saved as {csv_filename}")

    except Exception as e:
        messagebox.showerror("Error", f"Failed to fetch data: {str(e)}")

def plot_stock_data(stock_data, symbol):
    """Plot stock data on a graph."""
    ax.clear()
    ax.plot(stock_data.index, stock_data["Close"], label=f"{symbol} Closing Price")
    ax.set_title(f"{symbol} Stock Price")
    ax.set_xlabel("Date")
    ax.set_ylabel("Price (USD)")
    ax.legend()
    ax.grid()
    canvas.draw()

# Create GUI window
root = tk.Tk()
root.title("Stock Data Downloader")

# Input field (Combobox) for stock ticker
tk.Label(root, text="Select a Tech Stock:").pack(pady=5)
ticker_combobox = ttk.Combobox(root, values=tech_tickers)
ticker_combobox.pack(pady=5)
ticker_combobox.set("AAPL")  # Default selection

# Button to fetch data
btn_fetch = tk.Button(root, text="Download Data", command=fetch_stock_data)
btn_fetch.pack(pady=10)

# Matplotlib Figure and Canvas
fig, ax = plt.subplots(figsize=(5, 3))
canvas = FigureCanvasTkAgg(fig, master=root)
canvas.get_tk_widget().pack()

# Run the GUI
root.mainloop()

