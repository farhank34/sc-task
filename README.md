# sc-grad-2025

The technical take home for SC graduate program of 2025.

## Getting started

Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
```

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
  - within `folders.go`.
    - We would like you to read through, and run, the code.

     ### An attempt to run the main.go file
            When I attempted to run the code using the command go run main.go, I get the following errors:

                folders/folders.go:9:3: err declared and not used
                folders/folders.go:10:3: f1 declared and not used
                folders/folders.go:11:3: fs declared and not used
                folders/folders.go:15:6: k declared and not used
                folders/folders.go:19:6: k1 declared and not used

      ### The errors are due to unused variables in the file. I commented the variables to run the code as it is (minimal change)
      ### Here is the folder.go file after I commented out (or removed) the unused variables.

            func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
                /*
                var (
                    err error
                    f1  Folder
                    fs  []*Folder
                )*/
                f := []Folder{}
                r, _ := FetchAllFoldersByOrgID(req.OrgID)

                /* k is replace by _ */
                for _, v := range r {
                    f = append(f, *v)
                }
                var fp []*Folder
                
                /* k1 is replace by _ */
                for _, v1 := range f {
                    fp = append(fp, &v1)
                }
                var ffr *FetchFolderResponse
                ffr = &FetchFolderResponse{Folders: fp}
                return ffr, nil
            }
    
      ### There is a bug

          The code runs without any further errors, but the output shows only the same part of the data ("ultimate-pandemic") 999 times suggesting that there is a bug in the code.

      ### Looking for the bug

          The following two loops from the code seem to have no functional purpose in the system;
        
      ### 1st Loop
          for k, v := range r {
              f = append(f, *v)
          }

          -- This loop create a value array from pointer array
          
          
      ### 2nd loop
          var fp []*Folder
          for k1, v1 := range f {
              fp = append(fp, &v1)
          }


        -- This loop, creates a pointer array from the value array. 
        -- 'v1' is declared once in the loop and used (assigned) many times througout the loop. 
        -- Since the address of v1 is added to fp, fp contains only one address in the whole array (999 times same address)
        -- When we use fp for futher processing. The value point to the last content of the v1. It could be empty or something else at runtime if the garbage collector cleans it up or it gets allocated to something else.



    - Write some comments on what you think the code does.
        
        ### See folder.go file for comments on what I think the code does.

    - Suggest some improvements that can be made to the code.
    - Implement any suggested improvements.
    - Write up some unit tests in `folders_test.go` for your new `GetAllFolders` method

- Component 2:
  - Extend your improved code to now facilitate pagination.
  - You can copy over the existing methods into `folders_pagination.go` to get started.
  - Write a short explanation of your chosen solution.

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