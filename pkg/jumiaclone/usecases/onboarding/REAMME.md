# ONBOARDING
Onboarding is the action or process of familiarizing a new customer or client with one's products or services.For this project this is where the customers get to interact with our system first.The main goal of this usecase is to allow the customer to sign up for the product incase they are not registered, sign in incase they exist and also expose the inventory we currently have so as to enable the customer to interact with the product in order to get a feel of what the product is actually about.

## SIGN UP
The sign up process for a user includes the following steps:
1. We use either a the user's email address or phone number to check if they already exists in the system. If they do they get to sign in/log in back to the system and if not they need to sign up.
2. They enter the required details such as names in order to register.
3. We send them a verification code via sms.
4. They enter the code and we verify that it's correct and mark them as verified hence completing the registration process.

## SING IN


# FEATURES

### 1. Create User

This feature capture every users information. These details are persisted to a local database.
### 1.1 API Specification
> **PS**: This feature exposes a `RESTFul` API.

**Authorization** Basic Auth

| Key          | Value           |
| ------------ | --------------- |
| Username     | "username"      |
| Password     | "password"      |

**POST** create user

    {{baseURL}}/api/v1/user

**Body** raw (json)
```json
    {
    "first_name":"<first name>",
    "middle_name":"<middle name>",
    "last_name":"<last name>",
    "phone_number": "<phone number>",
    "email": "<email>",
    "password": "<password>",
   }
```

**Response**
```json
    {
        "message": "contact has been created successfully",
        "status_code": 200,
        "body": {
            "name": "Dorothy Mayert",
            "phone_number": "+254700000004"
        }
    }
```