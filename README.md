Pizza management microservice

About the implementation. I hesitated over storage options. 
Now I am practically decided to try redis with persistence setting. 
It is an overkill here but at least I will see how it works.

In file api.go I have added comments for clarity of what I want to do.
Also I have divided logic behind the functions and the API itself into 
separate files. I like that kind of clarity especially with bigger projects.