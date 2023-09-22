package main

import (
	"fmt"
)
func main() {
N:=0
M:=0
fmt.Println("Enter the number of inputs")
fmt.Scanln(&N)
fmt.Println("Enter the inputs")
arr:=make([]int, N)
for i:=0; i<N; i++ {
fmt.Scanln(&arr[i])
}
fmt.Println("Enter value of m: ")
fmt.Scanln(&M)
arr1:=make([]int, M)
for i:=0; i<M; i++ {
arr1[i]=arr[i]
}
for j:=M; j<N; j++ {
big:=0
for k:=1; k<M; k++ {
if arr1[k]>arr[big] {
big=k
}
}
if arr[j]<arr1[big] {
arr1[big]=arr[j]
}
}
for l:=0; l<M-1; l++ {
for m:=l+1; m<M; m++ {
if arr1[l] > arr1[m] {
arr1[l], arr1[m]=arr1[m], arr1[l]
}
}
}
fmt.Println("smallest", M, "numbers are", arr1)
}
