# sc-grad-2025

The technical take home for SC graduate program of 2025.

## Getting started

Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
1. Execute: export SESSION_KEY="something in the command prompt.
2. Run go run main.go
3. Open your browser, and enter one of the following urls in the address bar:

 //  http://127.0.0.1:8001?OrgId=c1556e17-b7c0-45a3-a6ae-9546248fb17a 
 //  http://127.0.0.1:8001/?OrgId=c1556e17-b7c0-45a3-a6ae-9546248fb17a&token=fd6dc80e4645c0f46d08
 //  http://127.0.0.1:8001?OrgId=4212d618-66ff-468a-862d-ea49fef5e183

```

To run the unit tests in your local machine

1. Go to the folders' folder using the command: cd folders
2. Execute: go test

## Folder structure

```
| go.mod
| README.md
| sample.json
| main.go
| folders
    | folders.go
    | folders_test.go
    | static.go
```

## Instructions

- This technical assessment consists of 2 components:
- Component 1:
    All the modifications and comments are in part-1 branch

- Component 2:

    Install two libraries:
    go get github.com/gorilla/mux
    go get github.com/gorilla/sessions

    -- Session is used to store the token
    -- The session is also used to save the index of the latest page.
    -- Part 2 is implemented considering a REST base request call to the go application



## What is pagination?
  - Pagination helps break down a large dataset into smaller chunks.
  - Those smaller chunks can then be served to the client, and are usually accompanied by a token pointing to the next chunk.
  - The end result could potentially look like this:
```
  original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

  This might result in the following payload to the client:
  { data: [1, 2, 3, ..., 10] }

  However, with pagination implemented, the payload might instead look like this:
  request() -> { data: [1, 2], token: "nQsjz" }

  The token could then be used to fetch more results:

  request("nQsjz") -> { data : [3, 4], token: "uJsnQ" }

  .
  .
  .

  And more results until there's no data left:

  { data: [9, 10], token: "" }
```

## Submission

Create a repo in your chosen git repository (make sure it is public so we can access it) and reply with the link to your code. We recommend using GitHub.


## Contact

If you have any questions feel free to contact us at: graduates@safetyculture.io
