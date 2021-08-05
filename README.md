# Go Mail Server

A standalone server which its purpose is to only send email.

## Required Environment Variables

`Note:` Before running the application, environment variables must be defined.  
`Instruction:` Create a file `.env` inside your project directory, then add the variables with its corresponding values indicated in the table below.  
`Important Gmail Account Config:` You will need to turn on `Allow less secure apps` on your gmail account.  

| Variables      |                     Definition                      |          Stage Values |           Prod Values |
| :------------- | :-------------------------------------------------: | --------------------: | --------------------: |
| ENV            | Defines the environment the app will be running on. |                  prod |                  prod |
| PORT           |           The port where the app will run           |                    80 |                    80 |
| EMAIL          |     Email address used by app for sending mails     |          <your-email> |          <your-email> |
| EMAIL_PASSWORD |               Password for the email                | <your-email-password> | <your-email-password> |

## Build application

```
go build -o mail .
```

## Build and run application

```
Linux: go build -o mail . && ./mail
Windows: go build -o mail . && go run mail
```

## Directory structure

```
mail
    - config
        - bash
        - docker
    - src
        - config
        - gmail
        - routes
    - main.go
```

##### config

Main Project configuration

| Directory | Definition                                                                  |
| :-------- | :-------------------------------------------------------------------------- |
| bash      | Script that allows docker to have access to environment variables in gitlab |
| docker    | Docker configuration for stage and prod                                     |

##### src

Contains the project source code.

| Directory | Definition                                                                    |
| :-------- | :---------------------------------------------------------------------------- |
| config    | Contains configuration for getting required values from environment variables |
| gmail     | Contains the logic for sending email using gmail api                          |
| routes    | Contains api route of the application for sending email                       |

##### main.go

Entry point of the application.

##### API: Send email

- `https://yourdomain.com/api/send/mail`
  ```
  body: {
      email: string
      message: string
      subject: string
  }
  method: "POST"
  ```
  Response:
  ```
  {
    Message: string
    Email: string
    Subject: string
  }
  ```
