package seaweedfs

type WeedMasterResponse struct {
	Count     int    `json:"count"`
	Fid       string `json:"fid"`
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
}

type WeedVolumeResponse struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	ETag string `json:"eTag"`
}

type WeedVolumeLoopUpResponse struct {
	VolumeOrFileId string     `json:"volumeOrFileId"`
	Locations      []Location `json:"locations"`
}

type WeedUploadFileResponse struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Etag string `json:"eTag"`
}

type Location struct {
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
}
