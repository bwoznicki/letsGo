# Overview

## Why GO ?
A list of organisations / people and reasons for them to use GO.  
https://github.com/golang/go/wiki/WhyGo

Success stories  
https://github.com/golang/go/wiki/SuccessStories

Go demand  
https://spectrum.ieee.org/view-from-the-valley/at-work/tech-careers/go-language-tops-list-of-indemand-software-skills

With so many languages out there why do we even need go ?
- design / syntax https://en.wikipedia.org/wiki/Go_(programming_language)#Design
- goroutines / threads https://www.geeksforgeeks.org/golang-goroutine-vs-thread/
- parallelism / concurrency https://vimeo.com/49718712

## Go syntax
- go syntax is strict, similar to C with great emphasis on readability
- 20yo battle tabs vs spaces - solved ! https://thenewstack.io/spaces-vs-tabs-a-20-year-debate-and-now-this-what-the-hell-is-wrong-with-go/
- verbose declarations: https://blog.golang.org/declaration-syntax  

C
```C
int x; # x is an integer
int *x; # x is a pointer to integer
int x[10]; # x is an array of 10 integers
int *(*p)( int, int *);
```

C's spiral rule http://c-faq.com/decl/spiral.anderson.html  

GO
```GO
var x int # x is an integer
var x *int # x is a pointer to integer
var x [10]int # x is an array of 10 integers
func p(int, *int) *int
```

## Performance

example benchmarks:
- GO vs C++ https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-gpp.html
- GO vs Python https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-gpp.html  

Simple increment value by 1  
comparing C, Go, Python with PERF

C program
```C
int main() {
    int x = 0;
    x += 1;
}
```

GO program
```GO
package main

func main() {
    var x int = 0
    x += 1
}
```
Python script
```python
#!/usr/bin/python3

x = 0
x += 1
```
# lets Go!
## Installation

Getting started https://golang.org/doc/install
1. download package for your system.
2. extract from download dir to /usr/local

```bash
ls ~/Downloads | grep "go1." # copy output
sudo tar -C /usr/local -xzf <output> # eg go1.14.1.linux-amd64.tar.gz
```
3. Add **GO** to your path.
```bash
# In .bash_profile / .profile etc depending on the system:
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile

# confirm by checking the version
go version
```
## Hello World

```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

## Parts of the GO program

- package - **main** package is the starting point of any GO program.
- import - for stl and other packages.
- main() - program entry point

## Type system

"the more important idea is the separation of concept: data and behavior are two distinct concepts in Go, not conflated into a single notion of class”. – Rob Pike

### Basic types
- numbers
  - **int**  int8  int16  int32  int64
  - uint uint8 uint16 uint32 uint64 uintptr
  - float32 float64 (1.555)
  - complex64 complex128 (eg. 5+2i)
  - rune (int32 representation of unicode code point)
  - byte (uint8)
- strings
  - string
- booleans
  - bool
- other
  - array
  - slice
  - map
  - struct
  - pointer
  - channel
  - interface
  - function


## Additional resources
- Go tour. Interactive hands on Go course. https://tour.golang.org/welcome/1
- The Little Book of Go. https://www.openmymind.net/The-Little-Go-Book/