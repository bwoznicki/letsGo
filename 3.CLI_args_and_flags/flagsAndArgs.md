# CLI arguments and flags

## RECAP

Why go install did not work last time ?

### local GOPATH path  
https://golang.org/cmd/go/#hdr-GOPATH_environment_variable  
in your ~/.bash_profile or ~/.profile etc.  

export PATH=$PATH:/usr/local/go/bin  
export GOPATH=$HOME/go  
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin

#### in your home dir:  
mkdir -p ~/go/pkg ~/go/src ~/go/bin

#### reload env  
source ~/.bash_profile

## Args and Flags in cli programs

`mycli -f --flag arg1 arg2`

flags:
- -f
- --flag

arguments:
- arg1
- arg2

## Aditional resources
- go os package - https://golang.org/pkg/os/
- go flag package - https://golang.org/pkg/flag/