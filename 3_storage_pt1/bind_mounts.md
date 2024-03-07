Great so you've made a container that contains a html page and serves it via nginx.
But at the moment its baked into the container during the build process and the only real way to update it is to destroy and rebuild the container.
This is great and how it should be done for production builds where you do not want persistant data to live in the container. However, during development or in a local non production environment im sure this can get frustrating.

Docker containers themselves are non persistant, if you were to take in user inputs and save them to a file, then delete the container and rebuild it e.g you have to update the app, you've lost all your data.

There are two solutions to these problems and we are going to use both to solve these problems.

Bind Mounts and DockerVolumes!

Firstly, let's talk about Bind mounts.

Bind mounts are used when you want to share directories between the host and the container. This enables the container to fetch the contents of its directory from the host.
This can be helpful when you have html, css and js that you wish to store on your machine to develop with but want your changes reflected in realtime inside of the container.
Lets test this out.

```docker run -d --name storage -p 80:80 --mount type=bind,src=$PWD,target=/usr/share/nginx/html my-first-container```
Before we run this command lets break it down, so we have the run detatched with name of storage, on port 80 from the previous chapter. Now with ```--mount``` we are mounting a host directory into the container by specifing
type of bind with the source being the current working directory and the target being the ```/usr/share/nginx/html``` directory. If you recall building your docker image you will see that we had copied index.html into this directory but now we are fetching it from the host system.

Now that you have run your container, lets update the index.html  file with some new content, Change first container to second container. then refresh your page.
Nice, now your containerised website is rendering updates from the bind mount.

Next we are going to cover  Docker Volumes but before this we need to create a backend service that can recieve the data from the frontend and store it locally.
Since nginx is only serving static content, it doesnt have a way of taking inputs from the client and storing  them in the filesystem safely. So we need to build an app to do this.
This is also a great transition into how we can use docker networks to send data from one container to another.
So onto part 4 we go before we revist storage part 2!