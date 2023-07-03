# Bibgen
Bibgen is a CLI tool for generating bibliography in ABNT2 standard.

## Installation
```go
go build -o bibgen *.go
```
I recommend you to move the binary to /usr/sbin or something like that.

## Usage
If you want to reference a book use
```
bibgen -b
```
And fill the informations that will be asked

For websites use
```
bibgen -w
```
And fill the informations

Feel free to edit the source code and adapt the software to your use
