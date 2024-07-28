package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type PagedFolderResponse struct {
    Folders []*Folder   `json:"data"`
    Token string        `json:"token"`
}