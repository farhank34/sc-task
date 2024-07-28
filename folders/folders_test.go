package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)


type folderTestCase struct {
    OrgId string
    FirstPageNumber int
    PageSize  int
    Expected int
}

var folderTestCases = []folderTestCase{
    folderTestCase{"", 0,10,0}, 
    folderTestCase{"74392cf9-4724-4cf8-a45d-dd2b03033342", 0,10,0}, 
    folderTestCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 0,10,10},
    folderTestCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 600,10,10},
    folderTestCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 664,10,2},
    folderTestCase{"4212d618-66ff-468a-862d-ea49fef5e183", 0,10,1},    
    folderTestCase{"74392cf9-4724-4cf8-a45d-dd2b092f2b30", 0,10,1},    
    
}


type folderTestContentCase struct {
    OrgId string
    FirstPageNumber int
    PageSize  int
    expectedNameInFirstRecord string
    expectedNameInLastRecord string
}

var folderTestContentCases = []folderTestContentCase{   
    folderTestContentCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 0,10, "current-excalibur","honest-beach-head"} ,  
    folderTestContentCase{"c1556e17-b7c0-45a3-a6ae-9546248fb17a", 0,15, "current-excalibur","devoted-ricochet"} ,  
    folderTestContentCase{"4212d618-66ff-468a-862d-ea49fef5e183", 0,10, "heroic-bella","heroic-bella"}   ,
}


func TestGetAllFolders(t *testing.T) {
    t.Run("test", func(t *testing.T) {
        

// Test the number of records return by the function
        for _, test := range folderTestCases {

            req := &folders.FetchFolderRequest{
                OrgID: uuid.FromStringOrNil(test.OrgId),
            }
            
            //res, err := folders.GetAllFolders(req)        

            //Read a page of data to serve
            res, err := folders.GetAllFolders(test.FirstPageNumber,test.PageSize,req)

                if err != nil {
                    panic(err)
                }


            assert.Equal(t,err,nil,"GetAllFolders throughs a panic error")

            got := len(res.Folders)
            
            assert.Equal(t,test.Expected,got,"Expected %d number of records but got %d",test.Expected,got)

        }
//****************************************************************************************************************************


        // Test if the content inside the array is correct for some sample records
        
        for _, test := range folderTestContentCases {

            req := &folders.FetchFolderRequest{
                OrgID: uuid.FromStringOrNil(test.OrgId),
            }

            
            res, err := folders.GetAllFolders(test.FirstPageNumber,test.PageSize,req)
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
