# Source code for the presentation: Containerization of the Go projects

On the `master` branch, there is the final version of the image (version 4)

To switch between different versions of the `Dockerfile` version, change the git branch:
* 1 - for the simplest approach
* 2 - for the basic multi-stage approach
* 3 - for the stripped of multi-stage approach (the most optimal solution)
* 4 - for the approach with the additional features of the Dockerfile 

>Note Remember, this application might be an example of the working `go` application but **it has been not** created with all good practices on mind :)

## Running

To run the app, in root directory run `docker-compose up -d`. It will setup the application alongside with the `mongoDB` instance.

### Endpoints
`GET /get`<br>
To fetch the all data from the DB

`POST /post`<br>
BODY:
```json
{
  "user": "str",
  "description": "str"
}
```
To add some data into DB
