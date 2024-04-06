# Problem 3.2
I have second stack that maintains the smallest element that has been added to the main stack.

Push: If the element is smaller or equal to the element at the top of the secondary stack, also push it to the secondary stack.

Pop: If the element is equal to the number at the top of the secondary stack, also pop from the secondary stack. Note the element can never be smaller than what is at the top of the secondary stack, which is what allows us to save time.

Min: Return the element at the top of the secondary stack
