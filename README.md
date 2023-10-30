# chirpy

A twitter clone built using Golang. To run, use go build --o && ./out and navigate to localhost:8080/app in your browser

## It currently supports the following APIs:

- User Create (POST localhost:8080/api/users)
  <br />
  JSON Body:
  {<br />
  "email": your email, <br />
  "password": your password" <br />
  }<br />
<br />
  Returns:
  {<br />
  "ID": your userID <br />
  "email": your email <br />
  "is_red_chirpy": false <br />
  }<br />
