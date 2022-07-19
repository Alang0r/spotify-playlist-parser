package lib

type File struct {
	Playlists []PlaylistElement `json:"playlists"`
}

type PlaylistElement struct {
	Name              string  `json:"name"`             
	LastModifiedDate  string  `json:"lastModifiedDate"` 
	Items             []Item  `json:"items"`            
	Description       *string `json:"description"`      
	NumberOfFollowers int64   `json:"numberOfFollowers"`
}

type Item struct {
	Track      *Track      `json:"track"`     
	Episode    interface{} `json:"episode"`   
	LocalTrack *LocalTrack `json:"localTrack"`
	AddedDate  string      `json:"addedDate"` 
}

type LocalTrack struct {
	URI string `json:"uri"`
}

type Track struct {
	TrackName  string `json:"trackName"` 
	ArtistName string `json:"artistName"`
	AlbumName  string `json:"albumName"` 
	TrackURI   string `json:"trackUri"`  
}
