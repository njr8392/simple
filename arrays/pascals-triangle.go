package main

/*Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:


 
        tri[i]=make([]int,i+1)      

LEETCODE 118


func generate(numRows int) [][]int {
	tri := make([][]int, numRows)

LEETCODE 118

*/

func generate(numRows int) [][]int {
	for i := range tri {

      tri:=make([][]int,numRows) 

    for i:= range tri{

        tri[i]=make([]int,i+1)      

    }
LEETCODE 118 LEETCODE 1 8     if j==0 ||       }else{
                tri[i][j]=	i-1][j]  
            }
        }
    }
    return tri
}
