# Sort String CSV

This function takes unordered data from a CSV in a string and returns the data sorted by name in ascending order.

## Resources
- Golang (version 1.12)
- Test and native packages

## How does it work?
The `SortCsvColumns()` function receives a `string` with the CSV data. For this purpose, the data has already been extracted previously.
First, the lines are separated with the `Split()` function. Then, each value of the lines is separated and everything is saved in a string array variable.
With all data stored separately, a `for` is made to merge all of a person's data into struct. Each line has their data grouped.
With the struct, it is finally done sorting by name in ascending order. After that just merge all the data again making a range of each line inside the struct.

The `Len()`, `Less()` and `Swap()` functions serve to sort the struct data.

## Important
This code was written for a specific purpose. The `name` was defined as a header, serving as the basis for sort. It wasn't specified what each value is, so I treated them as first and second values.