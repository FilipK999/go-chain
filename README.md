# Go Chain

<hr>

A very simple and basic implementation of a blockchain,
written in Go and stored in a Redis database.
Has no actual use, I built it just for fun, to get familiar with the subject. 


### If anyone ever wants to run the code for some reason:

#### Prerequisites
1. Install Go and Redis
2. Start the redis server (Program uses the default DB, so if you have ever used Redis, make sure there's nothing 
   important there)
3. That should be it :)

#### Run the code

1. Clone the repo and cd to it.
1. ```go run main.go add -block BLOCK_DATA``` to add a block
1. ```go run main.go print``` to show all created blocks

