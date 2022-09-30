# CourseraGoLangPractice
There are some my best works from the time i was passing Golang course with Coursera


# Links to my certificates 

**Getting started with go:**
https://coursera.org/share/e6d3307fb3b4e8018aac9c14f6780f4b


[https://coursera.org/share/a3c6c8cf5a9315bd6d4b32d3fd9666d1](Functions, Methods, and Interfaces in Go)

# Concurrency In Go


1. **chopsticks.go** :
   Implement the dining philosopher’s problem with the following constraints/modifications.

   There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

   Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

   The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

   In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

   The host allows no more than 2 philosophers to eat concurrently.

   Each philosopher is numbered, 1 through 5.

   When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

   When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

2. **numberSort.go** :
   Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
   The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
3. **raceCond.go** :
   Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.


# Functions, Methods and Interfaces in Go 


1. **bubbleSort.go** :
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
2. **distance.go** :
   Write a program which first prompts the user
   to enter values for acceleration, initial velocity, and initial displacement.
   Then the program should prompt the user to enter a value for time and the
   program should compute the displacement after the entered time.

   You will need to define and use a function
   called GenDisplaceFn() which takes three float64
   arguments, acceleration a, initial velocity vo, and initial
   displacement so. GenDisplaceFn()
   should return a function which computes displacement as a function of time,
   assuming the given values acceleration, initial velocity, and initial
   displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
   float64 argument which is the displacement travelled after time t.
3. **animal.go** :
   Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow,bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.
   Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

   You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.