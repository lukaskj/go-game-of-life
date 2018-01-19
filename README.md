# go-game-of-life

## Implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go

### Rules
> 1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.<br/>
> 1. Any live cell with two or three live neighbours lives on to the next generation.<br/>
> 1. Any live cell with more than three live neighbours dies, as if by overpopulation.<br/>
> 1. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.<br/>

## Usage
```
go run gameOfLife.go [nGenerations [width height [seed]]]
```
## Examples
* Generates the default 100 generations of 10x10 matrix with random seed:
```bash
go run gameOfLife.go 
```
* 500 generations of 10x10 (default size) matrix with random seed:
```
go run gameOfLife.go 500
```
* 500 generations of 20x10 matrix with random seed:
```
go run gameOfLife.go 500 20 10
```

* 500 generations of 20x10 matrix with fixed seed (seed is _int64_):
```
go run gameOfLife.go 500 20 10 65984445125
```

> _Drawing in console isn't the best but this is just for practice_