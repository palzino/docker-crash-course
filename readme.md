Are you new to docker? Heard the buzz around containers but don't understand the hype?
In this crash course I aim for you to be able to understand and do just more than the basics.
Going from zero to almost hero.

## What is docker?
Docker is a platform that allows developers to create an isolated yet shared environments for apps.
It's shared because it shares the same underlying linux kernal as its host machine to save resources.
It's isolated because the app instance is selfcontained similar to how a virtual machine runs.
Running multiple containers means that will be isolated from each other and need to communicate through a network, but they will all share the kernal of the base host operating system.

## Requirements and Download
To start this crash course you should already be familiar with a Linux distro or MacOS. You can follow along on windows using WSL but that will not be covered here nor will MacOS.
I recommend using an Ubuntu 22.04 virtual machine for ease of use.
```sudo apt install git``` then run
```git clone https://github.com/palzino/docker-crash-course.git``` to download the folder and ```cd docker-crash-course``` to enter the directory
Please ensure you have a code editor installed such as vscode. if you are using ubuntu then you can run ```sudo snap install code --classic``` and then ``` code .``` to open VSCode

