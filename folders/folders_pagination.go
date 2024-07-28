package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

import (
	"encoding/hex"
	"log"
	"math/rand"
)

//const pageSize = 10


func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}




func GetFolders(firstPageNo int, pageSize int, req *FetchFolderRequest) (*PagedFolderResponse, error) {
    
    var ffr PagedFolderResponse 
    var fp []*Folder
    

/***********************************************************************************************************************/
/***********************************************************************************************************************/
/* Get all folder names from the json file and filter according to the OrgID                                           */
/*  req.OrgID - Pass the OrgID from the global req /static.go file                                                     */
/*  r is a pointer array with folder structure                                                                         */
    r, err := FetchAllFoldersByOrgID(req.OrgID)
    
    if err != nil {
        log.Println(err)
    }




    for pageNo := firstPageNo; pageNo < (firstPageNo+pageSize) && pageNo <len(r); pageNo++ {
        fp = append(fp, r[pageNo])
    }


    ffr = PagedFolderResponse{Folders: fp,Token: "set-before-delivery"}

     
    return &ffr, nil

}


