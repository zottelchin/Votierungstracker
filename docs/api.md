# The API
Default endpoint: `/api`

## GET /tasks/:account/:course
Returns the tasks for the specified course.

### Response body
```
[
  {"id": 0, "name": "Example", "points": 4, "maxPoints": 5, "presentations": 0},
  …
]
```

## POST /tasks/:account/:course
Adds a task.

### Request body
```
{
  "name": "Example",
  "points": 4,
  "maxPoints": 5,
  "presentations": 0
}
```

### Response body
```
{ "id": 0 }
```

## PUT /tasks/:account/:course/:id
Updates a task.

### Request body
```
{
  "name": "Example",
  "points": 4,
  "maxPoints": 5,
  "presentations": 0
}
```

### Response body
```
{ "updated": 1 }
```

## DELETE /tasks/:account/:course/:id
Deletes a task.

### Response body
```
{ "deleted": 1 }
```

## GET /api/courses/:user
Get the courses of a specified user.

### Response body
```
[
  "example",
  …
]
```

## GET /api/pres/:account/:course
Get number of times a user have to present in a given course.

### Response body
```
{"presentations": 1}
```

## PUT /api/pres/:account/:course
Set number of times a user have to present in a given course.

### Request Body
```
{"presentations": 1}
```

### Response Body
```
{"updated presentations":1}
```
 The number provided by the response is the affected row in the database.

 ## GET /api/perc/:account/:course
Get the percentage of subtasks a user has to do in one course.

### Response body
```
{"percentage": 1}
```

## PUT /api/perc/:account/:course
Set the percentage of subtasks a user has to do in one course.

### Request Body
```
{"percentage": 1}
```

### Response Body
```
{"updated percentage":1}
```
 The number provided by the response is the affected row in the database.
