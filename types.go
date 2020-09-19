package gobluos

import (
	"encoding/xml"
)

type StateResponse struct {
	XMLName xml.Name `xml:"state"`
	State   string   `xml:",chardata"`
}

type PlaybackResponse struct {
	XMLName xml.Name `xml:"id"`
	ID      int      `xml:",chardata"`
}

type VolumeResponse struct {
	XMLName  xml.Name `xml:"volume"`
	Etag     string   `xml:"etag,attr"`
	Db       float32  `xml:"db,attr"`
	Mute     bool     `xml:"mute,attr"`
	OffsetDb float32  `xml:"offsetDb,attr"`
	Level    int      `xml:",chardata"`
}

type StatusResponse struct {
	XMLName xml.Name `xml:"status"`
	Etag    string   `xml:"etag,attr"`
	Actions struct {
		Action []struct {
			Name string `xml:"name,attr"`
			URL  string `xml:"url,attr"`
		} `xml:"action"`
	} `xml:"actions"`
	Album           string  `xml:"album"`
	AlbumID         string  `xml:"albumid"`
	Artist          string  `xml:"artist"`
	ArtistID        string  `xml:"artistid"`
	Autofill        string  `xml:"autofill"`
	CanMovePlayback string  `xml:"canMovePlayback"`
	CanSeek         string  `xml:"canSeek"`
	CurrentImage    string  `xml:"currentImage"`
	Cursor          string  `xml:"cursor"`
	DB              string  `xml:"db"`
	Image           string  `xml:"image"`
	Indexing        string  `xml:"indexing"`
	MID             string  `xml:"mid"`
	Mode            string  `xml:"mode"`
	Mute            string  `xml:"mute"`
	PID             string  `xml:"pid"`
	PrID            string  `xml:"prid"`
	Quality         string  `xml:"quality"`
	Repeat          string  `xml:"repeat"`
	SchemaVersion   string  `xml:"schemaVersion"`
	Service         string  `xml:"service"`
	ServiceIcon     string  `xml:"serviceIcon"`
	ServiceName     string  `xml:"serviceName"`
	Shuffle         string  `xml:"shuffle"`
	SID             string  `xml:"sid"`
	Sleep           string  `xml:"sleep"`
	Song            string  `xml:"song"`
	SongID          string  `xml:"songid"`
	State           string  `xml:"state"`
	StreamFormat    string  `xml:"streamFormat"`
	StreamURL       string  `xml:"streamUrl"`
	SyncStat        string  `xml:"syncStat"`
	Title1          string  `xml:"title1"`
	Title2          string  `xml:"title2"`
	Title3          string  `xml:"title3"`
	TotalLength     float32 `xml:"totlen"`
	Volume          int     `xml:"volume"`
	Secs            int     `xml:"secs"`
}
