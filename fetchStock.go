package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/piquette/finance-go/chart"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// List of major technology sector tickers
var techTickers = []string{
	"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA", "NVDA", "META",
	"ORCL", "IBM", "ADBE", "INTC", "CSCO", "AMD", "PYPL", "CRM",
}

// StockData holds the date and closing price.
type StockData struct {
	Date  time.Time
	Close float64
}

// fetchStockData downloads historical stock data for the given symbol.
func fetchStockData(symbol string) ([]StockData, error) {
	params := &chart.Params{
		Symbol:   symbol,
		Interval: "1d",
		// Removed unsupported "Range" field. By default, this will fetch available data.
	}
	iter := chart.Get(params)
	var data []StockData
	for iter.Next() {
		bar := iter.Bar()
		t := time.Unix(int64(bar.Timestamp), 0)
		// Convert decimal.Decimal to float64.
		closeVal, _ := bar.Close.Float64()
		data = append(data, StockData{
			Date:  t,
			Close: closeVal,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no data found for %s", symbol)
	}
	return data, nil
}

// saveToCSV writes the stock data to a CSV file.
func saveToCSV(symbol string, data []StockData) error {
	filename := fmt.Sprintf("%s_stock_data.csv", symbol)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header.
	writer.Write([]string{"Date", "Close"})
	for _, d := range data {
		writer.Write([]string{
			d.Date.Format("2006-01-02"),
			fmt.Sprintf("%f", d.Close),
		})
	}
	return nil
}

// plotStockData creates a line plot of the stock closing prices and returns it as an image.
func plotStockData(symbol string, data []StockData) (image.Image, error) {
	// plot.New() now returns a single value.
	p := plot.New()
	p.Title.Text = fmt.Sprintf("%s Stock Price", symbol)
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Price (USD)"

	pts := make(plotter.XYs, len(data))
	for i, d := range data {
		pts[i].X = float64(d.Date.Unix())
		pts[i].Y = d.Close
	}

	if err := plotutil.AddLinePoints(p, fmt.Sprintf("%s Closing Price", symbol), pts); err != nil {
		return nil, err
	}

	// Format X axis to display dates.
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	// Save plot to a temporary PNG file.
	imgFile := fmt.Sprintf("%s_stock_plot.png", symbol)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, imgFile); err != nil {
		return nil, err
	}

	// Open and decode the image file.
	f, err := os.Open(imgFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	// Remove the temporary file.
	os.Remove(imgFile)
	return img, nil
}

// imageToPNG converts an image.Image to a PNG byte slice.
func imageToPNG(img image.Image) []byte {
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return buf.Bytes()
}

func createBlankPNG() []byte {
	// Create a small blank image.
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// Fill with white.
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.White)
		}
	}
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return buf.Bytes()
}

func main() {
	// Create a new Fyne application.
	a := app.New()
	w := a.NewWindow("Stock Data Downloader")

	// Create a dropdown (Select widget) for tech tickers.
	tickerSelect := widget.NewSelect(techTickers, nil)
	tickerSelect.SetSelected("AAPL")

	// Create an initial valid (blank) PNG resource.
	blankPNG := createBlankPNG()
	plotResource := fyne.NewStaticResource("plot.png", blankPNG)
	imageWidget := widget.NewIcon(plotResource)

	// Button to fetch data, save CSV, and update the plot.
	btn := widget.NewButton("Download Data", func() {
		symbol := tickerSelect.Selected
		if symbol == "" {
			dialog.ShowError(fmt.Errorf("please select or enter a stock ticker"), w)
			return
		}

		// Fetch stock data.
		data, err := fetchStockData(symbol)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to fetch data: %v", err), w)
			return
		}

		// Save data to CSV.
		if err := saveToCSV(symbol, data); err != nil {
			dialog.ShowError(fmt.Errorf("failed to save CSV: %v", err), w)
			return
		}

		// Plot the stock data.
		img, err := plotStockData(symbol, data)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to plot data: %v", err), w)
			return
		}

		// Update the image widget with the new plot.
		imageWidget.SetResource(fyne.NewStaticResource("plot.png", imageToPNG(img)))
		imageWidget.Refresh()

		dialog.ShowInformation("Success", fmt.Sprintf("Data saved and plot updated for %s", symbol), w)
	})

	// Layout the UI.
	content := container.NewVBox(
		widget.NewLabel("Select a Tech Stock:"),
		tickerSelect,
		btn,
		imageWidget,
	)
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
