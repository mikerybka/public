package contacts

type Root struct {
    Friends map[string]Person
    Coworkers map[string]Person
}

type Person struct {
    Name string
    EmailAddress string
    PhoneNumber string
}
