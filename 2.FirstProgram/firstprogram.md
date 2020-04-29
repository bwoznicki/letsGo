# First Go program

Writing CAT in Go
- imports
- stl packages
- input args
- reading files
- writing to standard output

## local GOPATH path  
https://golang.org/cmd/go/#hdr-GOPATH_environment_variable  
in your ~/.bash_profile or ~/.profile etc.  

export PATH=$PATH:/usr/local/go/bin  
export GOPATH=$HOME/go  
export PATH=$PATH:$GOPATH/bin

### in your home dir:  
mkdir -p ~/go/pkg ~/go/src ~/go/bin

### reload env  
source ~/.bash_profile

### Basic Go tools
- go run
- go build 
- go install

### Additional content
- go **io** - https://golang.org/pkg/io/
- go **fmt** - https://golang.org/pkg/fmt/
- go **os** - https://golang.org/pkg/os/
- go code structure - https://talks.golang.org/2014/organizeio.slide#1  
- go commands - https://golang.org/cmd/go/
