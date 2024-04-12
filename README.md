[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/26dp6wek)

# Asynchronicity

Implement a function that takes an array and a key to search for and counts the
number of times key matches an element in the array (the count matches function
we talked about in lectures). Your implementation must count the number of
matches asynchronously, but does not need to do so in parallel. What type of
asynchronous execution you choose is up to you.

I have not provided a template; depending on how you choose to implement the
function, it will have a different signature.

I have also not provided any test code, but you can base yours on test code from
other exercises. Your tests must check the correctness of the result of running
the function and run automatically when you commit through a GitHub action.

The [async library](https://caolan.github.io/async/v3/) may be helpful with
this.

**Implimentation**

I will be using the Go programming language as I have been doing alot of work in that language recently and want to get better with channels. So I am using that along with Go's builtin testing xlibrary.

## Runtime Analysis

What is the time complexity of your implementation (worst-case $\Theta$)? Add
your answer, including your reasoning, to this markdown file.

**Analysis**

We start by iterating over the number of threads $t$ to generate each chunk index to pass to each thread. Then foreach thread we will iterate through the chunksize. This means a total of $n$ iterations occur during the checking phase. Whether we have 1 or 100 threads, we will always iterate through each element once.

Since this can be done in parallel for atleast as many cores as the user has on their machine, I would say it's $\Theta(n / t)$
runtime complexity in the best-case and $\Theta(n/1) \therefore \Theta(n)$ in the worst-case when the user has only one core to utilize.

However, since the number of cores or specs of a machine dont matter when analizing the runtime complexity of any algorithm, I say the average/worst-case complexity is always $\Theta(n)$.

## Resources

I had a hard time grasping channels at first. Here is a good website which has a series of diferent pages you can view which contain some code and a brief explaination. I also watched a few youtube videos on the topic.
https://gobyexample.com/closing-channels

For number generation I utilized this code snippet to get me a range of random integers.
https://stackoverflow.com/questions/23577091/generating-random-numbers-over-a-range-in-go

I used this testing code snippet for help with generating tests
https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
