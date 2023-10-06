## Bubble Sort
Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly steps through the list to be sorted, compares each pair of adjacent items and swaps them if they are in the wrong order. The pass through the list is repeated until no swaps are needed, which indicates that the list is sorted. The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list. Although the algorithm is simple, it is too slow and impractical for most problems even when compared to insertion sort. Bubble sort can be practical if the input is in mostly sorted order with some out-of-order elements nearly in position.

Time complexity analysis:

|Worst Case|Average Case|Best Case|
|---|---|---|
|O(n<sup>2</sup>)|Θ(n<sup>2</sup>)|Ω(n)|

|In-place?|Stable?|
|---|---|
|Yes|Yes|

## Shell Sort
Shellsort, also known as Shell sort or Shell's method, is an in-place comparison sort. It can be seen as either a generalization of sorting by exchange (bubble sort) or sorting by insertion (insertion sort).The method starts by sorting pairs of elements far apart from each other, then progressively reducing the gap between elements to be compared. Starting with far apart elements, it can move some out-of-place elements into position faster than a simple nearest neighbor exchange. Donald Shell published the first version of this sort in 1959.The running time of Shellsort is heavily dependent on the gap sequence it uses. For many practical variants, determining their time complexity remains an open problem.

Time complexity analysis:

|Worst Case|Average Case|Best Case|
|---|---|---|
|O(n<sup>2</sup>)|Depends on Implementation|Ω(n log(n))|

|In-place?|Stable?|
|---|---|
|Yes|No|

## Gnome sort
Gnome sort is a sorting algorithm which is similar to insertion sort, except that moving an element to its proper place is accomplished by a series of swaps, similar to a bubble sort. It is conceptually simple, requiring no nested loops. The average, or expected, running time is O(n2) but tends towards O(n) if the list is initially almost sorted.

The algorithm always finds the first place where two adjacent elements are in the wrong order and swaps them. It takes advantage of the fact that performing a swap can introduce a new out-of-order adjacent pair next to the previously swapped elements. It does not assume that elements forward of the current position are sorted, so it only needs to check the position directly previous to the swapped elements.

Time complexity analysis:

|Worst Case|Average Case|Best Case|
|---|---|---|
|O(n<sup>2</sup>)|O(n<sup>2</sup>)|Ω(n<sup>2</sup>)|

|In-place?|Stable?|
|---|---|
|Yes|Yes|

## Some ters

### Stable
A sorting algorithm is said to be stable if two objects with equal keys appear in the same order in sorted output as they appear in the input data set

### In Place
An in-place algorithm is an algorithm that does not need an extra space and produces an output in the same memory that contains the data by transforming the input 'in-place'.