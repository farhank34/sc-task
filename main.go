package main

//********** set this for development machine:  export SESSION_KEY="something"

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)


const pageSize = 10
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))



 // TEST Run in browser:
 //  http://127.0.0.1:8001?OrgId=c1556e17-b7c0-45a3-a6ae-9546248fb17a 
 //  http://127.0.0.1:8001/?OrgId=c1556e17-b7c0-45a3-a6ae-9546248fb17a&token=fd6dc80e4645c0f46d08
 //  http://127.0.0.1:8001?OrgId=4212d618-66ff-468a-862d-ea49fef5e183
  
 /*********************** The main function *****************************************************************************************/
/*                                                                                                                                  */
/*                                                                                                                                  */
 
  func main() {
    
    /***************** setup env for the web servers graceful shutdown *****************************************************************/
    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, 
                    "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()


    /************************ The web server startup code below *************************************************************************/
    srv := setupServer()

    /*********************** The following code is to stop the server gracefully on Ctrl-C  *********************************************/
    /*                                                                                                                                  */
    /*                                                                                                                                  */
    /*      Ref: https://github.com/gorilla/mux?tab=readme-ov-file#graceful-shutdown                                                    */
    
            c := make(chan os.Signal, 1)
            // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
            // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
            signal.Notify(c, os.Interrupt)

            // Block until we receive our signal.
            <-c

            // Create a deadline to wait for.
            ctx, cancel := context.WithTimeout(context.Background(), wait)
            defer cancel()
            // Doesn't block if no connections, but will otherwise wait
            // until the timeout deadline.
            srv.Shutdown(ctx)
            // Optionally, you could run srv.Shutdown in a goroutine and block on
            // <-ctx.Done() if your application should wait for other services
            // to finalize based on context cancellation.
            log.Println("shutting down")
            os.Exit(0)
    /*                                                                                                                                  */
    /*                                                                                                                                  */
    /************************************************************************************************************************************/      

  }


/*********************** This function creates a stand alone web server for development *********************************************/
/*                                                                                                                                  */
/*                                                                                                                                  */
func setupServer()(*http.Server) {

    httpRequestHandler := mux.NewRouter()
    httpRequestHandler.HandleFunc("/", REST_RequestHandler).Methods("GET")

    srv := &http.Server{
        Addr:         "127.0.0.1:8001",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: httpRequestHandler, // Pass our instance of gorilla/mux in.
    }


    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    
    return srv

  }
  

/*********************** This function handles the REST GET request for the default page ********************************************/
/*                                                                                                                                  */
/*                                                                                                                                  */
/*                                                                                                                                  */
/*   To keep the development effort at minimum, I presume that the REST API will use                                                 */
/*   HTTP Get method to get paginated output and token will be passed in query string                                               */
/*                                                                                                                                  */



  func REST_RequestHandler(w http.ResponseWriter, r *http.Request) {    
    
    /**************** Section variable declaration and initialization *********************/
    var err error   
    firstPageNo := 0

    OrgIdSupplied :=  r.URL.Query().Get("OrgId")
    fmt.Printf("\n\rOrgIdSupplied = %v",OrgIdSupplied)


    if(OrgIdSupplied != ""){
        

            // Read the token for the GET request query string
            tokenSupplied :=  r.URL.Query().Get("token")
            tokenStoredInSession := "" 
            
            // Get the current session for the session store
            session, errSessionGet := store.Get(r, "session-for-folder-project")
                    if errSessionGet != nil {
                        http.Error(w, errSessionGet.Error(), http.StatusInternalServerError)
                        return
                    }


            // If a token is stored in the session, then get it otherwise leave as empty as set on initialization block
            if(session.Values["token"] != ""){ tokenStoredInSession = fmt.Sprint(session.Values["token"]) }

            //check if the stored token is same as the supplied token (by query string or REST GET Method)
            if tokenStoredInSession == tokenSupplied {
                //fmt.Printf("\n\r Token matched\n\r")
                //A valid token is supplied so read the last page number served from the session store
                nextPageStartFrom := fmt.Sprint(session.Values["nextPageStartFrom"])
                //Convert to an integer for calculation
                firstPageNo, err = strconv.Atoi(nextPageStartFrom)
                    if (err != nil) {
                        firstPageNo  = 0  // On error start from first page      
                    }

            } else {
                //fmt.Printf("\n\r Token mismatched; so start from 1st page\n\r")
                // Start from beginning if the token is not supplied or is invalid
                firstPageNo  = 0  
            }


            
            
            // Prepare the backend to deliver next page by saving the page number
            nextPageNo :=  firstPageNo + pageSize
            session.Values["nextPageStartFrom"] = nextPageNo    

            
            
            // Create and save a new token for next page
            token_new := folders.GenerateSecureToken(10)
            session.Values["token"] = token_new


            
            filterOrg := &folders.FetchFolderRequest{
                OrgID: uuid.FromStringOrNil(OrgIdSupplied),
            }
            
        /************************************** Start writing the output buffer *****************************************************/
            // Save the session variables before we write to the response/return from the handler.
            err = session.Save(r, w)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Tell the backend to server json data
            w.Header().Set("Content-Type", "application/json")

                

            //Read a page of data to serve
            res, err := folders.GetFolders(firstPageNo,pageSize,filterOrg)

                if err != nil {
                    panic(err)
                }

            // Add the token to the json output    
            res.Token = fmt.Sprint(token_new)
            

            jsonBytes, err := json.Marshal(res)
                if err != nil {
                    panic(err)
                }

            //Finally write the content into the output buffer  
            w.Write([]byte(string(jsonBytes)))
    
    } else {
        w.Write([]byte("You must supply an OrgId in query string"))
    }

    
  }