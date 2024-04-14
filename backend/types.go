package backend

type File struct {
	Name         string `json:"name"`
	IsDir        bool   `json:"isDir"`
	Url          string `json:"url"`
	LastModified int64  `json:"lastModified"`
	DateCreated  int64  `json:"createdAt"`
}

type FileResponse struct {
	Files []File `json:"files"`
	Next  string `json:"next"`
}
