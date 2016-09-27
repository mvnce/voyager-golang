#Voyager Backend in Go with Gin Framework
Providing RESTful APIs for Voyager blog web application

[![Build Status](https://travis-ci.org/vincentsma/voyager-golang.svg?branch=master)](https://travis-ci.org/vincentsma/voyager-golang)
[![codebeat badge](https://codebeat.co/badges/aa561012-c016-40eb-a898-ee72da718e50)](https://codebeat.co/projects/github-com-vincentsma-voyager-golang)

#Task
- ~~Database schema~~
- ~~Module design~~
- ~~API design~~
- ~~Implementation~~
- Testing
- ...

#Examples
###Sign Up
`curl -i -X POST -H "Content-Type: application/json" -d "{\"email\": \"admin@domain.com\", \"password\": \"admin\"}" http://localhost:8080/api/v1/auth/signup`
###Sign In
`curl -i -X POST -H "Content-Type: application/json" -d "{\"email\": \"galao@gmail.com\", \"password\": \"galao\"}" http://localhost:8080/api/v1/auth/signin`
###Validate Token
`curl -H "Authorization: Bearer x.x.x" http://localhost:8080/api/v1/auth/validate`

###Get Posts
`http -f GET http://localhost:8080/api/v1/posts "Authorization:Bearer foo.bar.zoo"  "Content-Type: application/json"`
###Get Post
`http -f GET http://localhost:8080/api/v1/post/:id "Authorization:Bearer foo.bar.zoo" "Content-Type: application/json"`
###Add Post
`N/A`
###Delete Post
`N/A`
###Update Post
`N/A`



-Vincent
