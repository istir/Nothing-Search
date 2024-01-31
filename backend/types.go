package backend

type File struct {
	Name         string `json:"name"`
	IsDir        bool   `json:"isDir"`
	Url          string `json:"url"`
	LastModified int8   `json:"lastModified"`
}
