# first-rest-api

# user-service

# REST API

GET /users -- list pf users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204/200, 404, 400, 500
DELETE /users/:id - delete user -- 204, 404, 400, 500