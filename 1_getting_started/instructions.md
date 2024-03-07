## Getting Started
To start you'll need a linux or mac based machine.
If you wish to use a virtual machine please look into virtual box and follow this tutorial: https://www.youtube.com/watch?v=hYaCCpvjsEY

Next to install docker on linux:
Once you machine is updated you need to run the following commands
```sudo apt update && sudo apt upgrade  ```
```sudo apt install docker.io docker-compose curl wget git  ```
```sudo apt update && sudo apt upgrade  ```
```git clone //repo  ```
```sudo chown $USER /var/run/docker.sock ```
now if you run ```docker -v``` youll see the version printed in terminal

Great now that we have docker installed lets go to part 2 and learn about how to create a basic container.