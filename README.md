# graphql-learning
practice graphql  golang

# project's requirements

```
go version go1.21.0 linux/amd64
node v20.11.1
pnpm v9.1.1
Docker version 26.1.1 
Docker Compose version v2.27.0
```

# Start application

```#bash
$ docker-compose build
$ docker-compose up
```

# Use GraphQL Postman
Endpoint `localhost:8080/graphql`

# Add new Author
```
mutation CreateAuthor {
    createAuthor(email: "author1@gmail.com", name: "author1") {
        created_at
        email
        id
        name
    }
}
```

# Add new Post
```
mutation CreatePost {
    createPost(title: "Post 2", content: "Post 2", author_id: 1) {
        content
        created_at
        id
        title
    }
}
```

# Query all posts and it's author
```
query Posts {
    posts {
        content
        created_at
        id
        title
        author {
            email
            id
            name
            created_at
        }
    }
}

```