# gf-markdown link checker

cicd tool for validating markdown files contains no broken links, so that it can be added as a check as part of the build. If you rename a file and don't remember to update your readme, then `markdown` will stop the build until all the links are correct.

## Installing

```
go install github.com/goblinfactory/gf-markdown
```

## Usage

```css
gf-markdown testdata/**/*.md
```
*Run from the root of your solution using glob patterns e.g. `gf-markdown testdata/**/*.md`* 

![markdown testdata/**/*.md](markdown2.png)

## Ignoring folders

```css
gf-markdown **/*.md
```

## Verbose output

*Add -v for verbose output. (will display status of all links, valid as well as broken)*
![gf-markdown testdata/**/*.md -v](markdown1.png)


## Adding to makefile

After installing the tool, simply add the line `gf-markdown **/*.md` to your makefile. This will exit with (-1) fatal, and stop any build if added to a makefile and there are errors.


```yaml
.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet: fmt lint
		go vet ./...
.PHONY:vet

build: vet
		go test -tags integration ./...
		gf-markdown **/*.md
		go build ./markdown.go
.PHONY:build
```
## Does not support multiple glob patterns

Does not support multiple glob patterns, to support something like `gf-markdown **/*.txt **/*.json` simply call markdown multiple times in your makefile, one for each glob pattern e.g.

```yaml
build: 
		...
		gf-markdown **/*.txt
		gf-markdown **/*.json
```

## Calling from your own code (using the markdown API)

See [markdown_integration_test.go](markdown/markdown_integration_test.go) for example of calling markdown from your code.

(Hopefully will be able to link to Google's auto documenter as soon as I setup pipeline to create and sign packages.)

| method | params |
| --- | --- |
| [RunFromArgs](markdown/markdown.go)  | `(args []string, printer *Printer) Result` |
| [GetReport](markdown/reports.go)  | `(filename string) Report` |


## Printer package

- [Printer](markdown/printer.go) : *Buffered printer to support integration testing*

**How printer works**

Please note: If code exits via log.Fatal(), then defer does not run, and printer will not flush. 

```go

	func TestDoSomething(t *testing.T) {

		// create a buffered printer, and defer all printing
		p:= &Printer{}
		defer p.Flush()

		// pass printer to anything that would typically print to the console
		// when you're done, call flush to print to the console.
		// flush will print all buffered lines to console, and add printed 
		// lines to history.

		addNums(p, 1, 3)
		addNums(p, 2, 3)
		
		assert.Equal(t, p.Lines(), []string { 
			"1 + 2 = 3" 
			"2 + 3 = 5" 
		})
		p.Flush()
		Greet(p, "Greg")
		assert.Equal(t, p.Lines(), []string { 
			"Hello Greg"
		})
		p.Flush()
		
		// history contains everything printed
		assert.Equal(t, p.History(), []string { 
			"1 + 2 = 3" 
			"2 + 3 = 5" 
			"Hello Greg"
		})
	}

	func addNums(p *Printer, a int, b int) {
			p.Println("%d + %d = %d", a, b, a+b)
	}

	func Greet(p *Printer, name string) {
		p.Println("Hello %s", name)
	}

```
	
## Backlog, Ideas 

- Add a `-i ignore` to allow you to ignore folders.