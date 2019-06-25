package main

import "fmt"

func main() {
   var n [100]int /* n 是一个长度为 10 的数组 */
   var i int
   var sum int 

   /* 为数组 n 初始化元素 */         
   for i = 0; i < 100; i++ {
      n[i] = i+1
      fmt.Printf("Element[%d] = %d\n", i, n[i] )
      sum=sum+n[i]
   }

 
 
   fmt.Printf("sum = %d\n", sum )
 
}