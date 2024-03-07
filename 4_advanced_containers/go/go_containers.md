To continue with this I would recommend you use the precompiled binary called ```backend``` in the directory backend but you can also try and compile the app yourself for a challenge.

I have made a little Golang app that listens to data coming in on the /post route and saves it to a file, a route called /get that pushes that data to the webpage.
Now lets build this container.

We need a new ```Dockerfile``` inside of our backend directory so make that firstly.

Then lets break it down:

1. we need to define our base image that the container will use
2. set a working directory  (where the CMD command will be executed)
3. copy all files from the current directory into the container
4. expose our port
5. set a command to run the backend when the container starts

Lets go!!

```
    FROM alpine:latest

    WORKDIR /app

    # Copy the compiled binary and any other necessary files
    COPY backend /app/
    RUN mkdir /data
    EXPOSE 8080

    CMD ["/app/backend"]
```
Unlike defining NGINX which is a precompiled instance of nginx within an OS, here we are taking Alpine. A lightweight linux distribution and using this as our base image. 

```WORKDIR``` keyword is used to define the working directory of the app. This is where commands will be execued from and your path will be set to as default when calling or interacting with files inside of the container.

```RUN``` is the keyword that is used to execute bash commands from within the container. In this instance we are using it to execute the creation of a new directory at the path /data.
 ```EXPOSE``` will ensure the containers 8080 port is available
 and CMD``` tells Docker what command to use when starting the application.

 Now from within the backend directory , you can now build your docker image using the following command :
 ```docker build -t backend .```
 and to run the container do the following:
 ```docker run --name backend -container -p 8080:8080 backend```

 Now lets test our backend by using Postman. 
 
 Open up Postman, enter http://localhost:8080/post in the address bar, change get to post, in the body header click raw and enter ```{"name":"anyname"}``` then click send.

 This has sent a json object with which contains a key value pair of "name" & message of "anyname". If it was sucessful it should return ```{"message":"Data received successfully","name":"name"}```

 Now to test the get function lets change the request type from post to get and change /post to /get at the end of the url. This should now return ```{"name":"name"}```. 

 Great, how we have a very basic HTTP API working in GO!

 Following these primitives you are able to  create as complex or simple images as needed for your future projects. You can use a base linux layer to run your compiled binaries. However this only using a fraction of dockers power at so lets take a look at the more advanced section and compile code code within the container at build time.

 When you make an app and you test it locally, it works just fine, but then someone else tries to build it or it runs into runtime errors. This is a huge painpoint in operations and development alike, so containers standardise the build process. Lets replicate this and build the app at container run time so that it is interoperable on all machines.(At least for the same system architecture in this case).

 Lets firstly stop, delete and remove the image of the container we've just made
 ```docker stop backend; docker rm backend; docker image rm backend```

 Now lets go back to the Dockerfile you made and change it up.
 
```
FROM golang:1.22-alpine as builder

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o own_binary ./backend.go

FROM alpine:latest  

WORKDIR /app

COPY --from=builder /build/own_binary .

EXPOSE 8080

CMD ["./own_binary"]
```

1. ``` as builder ``` is used to define our first container image that is going to build our app using the golang docker image based on alpine
2. We set our WORKDIR and copy over our files, download  dependencies with `go mod download` (we dont have any in this example so they get skipped)
3. ```RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o own_binary ./backend.go``` is used to build the backend app and sets it to name of own_binary through the -o argument
4. we define a new base alpine image and set our workdir
5. We copy the binary from the previous build step into the new container into the working dir
6. expose 8080 and set the run command.