Python is a very popular high end language and you will often find yourself working with it for various reasons. It could range from datascience to web development. 

Since Python requires its depenancies to be installed into the environment it runs from and cannot be compiled, Docker helps standardise its runtime environment. We can use a container to ensure all dependancies are installed, the correct and same version will be used everytime and environment variables also set. Environment variables are variables we might not want to expose in our online version control(Github) such as on such as secrets or passwords, but are needed for the application to run. 

Now that we have some background, in the api directory we have a python app that uses Flask to create an API. When you hit the /get route it will return the environment variable.

First lets create our Dockerfile and talk through it:

```
FROM python:3.8-slim

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the current directory contents into the container at /usr/src/app
COPY . .

# Install any needed packages specified in requirements.txt
RUN pip install -r requirements.txt

# Make port 5000 available to the world outside this container
EXPOSE 5000

ENV NAME=myname

CMD ["flask", "run", "--host=0.0.0.0"]
```


```ENV``` keyword is used to define a variable inside of the operating system that can be accessed later from within running programs. In ourcase our flask API.
Inside of app.py we have the following code:
```name = os.getenv('NAME', 'No name set')
        return {'name': name}
```
This sets a variable called name to  either whatever is stored in the environment variable `NAME` or if that's not provided, defaults to No name is set.
Then returns the variable name.

The os.getenv is a built in function that python uses to interface with the operating system to fetch variables that are set or exported in the OS. In ourcase it is fetching the variable we have set using the ```ENV``` keyword.

This time I am not going to give you the commands to see how much you've learned.using the build commands with the tag set to api and using the docker run command with port set to 5000:5000 you should have your container built and running.

To test your container load up postman, set the request type to get, address to http://localhost:5000/get and hit the send button.
If all went well you should have recieved your environment variable in a json response such as ```{"name":"myname"}```

