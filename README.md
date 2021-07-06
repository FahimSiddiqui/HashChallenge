**Hash Generator:**

Hash generator is a tool to hit a resource URL and generate a hash value of it using hashing-algorithm. By default its using `MD5` hash generation technique but the tool has an ability to add more algorithms in future. 

`Only using standard GO library, there is no need to install any external go dependency.`

Hash generator internally uses a in-memory(local) cache to make sure same URLs are not being hit again. For each given URL, a hash code will be generated and placed temporarily into a map. Next time if same URL is hit, it would pick hash value from the cache rather than hitting the URL.
Note: There are no cache eviction logic written(as of now) but LRU could have been used to make it smarter.

Hash generator internally uses Goroutines to maximize the usage of CPUs and improving the performance of the operations. Tool also has an ability to limit the number of parallel requests to prevent local resources.



**Usage:**

In order to run it on local, kindly make sure you have GO installed and configured properly.
Tool accepts a flag called `-parallel` to limit the maximum number of parallel requests. If not provided, it will be set to 10 by default.

Build the project: Use command `go build`, which will create a `HashChallenge.exe` in the main folder of the repository.

Run the project: Use command `.\HashChallenge -parallel=2 http://google.com facebook.com twitter.com google.com`

Expected output of above command:


http://google.com a9c4c891e975d3cfcab4607b4d7c18a1

http://facebook.com 08165d33fe6ca698bfcbfb5a2c65eaee

http://google.com a9c4c891e975d3cfcab4607b4d7c18a1

http://twitter.com 1e49a4450a5c2adff5e3b367d58e0419


Refer: Please refer `Sample_Run_HashGenerator.txt` which has more integration test-cases defined. For integration test-cases in production, we can use integation-automations-tools to cover if not all then most of the end-to-end flows upfront.



**Test:**

HashGenerator tool comes with set of unit test cases across different packages. Not all the edge cases are captured but basic set of test cases written for each function/functionality test before the tool hits production.

How to test: Use command `go test -v ./...` it will go through each package and run the test case. Ideally all the test-cases should pass. Test cases are written in a niave approach where `if..else` is used to check if something will pass or fail. In production `GoConvey` package can be used to simplify & group more and more test cases.



**Assumptions:**

1. If HTTP scheme is not provided in the input URL, this tool will add as prefix.
2. Only MD5 hash will be created for the input. Used a factory-design pattern to switch to different hashing algorithms in future.
3. Used singleton design pattern to initialize a cache, which can be accessed by multiple goroutines simultaneously.
4. Write operation in cache requires a MUTEX lock. For read operation, its not needed.
5. Demonstrated the behaviour of Logging mechanism using simple print statements but it can be further well-designed & used using `Logrus` package.


**Future Scope:**

HashGenerator is a demonstration tool to take URLs as input, hit the URLs, generate hash value of the output and print it. In future we can implement it in the form of REST APIs, use better logging techniques, implement more and more hashing algorithms, a REDIS cache, other golang packages to simplify test cases, 

