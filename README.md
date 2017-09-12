# BillBoards

This is a Go implementation of a BillBoard. 

Given a width, height and text this program finds the maximum font-size that can be used for the BillBoard.

## Input
The program accepts input from _stdin_, the first line contains a single integer _T_, the number of test cases. 
_T_ lines follow, each representing a single test case in the form "W H S". W and H are the width and height 
in inches of the available space. S is the text to be written.

## Output
Output T lines, one for each test case. For each case, output "Case #t: s", where t is the test case number 
(starting from 1) and s is the maximum font size, in inches per character, we can use.  If the text does not fit when 
printed at a size of 1", then outputs 0.

## Constraints

- The font is monospace and every character have the same width and height
- 1 ≤ T ≤ 20
- 1 ≤ W, H ≤ 1000
- The text will contain only lower-case letters a-z, upper-case letters A-Z, digits 0-9 and the space character.
- The text will not start or end with the space character, and will never contain two adjacent space characters.
- The text in each case contains at least 1 character and at most 1000 characters.

# Running
## Using docker
```bash
cat yourinputfile.txt | docker-compose run --rm billboards
```

## Local development
In order to compile the program in your local environment, you must install [glide](http://glide.sh/), then run 
the following commands:
 
```
glide install
go build
cat yourinputfile.txt | ./billboards 
```

# Developing with docker-compose
## Testing
```bash
docker-compose run --rm test
```
### Verbose output
```bash
docker-compose run --rm test -v
```

## Install a new dependency
```bash
docker-compose run --rm glide get <packageName>
```

## Fmt your code
```bash
docker-compose run --rm fmt
```
