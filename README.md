# apsis
 <!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol> 
    <li><a href="#run">Built with</a></li>
    <li><a href="#run">How to run app locally</a></li>  
     <li><a href="#run">API doc</a></li>  
      <li><a href="#run">Demo</a></li>  
  </ol>
</details>

## Display app:

 App => [Click Here to view app](https://apsis-code.herokuapp.com/).

## Built With:

Discover some packages of this project:

* [chi](https://pkg.go.dev/github.com/go-chi/chi/v5)
* [testing](https://pkg.go.dev/testing)
* [fmt](https://pkg.go.dev/fmt)
* [log](https://pkg.go.dev/log)
* [os](https://pkg.go.dev/os)
* [bufio](https://pkg.go.dev/bufio)
* [log](https://pkg.go.dev/log)
* [path](https://pkg.go.dev/path)
* [strings](https://pkg.go.dev/strings)
* [strconv](https://pkg.go.dev/strconv)
* [unicode](https://pkg.go.dev/unicode)
* [React.js](https://reactjs.org/)
* [axios](https://www.npmjs.com/package/)
* [express](https://www.npmjs.com/package/express)
* [Material-UI](https://www.npmjs.com/package/@material-ui/core/)
* [React Testing Library](https://www.npmjs.com/package/@testing-library/react)
* [moment](https://www.npmjs.com/package/moment)
* [eslint](https://www.npmjs.com/package/eslint)
* [filefy](https://www.npmjs.com/package/filefy)
* [classnames](https://www.npmjs.com/package/classnames)
* [mongoose](https://www.npmjs.com/package/mongoose)
* [cors](https://www.npmjs.com/package/cors)
* [classnames](https://www.npmjs.com/package/classnames)
* [emotion/react](https://www.npmjs.com/package/@emotion/react) etc.

## Local setup:
### `#1`
Move main.go to **$GOPATH/src/apsis/server/main.go** and building should work just fine

### `#2`
For debug or run operation:
##### `git clone https://github.com/TaylanYildirim/apsis.git`
##### `cd apsis/server/ && go get && go run main.go`

## api-doc

1.
> As a User\
> I want to be able to **create a new counter**\
> So that steps can be accumulated for a team of one or multiple employees

```http
POST https://apsis-code.herokuapp.com/team

```

2.
> As a User\
> I want to be able to increment the value of a stored counter\
> So that I can get steps counted towards my team's score

```http

GET https://apsis-code.herokuapp.com/teams/{id}/employee
e.g. https://apsis-code.herokuapp.com/teams/0/employee

```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Team id |


3.
> As a User\
> I want to get the current total steps taken by a team\
> So that I can see how much that team have walked in total
```http

GET https://apsis-code.herokuapp.com/team/{id}
e.g. https://apsis-code.herokuapp.com/team/0

```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Team id |

4.
> As a User\
> I want to list all teams and see their step counts\
> So that I can compare my team with the others
```http

GET https://apsis-code.herokuapp.com/teams

```

5.
> As a User\
> I want to list all counters in a team\
> So that I can see how much each team member have walked
```http

GET https://apsis-code.herokuapp.com/team/{teamId}/employees
e.g. https://apsis-code.herokuapp.com/team/0/employees

```
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Team id |

6.
> As a User\
> I want to be able to add/delete teams
> So that I can manage teams
```http

POST https://apsis-code.herokuapp.com/team

DELETE https://apsis-code.herokuapp.com/team/{id}
e.g. https://apsis-code.herokuapp.com/team/0

```
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Team id |
7.
> As a User\
> I want to be able to add/delete counters
> So that I can manage team member's counters
```http

POST https://apsis-code.herokuapp.com/teams/{id}/employees
e.g. https://apsis-code.herokuapp.com/teams/0/employees

DELETE https://apsis-code.herokuapp.com/teams/{id}/employees/{eID}
e.g. https://apsis-code.herokuapp.com/teams/0/employees/0

```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Team id |
| `eID` | `string` | **Required**. Employee id |
<br></br>
## Demo:




https://user-images.githubusercontent.com/18633675/149615357-57c3a222-72d9-485c-be14-6745c84f0ebb.mov



