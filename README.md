# chirpy

A twitter clone backend built using Golang. To run, use go build --o && ./out and navigate to localhost:8080/app in your browser.

To run certain APIs, you will need to generate your own JWT Secret to retrieve an access and refresh token. A fake polka secret can be created to run the API to upgrade a user to a paid subscription. Both tokens need to be added to a local .env file.  

## .env variables
- JWT_SECRET=
- POLKA_SECRET=

## Supported features
- Create a User
- Password hashing
- Login with a User
- Update a User
- Create a chirp
- Retrieve Chirps (with optional author and sort)
- Retrieve chirp by ID
- Delete a Chirp
- Upgrade a user with a fake payment webhook
- See site visits metric
- Reset site visits metric

## Supported APIs:

- User Create (POST localhost:8080/api/users)
```
  Body:  
  {  
  "email": your email,  
  "password": your password"  
  }  
  Returns:  
  {  
  "ID": your userID  
  "email": your email  
  "is_red_chirpy": false  
  }

```
- User Login (POST localhost:8080/api/login)  
  Body:  
  {  
  "email": your email,  
  "password": your password"  
  }  
  Returns:  
  {  
  "ID": your userID  
  "email": your email  
  "is_red_chirpy": false  
  "token": your ACCESS token  
  "refresh token": your refresh token  
  }  

- User Update (PUT localhost:8080/api/users)  
  Header:  
  {  
  "Authorization": Bearer your ACCESS token
  }  
  Returns:  
  {  
  "ID": your userID  
  "email": your email  
  "is_red_chirpy": false  
  }  

- User Upgrade (PUT localhost:8080/polka/webhooks)  
  Body:  
    {  
    "data":  
  {  
      "user_id": 1  
  },  
  "event": "user.payment_failed"  
}  
   
  Returns:  
  Status code 200  

- Chirp Create (POST localhost:8080/api/chirps)  
  Header:  
  {  
  "Authorization": Bearer your ACCESS token  
  }  
  Returns:  
  {  
  "ID": chirpID  
  "Body": the chirp  
  "AuthorID": AuthorID of the chirp  
  }  

- Chirp Retrieve all or by author ID (GET localhost:8080/api/chirps?author_id=?sort=?)  
  Header:  
  {  
  "Authorization": Bearer your ACCESS token  
  }  
  Returns:  
  {  
  "ID": chirpID  
  "Body": the chirp  
  "AuthorID": AuthorID of the chirp  
  }  

- Chirp Retrieve by ID (GET localhost:8080/api/chirps/{chirpID})  
  Header:  
  {  
  "Authorization": Bearer your ACCESS token  
  }   

  Returns:  
  {  
  "ID": chirpID  
  "Body": the chirp  
  "AuthorID": AuthorID of the chirp  
  }  

- Chirp Delete by ID (DELETE localhost:8080/api/chirps/{chirpID})  
  Header:  
  {  
  "Authorization": Bearer your ACCESS token  
  }  
  Returns:  
  Status code 200  

- Retrieve new ACCESS token (DELETE localhost:8080/api/refresh)  
  Header:  
  {  
  "Authorization": Bearer your REFRESH token  
  }  
  Returns:  
  {  
  "ID": chirpID  
  "Body": the chirp  
  "AuthorID": AuthorID of the chirp  
  }  
  
- Revoke REFRESH token (DELETE localhost:8080/api/revoke)  
  Header:  
  {  
  "Authorization": Bearer your REFRESH token  
  }  
  Returns:  
  Status code 200  

- Metrics (localhost:8080/api/healthz)  
  This API tracks the number of visits the site has had  

- Reset metrics (localhost:8080/api/reset)  
  This will reset the visits metric  
