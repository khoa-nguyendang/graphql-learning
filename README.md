# graphql-learning
practice graphql  golang


# Start application

```#bash
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
mutation CreateAuthor {
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