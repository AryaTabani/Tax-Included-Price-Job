# Price Calculator CLI with a Modular I/O Architecture

A simple yet powerful command-line (CLI) tool written in **Go**, designed to calculate tax-inclusive prices. The architecture of this project is based on the principles of **Clean Architecture** and **Separation of Concerns**.

## ✨ Features & Key Concepts

-   **Interface-Based Design**: The core of the project's architecture is the use of an `IOManager` interface. This design allows the main logic to be completely decoupled from the input and output methods.
-   **Multiple I/O Implementations**:
    -   **Command Line (CMD)**: The application can read prices directly from the user's input in the terminal and print the result.
    -   **File I/O**: The application can read a list of prices from a `.txt` file and save the output to a new `JSON` file.
-   **High Extensibility**: Thanks to its modular design, adding new I/O methods (like reading from an API or a database) is straightforward and does not require any changes to the core application logic.
-   **Decoupled Processing Logic**: The price calculation module is completely independent of how data is received or displayed.

## 🛠️ Tech Stack

-   **Language**: Go (using only standard libraries)
-   **File Handling**: `os`, `bufio`
-   **JSON Handling**: `encoding/json`

## 🏛️ Architecture

This project utilizes an `IOManager` interface:
```go
type IOManager interface {
    ReadLines() ([]string, error)
    WriteResult(data interface{}) error
}
```
Two structs, `FileManager` and `CMDManager`, implement this interface. The core application logic (`TaxIncludedPriceJob`) interacts only with this interface, having no direct dependency on files or the command line.

## 🚀 How to Use

1.  Clone the project.
2.  Open the `main.go` file.
3.  Choose which I/O manager to use by uncommenting the desired line:

    ```go
    func main() {
        taxRates := []float64{0, 0.07, 0.1, 0.15}

        for _, taxRate := range taxRates {
            // To use the command line for I/O
            io := cmdmanager.New()

            // To use files for I/O (comment out the line above)
            // io := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

            priceJob := prices.NewTaxIncludedPriceJob(io, taxRate)
            priceJob.Process()
        }
    }
    ```
4.  Run the application:
    ```bash
    go run main.go
    ```
