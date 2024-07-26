package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	


	/************************ What the code does ************************/

	/* The code gets all the folder names from the json file and filters them according to the OrgID */

	/* req.OrgID - Passes the OrgID from the global req /static.go file */

	/* r is a pointer array with the structure of a folder */

	/************************ What the code does ************************/


	r, err := FetchAllFoldersByOrgID(req.OrgID)

	/************************ What the code does ************************/

	/* The code loops through the folders, gets the locations/address of the folder and builds a new pointer array */

	/************************ What the code does ************************/

	

	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: r}
	return ffr, err
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
