

# CSV to JSON Converter

This project is a simple Go module that converts CSV files to JSON format. The program reads a CSV file, processes its content, and outputs the result as a JSON file.

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.14+)

### Installing

1.  **Clone the repository:**
    
    sh
    
    Copy code
    
    ```sh
    git clone https://github.com/guilhermemcardoso/gotransform.git
    cd gotransform
    ``` 
    
2. **Build the project:**
        
    ```sh
    go build -o gotransform
    ``` 
    

## Usage

To use the CSV to JSON converter, run the following command:

sh

Copy code

`./gotransform <yourfile.csv>` 

This command will generate a JSON file with the same name as the input CSV file, but with a `.json` extension.

### Example

Suppose you have a CSV file named `example.csv`. To convert it to JSON, use:

```sh
./gotransform example.csv
``` 

The program will create a file named `example.json` containing the JSON representation of the CSV data.

## Code Overview

The main logic of the program is in the `main.go` file. The process is as follows:

1.  **Read the CSV file:**
    -   The CSV file is opened and read using the `encoding/csv` package.
2.  **Process the CSV content:**
    -   The first row is considered as headers.
    -   Each subsequent row is converted into a JSON object using the headers as keys.
3.  **Write to a JSON file:**
    -   The JSON data is written to a new file with a `.json` extension.

## Contributing

If you have suggestions for improving this project, feel free to submit a pull request or open an issue.