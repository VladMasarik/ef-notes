# Quick start
## Run the app
```
docker build -t notes-app .
docker run notes-app -p 8080:8080
```

## Test the app
```
curl http://localhost:8080/notes

curl -X POST http://localhost:8080/notes -d '{"text": "this is my note"}'

curl http://localhost:8080/notes/1

curl -X PUT http://localhost:8080/notes/1 -d '{"text": "this is my UPDATED note"}'

curl -X DELETE http://localhost:8080/notes/1

curl http://localhost:8080/notes/1
```

# Extra notes / my considerations
## Separation
I wanted to avoid putting everything into folders, but then some parts are just natural, and it is strange that other things are not in folders.

## Premature optimization
I dont think the struct calls on notes are really needed, like if you literally just want to take notes then just use functions. But overall, almost every API evolves, and I eventually start using the "class-like" structure, so I start with it every time as well, just because I dont want to be then refactoring everything.

## DB
Of course you can choose any DB, I used sqlite just so that you can run it yourself easily as well.

I have a package-wide DB client, as I prefer to always have just a single way of doing things, so that it forces me into being consistent.

## API versioning
I have exposed everything right at the root, and without any API versioning. Up to a discussion on how we want to version the API, whether through URL or headers, and especially where we want to have it exposes

# Comments / descriptions
I have not added package comments as I think they are obvious.

## API responses
Possibly dont return every single thing when quering the whole collection, but maybe just an ID in this case?

I would improve the error reporting of the API in terms of what messages it is sending back.

## User input
Of course we need to properly check whether the fields are valid, or define them as required in the DB schema.

## Over two hours
Unfortunatelly, I busted myself on the Golang docker image not having set the $GOROOT and $GOBIN paths by default, which caused the multistage build to fail, so I just used plain `go build`. Then my tests stopped working because the DB had issues with pathing as well.