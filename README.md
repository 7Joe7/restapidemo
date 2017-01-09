Pizza management microservice

About the implementation. I hesitated over storage options. 
Now I am practically decided to try redis with persistence setting. 
It is an overkill here but at least I will see how it works.

In file api.go I have added comments for clarity of what I want to do.
Also I have divided logic behind the functions and the API itself into 
separate files. I like that kind of clarity especially with bigger projects.

I decided for a redis hash for the pizzas. Each pizza will have an id and 
a JSON value with all ingredients.

I have divided code into multiple packages rest package with rest api and 
for now also logic (could be separated if it would get bigger). Db package 
for decoupling of used db and the rest of the code.

After some thinking I realized that ingredients shouldn't be a sub-resources as 
they are not unique to each pizza. This could be solved multiple ways 
(third resource unique to a pizza which would contain amount of ingredient and its id). 
I think for now separation is cleanest.