### Dynamic Data Structures: Slices and their Operations

**A slice in Go is a flexible, lightweight view of an underlying array. While arrays in Go have a fixed size—making them impractical for dynamic data like a user-provided list of web request logs or a shopping cart—slices provide the abstraction of a dynamically resizing list.**

**Under the hood, a slice is a header containing three things: a pointer to an element in an array, the length of the slice (how many elements it currently holds), and the capacity of the slice (how many elements it can hold before the underlying array needs to be reallocated).**

### Creating and Initializing Slices

*You can create a slice using a slice literal, which looks like an array declaration without a fixed size, or by using the make function.
go*

## // Literal declaration: compiler creates an underlying array for you
    tasks := []string{"email client", "fix bugs", "deploy service"}

    // Using make: specify type, length, and (optionally) capacity
    // make([]Type, length, capacity)
    connections := make([]int, 0, 10)


## Manipulating Slice Contents

*Slices are modified using the built-in append function. Because append might trigger a reallocation of the underlying array, it returns a new slice header. You must always assign the result back to your variable.*

func main() {
    // Start with an empty slice
    var logs []string

    // Append items
    logs = append(logs, "INFO: Request started")
    logs = append(logs, "INFO: Database queried")

    // Append multiple items at once using the ... variadic operator
    moreLogs := []string{"INFO: Response sent", "INFO: Connection closed"}
    logs = append(logs, moreLogs...)
}

### Understanding Slicing and Memory

You can create a new slice from an existing array or slice using the [low:high] syntax. This creates a new slice header that points to the same underlying data. Crucially, the elements are not copied. If you modify the elements in the new slice, the original array or slice will reflect those changes.
go

data := []int{10, 20, 30, 40, 50}
subset := data[1:3] // points to index 1 and 2: {20, 30}

subset[0] = 99
// data is now {10, 99, 30, 40, 50}

The length of data[low:high] is high - low. If you omit low, it defaults to 0; if you omit high, it defaults to the end of the source.


*Warning: Be cautious when slicing. Because a slice keeps a reference to the entire underlying array, a small slice can prevent a large underlying array from being garbage collected. If you need to keep only a tiny piece of a massive dataset, create a new slice and copy the data over using copy().*

##Slice Internals: Length vs. Capacity

**Understanding the difference between len and cap is the key to writing performant Go code.**

    len(): The number of elements currently reachable in the slice.
    cap(): The total number of elements in the underlying array starting from the slice's first pointer.

## Exercises

    Create a slice of integers with a length of 0 and a capacity of 5. Loop 5 times to append the numbers 10, 20, 30, 40, 50 to it. Print the slice, its length, and its capacity after each append.


