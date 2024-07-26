package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	/* var (
		err error
		f1  Folder
		fs  []*Folder
	) */

	f := []Folder{}

	/************************ What the code does ************************/

	/* The code gets all the folder names from the json file and filters them according to the OrgID */

	/* req.OrgID - Passes the OrgID from the global req /static.go file */

	/* r is a pointer array with the structure of a folder */

	/************************ What the code does ************************/


	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	/************************ What the code does ************************/

	/* The code loops through the folders, gets the locations/address of the folder and builds a new pointer array */

	/************************ What the code does ************************/

	for _, v := range r {

		fmt.Printf("\n\r contents inside the pointer *v = %v", *v) // '*v' provides the value inside the pointer 'v'
		fmt.Printf("\n\r memory location of v = %p", v) // This prints the actual location of the pointer

		f = append(f, *v)
	}

	// Now f got the actual values of the folders

	var fp []*Folder
	for _, v1 := range f {

		fmt.Printf("\n\r Address of an element of f, i.e. v -> &v1 = %p", &v1)
		fmt.Printf("\n\r Actual contents inside v1 = %v", v1)

		fp = append(fp, &v1)
	}

	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

/* 

Parameters: OrgID and UUID to denote the folder org OrgID

Returns:
	1. A pointer array with folder structure
	2. error

*/

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
