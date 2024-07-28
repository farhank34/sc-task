package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)


type folderTestCase struct {
    OrgId string
    expected int
}

// The test cases are listed below

var folderTestCases = []folderTestCase{
    folderTestCase{"", 0}, 
    folderTestCase{"74392cf9-4724-4cf8-a45d-dd2b03033342", 0}, 
    folderTestCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 666},
    folderTestCase{"4212d618-66ff-468a-862d-ea49fef5e183", 1},    
    folderTestCase{"74392cf9-4724-4cf8-a45d-dd2b092f2b30", 1},    
    
}


type folderTestContentCase struct {
    OrgId string
    expectedNameInFirstRecord string
    expectedNameInLastRecord string
}

var folderTestContentCases = []folderTestContentCase{   
    folderTestContentCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", "current-excalibur","ultimate-pandemic"} ,  
    folderTestContentCase{"4212d618-66ff-468a-862d-ea49fef5e183", "heroic-bella","heroic-bella"}   ,
}


func TestGetAllFolders(t *testing.T) {
    t.Run("test", func(t *testing.T) {
        

// The following for loop goes through the test cases and checks how many times they appear in the dataset.

        for _, test := range folderTestCases {

            req := &folders.FetchFolderRequest{
                OrgID: uuid.FromStringOrNil(test.OrgId),
            }

            res, err := folders.GetAllFolders(req)      
            assert.Equal(t,err,nil,"GetAllFolders throughs a panic error")

            got := len(res.Folders)
            
            assert.Equal(t,test.expected,got,"Expected %d number of records but got %d",test.expected,got)

        }
//****************************************************************************************************************************


        // The following tests if the content inside the array is correct for some sample records (just the name e.g. "ultimate-pandemic")
        for _, test := range folderTestContentCases {

            req := &folders.FetchFolderRequest{
                OrgID: uuid.FromStringOrNil(test.OrgId),
            }

            res, err := folders.GetAllFolders(req)      
            assert.Equal(t,err,nil,"GetAllFolders throughs a panic error")

            got := len(res.Folders)
            FirstRecordName := res.Folders[0].Name
            LastRecordName := res.Folders[got-1].Name
            
            assert.Equal(t,test.expectedNameInFirstRecord,FirstRecordName,"Expected %s number of records but got %s",test.expectedNameInFirstRecord,FirstRecordName)
            assert.Equal(t,test.expectedNameInLastRecord,LastRecordName,"Expected %s number of records but got %s",test.expectedNameInLastRecord,LastRecordName)

        }
//****************************************************************************************************************************


    })
}

// Ref: https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package