# Go basic CRUD API with Fiber
This mini-project is a simple CRUD (Create, Read, Update, Delete) API built using Golang-Fiber with local MongoDB. <br> Connetcing to local MongoDB by using Fiber, following a tutorial from freeCodeCamp.org.

## Features
__Create__: Add a new employee into the MongoDB.<br>
__Read__: Get all employee.<br>
__Update__: Edit employee.<br>
__Delete__: Delete a employee from the collection in the MongoDB.<br>

## API Endpoints
POST http://localhost:3000/employee - Add a new infomation to the local MongoDB..<br>
GET http://localhost:3000/employee - Retrieve all employee.<br>
PUT http://localhost:3000/employee/{id} - Update an existing employee by its ID.<br>
DELETE http://localhost:3000/employee/{id} - Delete a employee by its ID.<br>

## Tutorial Reference
This project is based on the freeCodeCamp.org tutorial:[⌨️ HRMS with Golang Fiber](https://www.youtube.com/watch?v=jFfo23yIWac) 
