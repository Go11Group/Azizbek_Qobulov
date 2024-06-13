package modul

type Person struct {
    ID         int    `form:"id"`
    FirstName  string `form:"first_name"`
    LastName   string `form:"last_name"`
    Age        int    `form:"age"`
    Gender     string `form:"gender"`
    Nation     string `form:"nation"`
    Field      string `form:"field"`
    ParentName string `form:"parent_name"`
    City       string `form:"city"`
}
