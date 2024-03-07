Containers are lightweight isolated istances of a linux environment that run your app. When we create a container we need to specify a couple of things. Lets look at the example docker file and break it down for Nginx.

Dockerfile:

FROM nginx:latest 
COPY ./index.html /usr/share/nginx/html/index.html
EXPOSE 80

Breakdown:
```FROM nginx:latest``` FROM is a keyword that tells the docker program to fetch the image name from the dockerhub, a container repository and specifies the version to use. When we build our container this will be the base image that is used.

```COPY ./index.html /usr/share/nginx/html/index.html``` COPY is a keyword used to move files from your local filesystem into the Docker container during the build time. Here we are copying index.html in the current directory and placing it into ```usr/share/nginx/html/``` as index.html.

```EXPOSE 80``` EXPOSE keyword is used to expose an internal container port to the host system. This enables communication through to the containers system. Without it we would not be able to access the container as it is an isolated instance.

To turn this definition into a container we must create a ```Dockerfile``` this file has no extension  but is recognized by the Docker daemon during the build process.

To build this container we must first ensure you are in the correct terminal by using ```cd``` to enter this directory
Next you  need to build the Docker image using the following command in your terminal:
    ```docker build  -t my-first-container .```

    docker specifies to use the docker app. build is the  subcommand which tells docker hey, we are going to build something, -t flag sets a tag for our newly created image as my first container. Then the . tells docker to use the current directory to find the Dockerfile that is to be built.
You should get the following terminal output:

    Sending build context to Docker daemon  4.608kB
    Step 1/3 : FROM nginx:latest
    latest: Pulling from library/nginx
    e1caac4eb9d2: Pull complete 
    88f6f236f401: Pull complete 
    c3ea3344e711: Pull complete 
    cc1bb4345a3a: Pull complete 
    da8fa4352481: Pull complete 
    c7f80e9cdab2: Pull complete 
    18a869624cb6: Pull complete 
    Digest: sha256:c26ae7472d624ba1fafd296e73cecc4f93f853088e6a9c13c0d52f6ca5865107
    Status: Downloaded newer image for nginx:latest
    ---> e4720093a3c1
    Step 2/3 : COPY ./index.html /usr/share/nginx/html/index.html
    ---> a4f5250c76da
    Step 3/3 : EXPOSE 80
    ---> Running in c20d21c56e6d
    Removing intermediate container c20d21c56e6d
    ---> 0064f5d17d5e
    Successfully built 0064f5d17d5e
    Successfully tagged my-first-container:latest

Congratulations you have sucessfully built your first container.

Now we need to run this container to test that it works. To do this we need to use the following command:
```docker run --name my-container my-first-container```

This will start running the container and give it the name "my-container". We have specified it to use the image we have built by the my-first-container argument.
In the terminal you should now have the following output:
    /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
    /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
    10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
    10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
    /docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
    /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
    /docker-entrypoint.sh: Configuration complete; ready for start up
    2024/03/06 16:17:26 [notice] 1#1: using the "epoll" event method
    2024/03/06 16:17:26 [notice] 1#1: nginx/1.25.4
    2024/03/06 16:17:26 [notice] 1#1: built by gcc 12.2.0 (Debian 12.2.0-14) 
    2024/03/06 16:17:26 [notice] 1#1: OS: Linux 6.6.13-amd64
    2024/03/06 16:17:26 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
    2024/03/06 16:17:26 [notice] 1#1: start worker processes
    2024/03/06 16:17:26 [notice] 1#1: start worker process 29
    2024/03/06 16:17:26 [notice] 1#1: start worker process 30
    2024/03/06 16:17:26 [notice] 1#1: start worker process 31
    2024/03/06 16:17:26 [notice] 1#1: start worker process 32
    2024/03/06 16:17:26 [notice] 1#1: start worker process 33
    2024/03/06 16:17:26 [notice] 1#1: start worker process 34
    2024/03/06 16:17:26 [notice] 1#1: start worker process 35
    2024/03/06 16:17:26 [notice] 1#1: start worker process 36
    2024/03/06 16:17:26 [notice] 1#1: start worker process 37
    2024/03/06 16:17:26 [notice] 1#1: start worker process 38
    2024/03/06 16:17:26 [notice] 1#1: start worker process 39
    2024/03/06 16:17:26 [notice] 1#1: start worker process 40
    2024/03/06 16:17:32 [notice] 1#1: signal 28 (SIGWINCH) received

Now lets go over to our web browser and type in  http://localhost:80 into the browser.
Notice how we cannot connect and we get an error dispite the fact our container is running?

This is because the container is exposing its port 80 but we have not mapped this to a port on our host system to enable communication.
Lets run control+c in the terminal to kill the container and change the command used to run the container.

Before we change our containers run command we must delete it to prevent any conflicts moving forward. You can do this with the following command:
```docker rm my-container```

Now lets change the command to:
```docker run --name my-container -p 80:80 my-first-container```

The -p flag specifies a port  mapping, where source ip:sourceport -> destinationport
In this case we are saying map port 80 our host machine to port 80 inside of the container.
the sourceip:sourceport can used to map the port to a specific interface such as a dedicated network card or just to the localhost via 127.0.0.1 so that connections outside of your localmachine cannot access this service.

In ourcase we are not mapping it to an IP so any device on your local network will be able to connect, if your are using a cloud provided VPS this means anyone with your IP can connect.

Now lets test this works by heading to http://localhost:80 on your browser and you should be greeted with your first container.

Congratulations! You have now completed your first docker container! Before we go onto part 3 lets do the following:
1. kill the container and remove it
2. run it in the background so it runs more like a service
3. remove the container

to kill the container do control + c again

then ```docker rm my-container```

to run the container in the background we need to pass the -d argument. This means to deteach, to run not in the current shell but the background:
```docker run -d --name my-container -p 80:80 my-first-container```

Now after running this command you will only see a hash in the commandline such as "8307ae64bcf6ba3808947db06e37ae8a3f99dbb97c432c9203d2479f6f3d91a2".
This lets you know your container is now running and this can be verified with the ```docker logs my-container``` command.
To further verify your container is running you can use ```docker ps``` to see all running containers and their ID, IMAGE, COMMAND, CREATED, STATUS, PORTS and NAME.

Nice now we have run through some important side commands that you will be using plentifully across your docker career, lets move on!
