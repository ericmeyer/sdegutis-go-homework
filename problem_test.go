package sdegutishomework

import (
  "testing"
  . "gobdd"
)

func TestInitializer(t *testing.T) {}

func TestFib(t *testing.T) {
  defer PrintSpecReport()
  
  Describe("fibonacci", func() {
    It("has the basic values", func() {
      firstNums := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
      
      for i, num := range firstNums {
        Expect(fibonacci(i), ToEqual, num)
      }
    })
  })
}
