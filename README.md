# gomu

gomu is a "reflect" based Go helper package that unpacks slice into map. 

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/rosmak/gomu)



## Installation

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/rosmak/gomu
```

## Use case

If we want our rest-api response to be in format:

```json
{
  "uniqueID": {
    "Obj1":{}   
  },
  "uniqueID": {
    "Obj2": {} 
  }
}
```
Instead of

```go
[
  Obj1,
  Obj2
]
```
## Example

Let's imagine a situation where we have a User who can have multiple Phone Numbers and those Phone Numbers can be connected to multiple Social Media Accounts. So our models would look something like this:

```go
type User struct {
    ID              uuid             `json:"id"`
    Name            string           `json:"name"`
    PhoneNumbers    []PhoneNumber    `json:"phone_numbers"`
}

type PhoneNumber struct {
    ID              uuid             `json:"id"`
    Number          string           `json:"number"`
    Accounts        []Account        `json:"accounts"`
}

type Account struct {
    ID              uuid             `json:"id"`
    Platform        string           `json:"platform"`
}
```

When we are done with our db queries, and and populate all connections we get slices  
If we want to turn those slices into maps we need to add our map placeholder field to struct like this:

```go
type User struct {
    ID              uuid               `json:"id"`
    Name            string             `json:"name"`
    PhoneNumbers    []PhoneNumber      `json:"phone_numbers"`
    //important part is to add gomu thag that specify 
    //name of field which values would be used 
    //this line will say to gomu that "PhoneNumber" 
    //values will be used for "PhoneNumberMap" field
    PhoneNumbersMap map[string]Account `json:"phone_num_map" gomu:"PhoneNumbers"`
}

type PhoneNumber struct {
    ID              uuid               `json:"id"`
    Number          string             `json:"number"`
    Accounts        []Account          `json:"accounts"`
    //same here
    AccountMap      map[string]Account `json:"account_map" gomu:"Accounts"`
}

```

To get those map values populated  

If we have one user from db, just supply user to gomu function like this: 

```go
func main() {
    user := User{}
    err := gomu.MapAll(&user, nil)
    if err != nil {
        //handle error
    }
    //user.PhoneNumbersMap and PhoneNumber.AccountsMap are populated with walues
    //note that user.PhoneNumbers is set to nil, same for PhoneNumber.Accounts
}

```

If we have slice users from db, just supply users to gomu function like this: 

```go
func main() {
    users := []User{}
    err := gomu.MapAll(&users, nil)
    if err != nil {
        //handle error
    }
    //fields are populated 
}

```


In previous case all field are populated but we are still left with slice of users  
To get map of users we add our second argument to function like this:
```go
func main() {
    users := []User{}
    usersMap := make(map[string]User)
    err := gomu.MapAll(&users, &usersMap)
    if err != nil {
        //handle error
    }
    //fields are populated 
}

```

## IMPORTANT NOTICE

For now key of map is "ID" field of a struct, so in order to work your struct need to have ID field.  
Additionally "ID" type needs to implement Stringer interface, so it could be converted to string.  
For only string keys are supported.

## License
[MIT](https://choosealicense.com/licenses/mit/)
