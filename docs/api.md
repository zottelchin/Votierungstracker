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
