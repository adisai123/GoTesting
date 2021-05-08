# GOSample

For go in-built unit testing steps:

    1. Create file with name like <extingfilename>_test.go (ex. for sort.go test filename will be sort_test.go)
    2. add import "testing" package
    3. Add method with prefix Test and having parameter t testing.T 
        ex : (func TestBubbleSortNoError(t *testing.T))
    4. add validations : Example
        //validation
        if elements[0] != 9 {
            t.Error("first element should not be 9")
        }

Commands using with this project

    1.  go test -cover     (inside folder)
    2.   go test -bench=.

Benchmark:
    Create function Name starting with Benchmark (example:
    func BenchmarkSort(t *testing.B){  for i = 0 ; i < b.N ; i++{} })

    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/testinggo/GoTesting/src/api/services$ go test -bench=.
    [1 2 2 2 3 3 4 4 4 5]
    goos: linux
    goarch: amd64
    pkg: mygo/testinggo/GoTesting/src/api/services
    BenchmarkSort-4          2894872               402 ns/op
    PASS
    ok      mygo/testinggo/GoTesting/src/api/services       1.592s
            

----------------------------------------

Value Type - int float string bool struct 
Reference type -  slices maps channels pointers functions 
