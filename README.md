# Password Validator
> Detects weak passwords 

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This program accepts a newline delimited list of passwords from STDIN and prints invalid passwords that do not meet the following requirements:

1.  8 character minimum
2. 64 character maximum
3. only ASCII characters
4. Not be a common password



## Requirements

Golang >= 1.15

Supply a weak_password file named "weak_password_list.txt"

#### Build and test
```
go build -o password_validator
go test -v
```
## Usage

An example input_password.txt file is provided. Run with:
```
cat input_password.txt| ./password_validator
```
