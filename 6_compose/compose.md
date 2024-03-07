So now you should be pretty confident in building some containers. Now you have made a front end container and backend container, they can talk to each other. But how can we orchestrate them togther instead of typing lots of commands to manage them? This is where docker-compose comes in.

docker-compose.yml files are yaml scripts that describe how a container or set of containers should be deployed and uses a nifty single docker-compose up to run the container.

This not only saves time but reduces the human error and makes the running of complex mulit container systems a lot easier.