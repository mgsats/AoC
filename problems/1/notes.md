### Part 1:

#### Time Complexity: O(n)
#### Space Complexity: O(n)

---

### Part 2:

#### Time Complexity: O(n^2)
#### Space Complexity: O(n^2)

A better solution would involve first creating an array of the `n` integers, sorting it in `O(nlogn)` time, and then for each element `e` at index `i` of the array using the [two pointer approach](https://leetcode.com/articles/two-pointer-technique/) to search for two elements in sub array `arr[i+1:len(n)]` that sum to `2020-e`, giving `O(n^2)` and `O(1)` space and time complexity.


