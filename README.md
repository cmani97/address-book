# Address-Book

1. Run the main.go file or executable file to start the program
2. Provide your option to add or search contact
3. To search by name you need to provide the fullname -> `firstName lastName`
4. To make searchable easier, we are storing records by combination of `firstName+' '+lastName`
5. We can also make changes to search by either `firstName` or `lastName`, for this memory usage will be high. Because we need to store contact based on both `firstName` and `lastName`, to maintain performance

Sample output:
![Example Image](image1.png)
![Example Image](image2.png)

## Further Optimization:
### Concurrency:
#### The current implementation uses a simple read-write mutex for thread safety. For higher performance, especially with concurrent searches, consider using more advanced concurrency patterns or lock-free data structures.

