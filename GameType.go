package main

import (
   "time"
   "math/rand"
)

type Matrix [][]int

type GameOfLife struct {
   Rows int
   Cols int
   Generations []Matrix
   CurrentGeneration int
   Seed int64
}

func (g *GameOfLife) Init(cols, rows, nGens int, seed int64) {
   g.Cols = cols
   g.Rows = rows
   
   g.Generations = make([]Matrix, nGens)
   for i := 0; i < nGens; i++ {
      g.Generations[i] = make(Matrix, g.Cols)
      for x := range g.Generations[i] {
         g.Generations[i][x] = make([]int, g.Rows)
      }
   }

   g.Seed = seed
}

func (g *GameOfLife) CreateGeneration(n int) {
   nextGen := g.Generations[n] // Points current generation
   if n == 0 { // create first generation randomly or by seed if informed
      if g.Seed == 0 {
         g.Seed = time.Now().UnixNano()
      }
      r := rand.New(rand.NewSource(g.Seed))
      var x, y int
      for x = range nextGen {
         for y = range nextGen[x] {
            nextGen[x][y] = r.Intn(2)
         }
      }
   } else {
      // Copy last gen to current
      currentGen := g.Generations[n - 1]
      for i := range currentGen {
         for j := range currentGen[i] {
            var currentState int = currentGen[i][j]
            var aliveNeighbors int = g.countAliveNeighbors(&currentGen, i, j)
            
            // Apply rules
            if currentState == 0 && aliveNeighbors == 3 {
               nextGen[i][j] = 1
            } else if currentState == 1 && (aliveNeighbors < 2 || aliveNeighbors > 3) {
               nextGen[i][j] = 0
            } else {
               nextGen[i][j] = currentState
            }
         }
      }
   }
}

// Count alive neighbors with matrix wrap
func (g *GameOfLife) countAliveNeighbors(mP *Matrix, x, y int) int {
   m := *mP
   sum := 0
   for i := -1; i <= 1; i++ {
      for j := -1; j <= 1; j++ {
         col := (x + i + g.Cols) % g.Cols
         row := (y + j + g.Rows) % g.Rows
         sum += m[col][row]
      }
   }
   sum -= m[x][y]

   return sum
}

func (g *GameOfLife) GenerateAll() {
   for i := range g.Generations {
      g.CreateGeneration(i)
   }
}