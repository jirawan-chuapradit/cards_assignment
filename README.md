

# GoLang Application for Cards Service

This application provides essential APIs for managing cards.

## API Documentation

### 1. Get Cards Details

- **Description:** Retrieve details of a specific card.
- **Method:** GET
- **URL:** `http://localhost:8080/api/cards/:cardId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}

### 2. Create Card

- **Description:** Create a new card.
- **Method:** POST
- **URL:** `http://localhost:8080/api/cards`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
      "title": "title 4",
      "description": "mock data 4",
      "status": "TO DO"
  }

### 3. Update Card
- **Description:** Update an existing card.
- **Method:** PUT
- **URL:** `http://localhost:8080/api/cards/:cardId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
        "id" : "65dc37baf9f5187ef06653eb",
        "title": "modify title",
        "description": "mock data title",
        "status": "TO DO"
  }

### 4. Store Card
- **Description:** Archive a card.
- **Method:** PUT
- **URL:** `http://localhost:8080/api/cards/archive/:cardId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}

### 5. Get Cards
- **Description:** Retrieve all cards.
- **Method:** GET
- **URL:** `http://localhost:8080/api/cards`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
  
### 6. Get Cards History
- **Description:** Retrieve history of changes for a specific card.
- **Method:** GET
- **URL:** `http://localhost:8080/api/cards/history/:cardId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}

### 7. Delete Comment
- **Description:** Delete a comment.
- **Method:** DELETE
- **URL:** `http://localhost:8080/api/comments/:commentId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}

### 8. Update Comment
- **Description:** Update an existing comment.
- **Method:** PUT
- **URL:** `http://localhost:8080/api/comments/:commentId`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
        "id": "65dc37b5f9f5287ef06653ea",
        "description": "update comment by bob4_jones"
  }

### 9. Create Comment
- **Description:** Create a new comment.
- **Method:** PUT
- **URL:** `http://localhost:8080/api/comments`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
        "card_id": "65dc5e95018757f994149169",
        "description": "y Lorem Ipsum is simply dummy text of the printing and typesetting"
  }

### 10. Ping
- **Description:** Check if the service is up.
- **Method:** GET
- **URL:** `URL: http://localhost:8080/ping`

### 11. Test
- **Description:** Test the service.
- **Method:** GET
- **URL:** `http://localhost:8080/test`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}

### 12. Login
- **Description:** Login to the service.
- **Method:** POST
- **URL:** `http://localhost:8080/login`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
        "email": "bob4@example.com",
        "password": "hashed_password"
  }

### 13. Signup
- **Description:** Sign up for the service.
- **Method:** POST
- **URL:** `http://localhost:8080/signup`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}
- **Body:**
  ```json
  {
        "username": "alice3_jones",
        "email": "alice3@example.com",
        "password": "hashed_password"
  }

### 13. Logout
- **Description:** Logout from the service.
- **Method:** POST
- **URL:** `http://localhost:8080/logout`
- **Headers:**
  - `Authorization`: Bearer {{access_token}}



This markdown document provides an overview of the APIs provided by the GoLang application for managing cards. Each API includes details such as method, URL, headers, body (if applicable), and description.
