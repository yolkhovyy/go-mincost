# Min Cost Path

This project contains a Go implementation to calculate the minimum cost path in a 2D grid.

## Description

The `minCost` function calculates the minimum cost to reach a specific cell in a 2D grid from the top-left corner. 

### Parameters:
- `cost [][]int`: A 2D slice representing the cost grid.
- `m int`: The row index of the target cell.
- `n int`: The column index of the target cell.

### Function Log# Min Cost Path

This project contains a Go implementation to calculate the minimum cost path in a 2D grid.

## Description

The `minCost` function calculates the minimum cost to reach a specific cell in a 2D grid from the top-left corner. Here's a detailed description of how it works:

### Parameters:
- `cost [][]int`: A 2D slice representing the cost grid.
- `m int`: The row index of the target cell.
- `n int`: The column index of the target cell.

### Function Logic:
1. **Base Cases**:
   - If either `m` or `n` is negative, the function returns `math.MaxInt64`, indicating an invalid path.
   - If both `m` and `n` are zero, the function returns the cost of the top-left cell (`cost[0][0]`), as this is the starting point.

2. **Recursive Calculation**:
   - The function recursively calculates the minimum cost to reach the target cell `(m, n)` by considering three possible previous cells:
     - The cell directly above `(m-1, n)`.
     - The cell directly to the left `(m, n-1)`.
     - The cell diagonally above and to the left `(m-1, n-1)`.
   - It adds the cost of the current cell (`cost[m][n]`) to the minimum of the costs to reach these three previous cells.

### Return Value:
- The function returns the minimum cost to reach the cell `(m, n)` from the top-left corner `(0, 0)`.

