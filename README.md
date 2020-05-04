# GOSample

For go in-built unit testing steps:

    1. Create file with name like <extingfilename>_test.go (ex. for sort.go test filename will be sort_test.go)
    2. add import "testing" package
    3. Add method with prefix Test and having parameter t testing.T ex : (func TestBubbleSortNoError(t *testing.T))
    4. add validations : Example
        //validation
        if elements[0] != 9 {
            t.Error("first element should not be 9")
        }

Commands using with this project
    1.  go test -cover     (inside folder)
