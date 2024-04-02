# Website Analytics

## Prerequisites

Before running the program, ensure you have Go installed on your system. You can download and install Go from the [official website](https://go.dev/dl/).

## How to Run

1.  Clone this repository to your local machine.
2.  Navigate to the project directory in your terminal.
3.  Run the following command to build and execute the program:

Linux/Mac:

    go run cmd/main.go

Win 10:

    go run cmd\main.go

## Benchmarking

To run the benchmark tests, execute the following command:

Linux/Mac:

    go test -bench=.

Win10:

    go test -bench=.

## Approach, Efficiency and Big O Notation

In this case i picked Hash Table Approach

The algorithm used in this project is based on the concept of iteration and boolean checks. It's a straightforward algorithm designed to iterate through each user's visit data and determine if they visited all products on a given day.

**The steps are**:

- Read both CSV files and construct hash tables where keys are user IDs and values are sets of product IDs visited by each user on each day.
- Iterate through the hash tables to find common users between the two days.
- For each common user, check if their set of product IDs on the second day contains any product ID not present on the first day.

The time complexity (Big O) of this algorithm depends on the size of the input data

The algorithm iterates over each user's visit data and checks if they visited all products.
Therefore, the overall time complexity is O(n×m), where n is the number of users and m is the average number of products visited per user.

## Time

~3-4 hours spent solving the problem
