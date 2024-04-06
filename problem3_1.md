# Three in One: Describe how you could use a single array to implement three stacks.
I would have three sub-arrays of size k in a single size k*3 array.

I would keep track of the filled size of each subarray with 3 integers. The API for an operation would look something like `push(stack_idx, value)`, so it would be easy to keep track of the integers in a size 3 array and update it using the respective index.

On the push and pop operations I would increment and decrement the counts respectively

On each pop operation I would check if `current_size + 1` > k, and if so I would reallocate the entire array to double the size.
 
Potentially there are smarter ways to do this, for example one weakness is that if two of the stacks generally weren't used, we would be allocating three times more space than is really required. To solve this you could use a similar approach, execept when you reallocate you only double the size of the stack that exceeded its size. Using this approach you would also need to keep track of where the base of each stack is. This would be more complicated to write, and may not be worth it for small scale use cases.

Note that this question is not entirely arbitrary, and a real use case for this solution is if you were writing a memory allocator. The difference being that you would have an dynamic amount of 'data structures' in your memory space, meaning there is a lot more housekeeping, and you might need to worry about fragmentation.
