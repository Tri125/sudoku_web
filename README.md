# sudoku_web

sudoku_web is a sudoku puzzle solving web app written entirely in Go.
The solving algorithm is a recursive backtracking algorithm, allowing fast solving of standard 9x9 grids. 

In the current version, net/http handle all the web aspect of the apps.

![Running sudoku_web](https://i.imgur.com/i8yqkrj.png)

## Dependencies
* Go 1.6 
* gorilla/mux for routing
* jcelliott/lumber for logging
* tri125/sudoku for handling the data and solving the grid

## Configuration

In the main function of the file main.go edit the port number to whichever port you want the application to be listening to.

The default port is 80.

~~~
func main() {
	//Port that the app will listen for requests
	var port int = 80
~~~

## Installation

You can compile and run the code directly from a Linux machine with Go 1.6 installed or you can create a Docker image from the provided Dockerfile.

### Compiling the code

Make sure that your system have **Go 1.6** installed and download sudoku_web.

Install the dependencies:
~~~~
go get github.com/gorilla/mux
go get github.com/jcelliott/lumber
go get github.com/tri125/sudoku
~~~~

You can then compile & run the code by running the following command from the sudoku_web folder.
~~~~
go run /src/main/main.go
~~~~
It is important that this command is ran from the project directory, otherwise the web server will have problems serving the static files.
If you want to keep the app running, make sure to send the command to the background like so:
~~~~
go run /src/main/main.go &
~~~~

### Creating the Docker image

If you don't have Go installed on your machine you might prefer creating a docker image, it is faster to set up.

The sudoku_web Docker image is based from Linux wheezy which can make the image sizeable (650 MB for the base image).

**Build the image:**
~~~~
docker build -t sudoku_web .
~~~~
The Dockerfile will handle most of the work for you and will create an image with the name "sudoku_web".

**Create & run a container:**
~~~~
docker run -d -p 80:80 sudoku_web
~~~~

The -d parameter will run the container in a detached state. 

The -p parameter will map your linux machine port 80 (the leftmost) to the internal port 80 of the container (rightmost).
If you want to run the app on port 500 of your linux machine you would run the following command:
~~~~
docker run -d -p 500:80 sudoku_web
~~~~

If the container is already created, just run it with the **docker start** command.

In the current state, there is no graceful shutdown.

*SIGTERM* is not handled and so the **docker stop** command will result in a 10 seconds delay followed by a *SIGKILL* to the container.

## Known issues

* Due to the limited if not complete lack of data validation in the current version of sudoku_web, 
it is possible to block goroutines by inputting an incorrect puzzle to solve. However, the application will continue running correctly for every other users.

* The logging functionality provided by lumber is not entirely functional.

* The Docker image doesn't handle *SIGTERM*
