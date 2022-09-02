# firead examples

Examples demonstrating how to use [firead](https://github.com/JonnyOrman/firead).

It contains the following examples:
- `example1` - gets documents with string IDs
- `example2` - gets documents with int IDs
- `example3` - uses a struct to define the data to be returned with a string ID
- `example4` - uses a struct to define the data to be returned with an int ID
- `example5` - config retrieved from the environment
- `example6` - config retrieved from both the config file and the environment

# How to run

With Docker installed, clone the repo and run
```
docker-compose up --build
```

Create the collection "MyCollection" and seed test data for the Firestore emulator at `localhost:4000/firestore/data`

To use each of the examples, `GET` the following, substituting in the target document ID:
- `example1` - `localhost:3001/{id}`
- `example2` - `localhost:3002/{id}`
- `example3` - `localhost:3003/{id}`
- `example4` - `localhost:3004/{id}`
- `example5` - `localhost:3005/{id}`
- `example6` - `localhost:3006/{id}`

e.g. to get a document with ID "abc" from `example1`, request:
```
localhost:3001/abc
```