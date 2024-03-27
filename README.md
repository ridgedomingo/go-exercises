# Golang sample exercises
This project contains different golang packages that you can use when learning go lang.

## Exercise 1 - Password generator
This package generates random password based on cli arguments.
### Installation
run ```go get github.com/ridgedomingo/go-exercises/pkg/generator```

### Usage
Below is a sample on how you can use the package from your local. For a list of complete arguments see table [below.](#cli-args)

```go run . --length=6 --includeNumbers=true```

### CLI Args
| Args        | Description                                               |
|-------------|-----------------------------------------------------------|
| --type       | kind of password that will be  generated, valid values are random, alphanumeric, and pin(defaults to random)              |
| --includeNumbers        |  Set to true if password should contain numbers                                                                |
| --includeSymbols | Set to true if password should contain numbers(Only works if type is set to random)                                   |
| --includeUppercase  | Set to true if password should contain an uppercase letter                                                         |
| --length         | How many characters will be generated. (defaults to 6 if type is pin, and 12 if type is random or alphanumeric        |

### Notes
- Alphanumeric type will only generate characters, and number regardless if includeSymbols flag is set.
- Pin type will only generate numeric password regardless of other flags set
