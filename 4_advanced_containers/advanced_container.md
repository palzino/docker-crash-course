So far you've made a container and added a bind mount. Now lets step it up a notch and build a more advanced container from scratch.

Now that we have a basic webserver from the Nginx base image that serves some html content. Lets create a backend so that it can fetch and store data.

To give you a taste of how docker works we aren't going to use a single framework or language, we are going to run through 3 example API's written in different languages and talk through turning them into containers.

To test our API's work you need to install postman which is available by using ```sudo snap install postman --classic```

Now in each of these sections I will cover a different useful topic:
Inside of go we have builder containers
Inside of python we have environment variables

Inside of javascript we have building a node environment because javascript is enough pain in itself.

See you inside the subdirectories of this more advanced section!

