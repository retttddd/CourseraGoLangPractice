# CourseraGoLangPractice
There are some my best works from the time i was passing Golang course with Coursera


# Concurrency In Go


1. chopsticks.go :
   Implement the dining philosopher’s problem with the following constraints/modifications.

   There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

   Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

   The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

   In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

   The host allows no more than 2 philosophers to eat concurrently.

   Each philosopher is numbered, 1 through 5.

   When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

   When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

2. numberSort.go :
   Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
   The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
3. raceCond.go :
   Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.


# Functions, Methods and Interfaces in Go 


1. bubbleSort.go :
   Write a Bubble Sort program in Go. The program
   should prompt the user to type in a sequence of up to 10 integers. The program
   should print the integers out on one line, in sorted order, from least to
   greatest. Use your favorite search tool to find a description of how the bubble
   sort algorithm works.

   As part of this program, you should write a
   function called BubbleSort() which
   takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
   order.

   A recurring operation in the bubble sort algorithm is
   the Swap operation which swaps the position of two adjacent elements in the
   slice. You should write a Swap() function which performs this operation. Your Swap()
   unction should take two arguments, a slice of integers and an index value i which
   indicates a position in the slice. The Swap() function should return nothing, but it should swap
   the contents of the slice in position i with the contents in position i+1.
2. 



