# gotasks

**Run**

`sh run.sh`

... and go to http://localhost:8080/

## API

**Models**

Task:
* id: `string`
* name: `string`
* description: `string`
* active: `bool`
* time: `number`
* username: `string`

--

*Create a new task*

Route: `/add`
Method: `POST`
JSON Body: `Task`

*Get all tasks*

Route: `/getall`
Method: `GET`

*Update a task*

Route: `/update`
Method: `POST`
JSON Body: `Task` (id required)
