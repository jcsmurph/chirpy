# chirpy

A twitter clone built using Golang. To run, use go build --o && ./out and navigate to localhost:8080/app in your browser

# It currently supports the following APIs:

- User Create (POST localhost:8080/api/users)
  JSON Body:
  {
  "email": <your email>,
  "password": <your password>"
  }

  Returns:
  {
  "ID": <your userID>
  "email": <your email>
  "is_red_chirpy": false
  }
