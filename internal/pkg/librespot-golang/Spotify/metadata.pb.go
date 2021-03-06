// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Album_Type int32

const (
	Album_ALBUM       Album_Type = 1
	Album_SINGLE      Album_Type = 2
	Album_COMPILATION Album_Type = 3
	Album_EP          Album_Type = 4
)

var Album_Type_name = map[int32]string{
	1: "ALBUM",
	2: "SINGLE",
	3: "COMPILATION",
	4: "EP",
}
var Album_Type_value = map[string]int32{
	"ALBUM":       1,
	"SINGLE":      2,
	"COMPILATION": 3,
	"EP":          4,
}

func (x Album_Type) Enum() *Album_Type {
	p := new(Album_Type)
	*p = x
	return p
}
func (x Album_Type) String() string {
	return proto.EnumName(Album_Type_name, int32(x))
}
func (x *Album_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Album_Type_value, data, "Album_Type")
	if err != nil {
		return err
	}
	*x = Album_Type(value)
	return nil
}
func (Album_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{5, 0} }

type Image_Size int32

const (
	Image_DEFAULT Image_Size = 0
	Image_SMALL   Image_Size = 1
	Image_LARGE   Image_Size = 2
	Image_XLARGE  Image_Size = 3
)

var Image_Size_name = map[int32]string{
	0: "DEFAULT",
	1: "SMALL",
	2: "LARGE",
	3: "XLARGE",
}
var Image_Size_value = map[string]int32{
	"DEFAULT": 0,
	"SMALL":   1,
	"LARGE":   2,
	"XLARGE":  3,
}

func (x Image_Size) Enum() *Image_Size {
	p := new(Image_Size)
	*p = x
	return p
}
func (x Image_Size) String() string {
	return proto.EnumName(Image_Size_name, int32(x))
}
func (x *Image_Size) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Image_Size_value, data, "Image_Size")
	if err != nil {
		return err
	}
	*x = Image_Size(value)
	return nil
}
func (Image_Size) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{7, 0} }

type Copyright_Type int32

const (
	Copyright_P Copyright_Type = 0
	Copyright_C Copyright_Type = 1
)

var Copyright_Type_name = map[int32]string{
	0: "P",
	1: "C",
}
var Copyright_Type_value = map[string]int32{
	"P": 0,
	"C": 1,
}

func (x Copyright_Type) Enum() *Copyright_Type {
	p := new(Copyright_Type)
	*p = x
	return p
}
func (x Copyright_Type) String() string {
	return proto.EnumName(Copyright_Type_name, int32(x))
}
func (x *Copyright_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Copyright_Type_value, data, "Copyright_Type")
	if err != nil {
		return err
	}
	*x = Copyright_Type(value)
	return nil
}
func (Copyright_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{11, 0} }

type Restriction_Type int32

const (
	Restriction_STREAMING Restriction_Type = 0
)

var Restriction_Type_name = map[int32]string{
	0: "STREAMING",
}
var Restriction_Type_value = map[string]int32{
	"STREAMING": 0,
}

func (x Restriction_Type) Enum() *Restriction_Type {
	p := new(Restriction_Type)
	*p = x
	return p
}
func (x Restriction_Type) String() string {
	return proto.EnumName(Restriction_Type_name, int32(x))
}
func (x *Restriction_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Restriction_Type_value, data, "Restriction_Type")
	if err != nil {
		return err
	}
	*x = Restriction_Type(value)
	return nil
}
func (Restriction_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{12, 0} }

type AudioFile_Format int32

const (
	AudioFile_OGG_VORBIS_96  AudioFile_Format = 0
	AudioFile_OGG_VORBIS_160 AudioFile_Format = 1
	AudioFile_OGG_VORBIS_320 AudioFile_Format = 2
	AudioFile_MP3_256        AudioFile_Format = 3
	AudioFile_MP3_320        AudioFile_Format = 4
	AudioFile_MP3_160        AudioFile_Format = 5
	AudioFile_MP3_96         AudioFile_Format = 6
	AudioFile_MP3_160_ENC    AudioFile_Format = 7
	AudioFile_OTHER2         AudioFile_Format = 8
	AudioFile_OTHER3         AudioFile_Format = 9
	AudioFile_AAC_160        AudioFile_Format = 10
	AudioFile_AAC_320        AudioFile_Format = 11
	AudioFile_OTHER4         AudioFile_Format = 12
	AudioFile_OTHER5         AudioFile_Format = 13
)

var AudioFile_Format_name = map[int32]string{
	0:  "OGG_VORBIS_96",
	1:  "OGG_VORBIS_160",
	2:  "OGG_VORBIS_320",
	3:  "MP3_256",
	4:  "MP3_320",
	5:  "MP3_160",
	6:  "MP3_96",
	7:  "MP3_160_ENC",
	8:  "OTHER2",
	9:  "OTHER3",
	10: "AAC_160",
	11: "AAC_320",
	12: "OTHER4",
	13: "OTHER5",
}
var AudioFile_Format_value = map[string]int32{
	"OGG_VORBIS_96":  0,
	"OGG_VORBIS_160": 1,
	"OGG_VORBIS_320": 2,
	"MP3_256":        3,
	"MP3_320":        4,
	"MP3_160":        5,
	"MP3_96":         6,
	"MP3_160_ENC":    7,
	"OTHER2":         8,
	"OTHER3":         9,
	"AAC_160":        10,
	"AAC_320":        11,
	"OTHER4":         12,
	"OTHER5":         13,
}

func (x AudioFile_Format) Enum() *AudioFile_Format {
	p := new(AudioFile_Format)
	*p = x
	return p
}
func (x AudioFile_Format) String() string {
	return proto.EnumName(AudioFile_Format_name, int32(x))
}
func (x *AudioFile_Format) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AudioFile_Format_value, data, "AudioFile_Format")
	if err != nil {
		return err
	}
	*x = AudioFile_Format(value)
	return nil
}
func (AudioFile_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{15, 0} }

type TopTracks struct {
	Country          *string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	Track            []*Track `protobuf:"bytes,2,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TopTracks) Reset()                    { *m = TopTracks{} }
func (m *TopTracks) String() string            { return proto.CompactTextString(m) }
func (*TopTracks) ProtoMessage()               {}
func (*TopTracks) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TopTracks) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *TopTracks) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type ActivityPeriod struct {
	StartYear        *int32 `protobuf:"zigzag32,1,opt,name=start_year,json=startYear" json:"start_year,omitempty"`
	EndYear          *int32 `protobuf:"zigzag32,2,opt,name=end_year,json=endYear" json:"end_year,omitempty"`
	Decade           *int32 `protobuf:"zigzag32,3,opt,name=decade" json:"decade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActivityPeriod) Reset()                    { *m = ActivityPeriod{} }
func (m *ActivityPeriod) String() string            { return proto.CompactTextString(m) }
func (*ActivityPeriod) ProtoMessage()               {}
func (*ActivityPeriod) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *ActivityPeriod) GetStartYear() int32 {
	if m != nil && m.StartYear != nil {
		return *m.StartYear
	}
	return 0
}

func (m *ActivityPeriod) GetEndYear() int32 {
	if m != nil && m.EndYear != nil {
		return *m.EndYear
	}
	return 0
}

func (m *ActivityPeriod) GetDecade() int32 {
	if m != nil && m.Decade != nil {
		return *m.Decade
	}
	return 0
}

type Artist struct {
	Gid                  []byte            `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name                 *string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Popularity           *int32            `protobuf:"zigzag32,3,opt,name=popularity" json:"popularity,omitempty"`
	TopTrack             []*TopTracks      `protobuf:"bytes,4,rep,name=top_track,json=topTrack" json:"top_track,omitempty"`
	AlbumGroup           []*AlbumGroup     `protobuf:"bytes,5,rep,name=album_group,json=albumGroup" json:"album_group,omitempty"`
	SingleGroup          []*AlbumGroup     `protobuf:"bytes,6,rep,name=single_group,json=singleGroup" json:"single_group,omitempty"`
	CompilationGroup     []*AlbumGroup     `protobuf:"bytes,7,rep,name=compilation_group,json=compilationGroup" json:"compilation_group,omitempty"`
	AppearsOnGroup       []*AlbumGroup     `protobuf:"bytes,8,rep,name=appears_on_group,json=appearsOnGroup" json:"appears_on_group,omitempty"`
	Genre                []string          `protobuf:"bytes,9,rep,name=genre" json:"genre,omitempty"`
	ExternalId           []*ExternalId     `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Portrait             []*Image          `protobuf:"bytes,11,rep,name=portrait" json:"portrait,omitempty"`
	Biography            []*Biography      `protobuf:"bytes,12,rep,name=biography" json:"biography,omitempty"`
	ActivityPeriod       []*ActivityPeriod `protobuf:"bytes,13,rep,name=activity_period,json=activityPeriod" json:"activity_period,omitempty"`
	Restriction          []*Restriction    `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related              []*Artist         `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	IsPortraitAlbumCover *bool             `protobuf:"varint,16,opt,name=is_portrait_album_cover,json=isPortraitAlbumCover" json:"is_portrait_album_cover,omitempty"`
	PortraitGroup        *ImageGroup       `protobuf:"bytes,17,opt,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized     []byte            `json:"-"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Artist) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Artist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artist) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Artist) GetTopTrack() []*TopTracks {
	if m != nil {
		return m.TopTrack
	}
	return nil
}

func (m *Artist) GetAlbumGroup() []*AlbumGroup {
	if m != nil {
		return m.AlbumGroup
	}
	return nil
}

func (m *Artist) GetSingleGroup() []*AlbumGroup {
	if m != nil {
		return m.SingleGroup
	}
	return nil
}

func (m *Artist) GetCompilationGroup() []*AlbumGroup {
	if m != nil {
		return m.CompilationGroup
	}
	return nil
}

func (m *Artist) GetAppearsOnGroup() []*AlbumGroup {
	if m != nil {
		return m.AppearsOnGroup
	}
	return nil
}

func (m *Artist) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Artist) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Artist) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Artist) GetBiography() []*Biography {
	if m != nil {
		return m.Biography
	}
	return nil
}

func (m *Artist) GetActivityPeriod() []*ActivityPeriod {
	if m != nil {
		return m.ActivityPeriod
	}
	return nil
}

func (m *Artist) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Artist) GetRelated() []*Artist {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Artist) GetIsPortraitAlbumCover() bool {
	if m != nil && m.IsPortraitAlbumCover != nil {
		return *m.IsPortraitAlbumCover
	}
	return false
}

func (m *Artist) GetPortraitGroup() *ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type AlbumGroup struct {
	Album            []*Album `protobuf:"bytes,1,rep,name=album" json:"album,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AlbumGroup) Reset()                    { *m = AlbumGroup{} }
func (m *AlbumGroup) String() string            { return proto.CompactTextString(m) }
func (*AlbumGroup) ProtoMessage()               {}
func (*AlbumGroup) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *AlbumGroup) GetAlbum() []*Album {
	if m != nil {
		return m.Album
	}
	return nil
}

type Date struct {
	Year             *int32 `protobuf:"zigzag32,1,opt,name=year" json:"year,omitempty"`
	Month            *int32 `protobuf:"zigzag32,2,opt,name=month" json:"month,omitempty"`
	Day              *int32 `protobuf:"zigzag32,3,opt,name=day" json:"day,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *Date) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

type Album struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,3,rep,name=artist" json:"artist,omitempty"`
	Typ              *Album_Type    `protobuf:"varint,4,opt,name=typ,enum=Spotify.Album_Type" json:"typ,omitempty"`
	Label            *string        `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Date             *Date          `protobuf:"bytes,6,opt,name=date" json:"date,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,7,opt,name=popularity" json:"popularity,omitempty"`
	Genre            []string       `protobuf:"bytes,8,rep,name=genre" json:"genre,omitempty"`
	Cover            []*Image       `protobuf:"bytes,9,rep,name=cover" json:"cover,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Disc             []*Disc        `protobuf:"bytes,11,rep,name=disc" json:"disc,omitempty"`
	Review           []string       `protobuf:"bytes,12,rep,name=review" json:"review,omitempty"`
	Copyright        []*Copyright   `protobuf:"bytes,13,rep,name=copyright" json:"copyright,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related          []*Album       `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,16,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
	CoverGroup       *ImageGroup    `protobuf:"bytes,17,opt,name=cover_group,json=coverGroup" json:"cover_group,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Album) Reset()                    { *m = Album{} }
func (m *Album) String() string            { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()               {}
func (*Album) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *Album) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Album) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Album) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Album) GetTyp() Album_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Album_ALBUM
}

func (m *Album) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Album) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Album) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Album) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Album) GetCover() []*Image {
	if m != nil {
		return m.Cover
	}
	return nil
}

func (m *Album) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Album) GetDisc() []*Disc {
	if m != nil {
		return m.Disc
	}
	return nil
}

func (m *Album) GetReview() []string {
	if m != nil {
		return m.Review
	}
	return nil
}

func (m *Album) GetCopyright() []*Copyright {
	if m != nil {
		return m.Copyright
	}
	return nil
}

func (m *Album) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Album) GetRelated() []*Album {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Album) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Album) GetCoverGroup() *ImageGroup {
	if m != nil {
		return m.CoverGroup
	}
	return nil
}

type Track struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Album            *Album         `protobuf:"bytes,3,opt,name=album" json:"album,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,4,rep,name=artist" json:"artist,omitempty"`
	Number           *int32         `protobuf:"zigzag32,5,opt,name=number" json:"number,omitempty"`
	DiscNumber       *int32         `protobuf:"zigzag32,6,opt,name=disc_number,json=discNumber" json:"disc_number,omitempty"`
	Duration         *int32         `protobuf:"zigzag32,7,opt,name=duration" json:"duration,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,8,opt,name=popularity" json:"popularity,omitempty"`
	Explicit         *bool          `protobuf:"varint,9,opt,name=explicit" json:"explicit,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,11,rep,name=restriction" json:"restriction,omitempty"`
	File             []*AudioFile   `protobuf:"bytes,12,rep,name=file" json:"file,omitempty"`
	Alternative      []*Track       `protobuf:"bytes,13,rep,name=alternative" json:"alternative,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,14,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
	Preview          []*AudioFile   `protobuf:"bytes,15,rep,name=preview" json:"preview,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}
func (*Track) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *Track) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Track) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Track) GetAlbum() *Album {
	if m != nil {
		return m.Album
	}
	return nil
}

func (m *Track) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Track) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Track) GetDiscNumber() int32 {
	if m != nil && m.DiscNumber != nil {
		return *m.DiscNumber
	}
	return 0
}

func (m *Track) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Track) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Track) GetExplicit() bool {
	if m != nil && m.Explicit != nil {
		return *m.Explicit
	}
	return false
}

func (m *Track) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Track) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Track) GetFile() []*AudioFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Track) GetAlternative() []*Track {
	if m != nil {
		return m.Alternative
	}
	return nil
}

func (m *Track) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Track) GetPreview() []*AudioFile {
	if m != nil {
		return m.Preview
	}
	return nil
}

type Image struct {
	FileId           []byte      `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Size             *Image_Size `protobuf:"varint,2,opt,name=size,enum=Spotify.Image_Size" json:"size,omitempty"`
	Width            *int32      `protobuf:"zigzag32,3,opt,name=width" json:"width,omitempty"`
	Height           *int32      `protobuf:"zigzag32,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *Image) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *Image) GetSize() Image_Size {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return Image_DEFAULT
}

func (m *Image) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type ImageGroup struct {
	Image            []*Image `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImageGroup) Reset()                    { *m = ImageGroup{} }
func (m *ImageGroup) String() string            { return proto.CompactTextString(m) }
func (*ImageGroup) ProtoMessage()               {}
func (*ImageGroup) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *ImageGroup) GetImage() []*Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Biography struct {
	Text             *string       `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Portrait         []*Image      `protobuf:"bytes,2,rep,name=portrait" json:"portrait,omitempty"`
	PortraitGroup    []*ImageGroup `protobuf:"bytes,3,rep,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Biography) Reset()                    { *m = Biography{} }
func (m *Biography) String() string            { return proto.CompactTextString(m) }
func (*Biography) ProtoMessage()               {}
func (*Biography) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *Biography) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Biography) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Biography) GetPortraitGroup() []*ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type Disc struct {
	Number           *int32   `protobuf:"zigzag32,1,opt,name=number" json:"number,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Track            []*Track `protobuf:"bytes,3,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Disc) Reset()                    { *m = Disc{} }
func (m *Disc) String() string            { return proto.CompactTextString(m) }
func (*Disc) ProtoMessage()               {}
func (*Disc) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *Disc) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Disc) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Disc) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type Copyright struct {
	Typ              *Copyright_Type `protobuf:"varint,1,opt,name=typ,enum=Spotify.Copyright_Type" json:"typ,omitempty"`
	Text             *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Copyright) Reset()                    { *m = Copyright{} }
func (m *Copyright) String() string            { return proto.CompactTextString(m) }
func (*Copyright) ProtoMessage()               {}
func (*Copyright) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *Copyright) GetTyp() Copyright_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Copyright_P
}

func (m *Copyright) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type Restriction struct {
	CountriesAllowed   *string           `protobuf:"bytes,2,opt,name=countries_allowed,json=countriesAllowed" json:"countries_allowed,omitempty"`
	CountriesForbidden *string           `protobuf:"bytes,3,opt,name=countries_forbidden,json=countriesForbidden" json:"countries_forbidden,omitempty"`
	Typ                *Restriction_Type `protobuf:"varint,4,opt,name=typ,enum=Spotify.Restriction_Type" json:"typ,omitempty"`
	CatalogueStr       []string          `protobuf:"bytes,5,rep,name=catalogue_str,json=catalogueStr" json:"catalogue_str,omitempty"`
	XXX_unrecognized   []byte            `json:"-"`
}

func (m *Restriction) Reset()                    { *m = Restriction{} }
func (m *Restriction) String() string            { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()               {}
func (*Restriction) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{12} }

func (m *Restriction) GetCountriesAllowed() string {
	if m != nil && m.CountriesAllowed != nil {
		return *m.CountriesAllowed
	}
	return ""
}

func (m *Restriction) GetCountriesForbidden() string {
	if m != nil && m.CountriesForbidden != nil {
		return *m.CountriesForbidden
	}
	return ""
}

func (m *Restriction) GetTyp() Restriction_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Restriction_STREAMING
}

func (m *Restriction) GetCatalogueStr() []string {
	if m != nil {
		return m.CatalogueStr
	}
	return nil
}

type SalePeriod struct {
	Restriction      []*Restriction `protobuf:"bytes,1,rep,name=restriction" json:"restriction,omitempty"`
	Start            *Date          `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End              *Date          `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SalePeriod) Reset()                    { *m = SalePeriod{} }
func (m *SalePeriod) String() string            { return proto.CompactTextString(m) }
func (*SalePeriod) ProtoMessage()               {}
func (*SalePeriod) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{13} }

func (m *SalePeriod) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *SalePeriod) GetStart() *Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SalePeriod) GetEnd() *Date {
	if m != nil {
		return m.End
	}
	return nil
}

type ExternalId struct {
	Typ              *string `protobuf:"bytes,1,opt,name=typ" json:"typ,omitempty"`
	Id               *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalId) Reset()                    { *m = ExternalId{} }
func (m *ExternalId) String() string            { return proto.CompactTextString(m) }
func (*ExternalId) ProtoMessage()               {}
func (*ExternalId) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{14} }

func (m *ExternalId) GetTyp() string {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return ""
}

func (m *ExternalId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type AudioFile struct {
	FileId           []byte            `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Format           *AudioFile_Format `protobuf:"varint,2,opt,name=format,enum=Spotify.AudioFile_Format" json:"format,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *AudioFile) Reset()                    { *m = AudioFile{} }
func (m *AudioFile) String() string            { return proto.CompactTextString(m) }
func (*AudioFile) ProtoMessage()               {}
func (*AudioFile) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{15} }

func (m *AudioFile) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *AudioFile) GetFormat() AudioFile_Format {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return AudioFile_OGG_VORBIS_96
}

func init() {
	proto.RegisterType((*TopTracks)(nil), "Spotify.TopTracks")
	proto.RegisterType((*ActivityPeriod)(nil), "Spotify.ActivityPeriod")
	proto.RegisterType((*Artist)(nil), "Spotify.Artist")
	proto.RegisterType((*AlbumGroup)(nil), "Spotify.AlbumGroup")
	proto.RegisterType((*Date)(nil), "Spotify.Date")
	proto.RegisterType((*Album)(nil), "Spotify.Album")
	proto.RegisterType((*Track)(nil), "Spotify.Track")
	proto.RegisterType((*Image)(nil), "Spotify.Image")
	proto.RegisterType((*ImageGroup)(nil), "Spotify.ImageGroup")
	proto.RegisterType((*Biography)(nil), "Spotify.Biography")
	proto.RegisterType((*Disc)(nil), "Spotify.Disc")
	proto.RegisterType((*Copyright)(nil), "Spotify.Copyright")
	proto.RegisterType((*Restriction)(nil), "Spotify.Restriction")
	proto.RegisterType((*SalePeriod)(nil), "Spotify.SalePeriod")
	proto.RegisterType((*ExternalId)(nil), "Spotify.ExternalId")
	proto.RegisterType((*AudioFile)(nil), "Spotify.AudioFile")
	proto.RegisterEnum("Spotify.Album_Type", Album_Type_name, Album_Type_value)
	proto.RegisterEnum("Spotify.Image_Size", Image_Size_name, Image_Size_value)
	proto.RegisterEnum("Spotify.Copyright_Type", Copyright_Type_name, Copyright_Type_value)
	proto.RegisterEnum("Spotify.Restriction_Type", Restriction_Type_name, Restriction_Type_value)
	proto.RegisterEnum("Spotify.AudioFile_Format", AudioFile_Format_name, AudioFile_Format_value)
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 1430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x6d, 0x6f, 0xdb, 0x36,
	0x10, 0x8e, 0x6c, 0x59, 0xb6, 0xce, 0x89, 0xa3, 0xb0, 0x59, 0xa3, 0x16, 0xd8, 0x9a, 0xa9, 0xdd,
	0x9a, 0xae, 0x43, 0x9a, 0xa6, 0x4d, 0x80, 0x0e, 0x18, 0x50, 0x27, 0x4d, 0x32, 0x63, 0x79, 0x03,
	0xed, 0x0e, 0xdd, 0x27, 0x8d, 0xb1, 0x18, 0x87, 0x98, 0x6c, 0x09, 0x12, 0x9d, 0xd6, 0xfd, 0x03,
	0x03, 0xf6, 0x57, 0x06, 0x0c, 0xfb, 0x3a, 0xec, 0x57, 0x0c, 0xd8, 0x0f, 0x1a, 0xf8, 0x22, 0xd9,
	0x91, 0xe3, 0x36, 0xe8, 0xa7, 0xf0, 0xf8, 0xdc, 0x9d, 0xc5, 0xe3, 0xf3, 0xdc, 0x31, 0xd0, 0xe8,
	0x53, 0x4e, 0x02, 0xc2, 0xc9, 0x7a, 0x9c, 0x44, 0x3c, 0x42, 0xd5, 0x76, 0x1c, 0x71, 0x76, 0x3e,
	0xf2, 0x7e, 0x04, 0xbb, 0x13, 0xc5, 0x9d, 0x84, 0x74, 0x7f, 0x4d, 0x91, 0x0b, 0xd5, 0x6e, 0x34,
	0x1c, 0xf0, 0x64, 0xe4, 0x1a, 0xab, 0xc6, 0x9a, 0x8d, 0x33, 0x13, 0x3d, 0x80, 0x0a, 0x17, 0x3e,
	0x6e, 0x69, 0xb5, 0xbc, 0x56, 0xdf, 0x6c, 0xac, 0xeb, 0xf8, 0x75, 0x19, 0x89, 0x15, 0xe8, 0x9d,
	0x41, 0xa3, 0xd9, 0xe5, 0xec, 0x92, 0xf1, 0xd1, 0x29, 0x4d, 0x58, 0x14, 0xa0, 0xcf, 0x01, 0x52,
	0x4e, 0x12, 0xee, 0x8f, 0x28, 0x49, 0x64, 0xd2, 0x25, 0x6c, 0xcb, 0x9d, 0x9f, 0x29, 0x49, 0xd0,
	0x1d, 0xa8, 0xd1, 0x41, 0xa0, 0xc0, 0x92, 0x04, 0xab, 0x74, 0x10, 0x48, 0xe8, 0x36, 0x58, 0x01,
	0xed, 0x92, 0x80, 0xba, 0x65, 0x09, 0x68, 0xcb, 0xfb, 0xcb, 0x02, 0xab, 0x99, 0x70, 0x96, 0x72,
	0xe4, 0x40, 0xb9, 0xc7, 0x02, 0x99, 0x75, 0x1e, 0x8b, 0x25, 0x42, 0x60, 0x0e, 0x48, 0x9f, 0xca,
	0x5c, 0x36, 0x96, 0x6b, 0xf4, 0x05, 0x40, 0x1c, 0xc5, 0xc3, 0x90, 0x24, 0x8c, 0x8f, 0x74, 0xb2,
	0x89, 0x1d, 0xf4, 0x04, 0x6c, 0x1e, 0xc5, 0xbe, 0x3a, 0x9e, 0x29, 0x8f, 0x87, 0xc6, 0xc7, 0xcb,
	0x6a, 0x83, 0x6b, 0x5c, 0x2f, 0xd1, 0x73, 0xa8, 0x93, 0xf0, 0x6c, 0xd8, 0xf7, 0x7b, 0x49, 0x34,
	0x8c, 0xdd, 0x8a, 0x0c, 0xb9, 0x95, 0x87, 0x34, 0x05, 0x76, 0x20, 0x20, 0x0c, 0x24, 0x5f, 0xa3,
	0x6d, 0x98, 0x4f, 0xd9, 0xa0, 0x17, 0x52, 0x1d, 0x66, 0xcd, 0x0e, 0xab, 0x2b, 0x47, 0x15, 0xf7,
	0x12, 0x96, 0xba, 0x51, 0x3f, 0x66, 0x21, 0xe1, 0x2c, 0x1a, 0xe8, 0xe0, 0xea, 0xec, 0x60, 0x67,
	0xc2, 0x5b, 0x65, 0xf8, 0x1e, 0x1c, 0x12, 0xc7, 0x94, 0x24, 0xa9, 0x9f, 0x27, 0xa8, 0xcd, 0x4e,
	0xd0, 0xd0, 0xce, 0x27, 0x3a, 0x7c, 0x19, 0x2a, 0x3d, 0x3a, 0x48, 0xa8, 0x6b, 0xaf, 0x96, 0xd7,
	0x6c, 0xac, 0x0c, 0x51, 0x04, 0xfa, 0x8e, 0xd3, 0x64, 0x40, 0x42, 0x9f, 0x05, 0x2e, 0x14, 0xf2,
	0xed, 0x69, 0xac, 0x15, 0x60, 0xa0, 0xf9, 0x1a, 0x7d, 0x03, 0xb5, 0x38, 0x4a, 0x78, 0x42, 0x18,
	0x77, 0xeb, 0x05, 0x26, 0xb5, 0xfa, 0xa4, 0x47, 0x71, 0x8e, 0xa3, 0x0d, 0xb0, 0xcf, 0x58, 0xd4,
	0x4b, 0x48, 0x7c, 0x31, 0x72, 0xe7, 0x0b, 0xf7, 0xb2, 0x93, 0x21, 0x78, 0xec, 0x84, 0x5e, 0xc2,
	0x22, 0xd1, 0xf4, 0xf3, 0x63, 0xc9, 0x3f, 0x77, 0x41, 0xc6, 0xad, 0x8c, 0xcf, 0x79, 0x85, 0x9e,
	0xb8, 0x41, 0xae, 0xd2, 0x75, 0x1b, 0xea, 0x09, 0x4d, 0x79, 0xc2, 0xba, 0xa2, 0x7c, 0x6e, 0x43,
	0x46, 0x2f, 0xe7, 0xd1, 0x78, 0x8c, 0xe1, 0x49, 0x47, 0xf4, 0x08, 0xaa, 0x09, 0x0d, 0x09, 0xa7,
	0x81, 0xbb, 0x28, 0x63, 0x16, 0xc7, 0xbf, 0x28, 0xb9, 0x8a, 0x33, 0x1c, 0x6d, 0xc1, 0x0a, 0x4b,
	0xfd, 0xec, 0x94, 0xbe, 0x62, 0x52, 0x37, 0xba, 0xa4, 0x89, 0xeb, 0xac, 0x1a, 0x6b, 0x35, 0xbc,
	0xcc, 0xd2, 0x53, 0x8d, 0xca, 0x5b, 0xd9, 0x15, 0x18, 0xfa, 0x0e, 0x1a, 0x79, 0x8c, 0xba, 0xc2,
	0xa5, 0x55, 0xe3, 0x4a, 0xc9, 0x65, 0xfd, 0xd4, 0x15, 0x2e, 0x64, 0xae, 0xd2, 0xf4, 0x36, 0x01,
	0xc6, 0xf7, 0x2b, 0xa4, 0x2c, 0x7f, 0xd4, 0x35, 0x0a, 0x17, 0x20, 0x7d, 0xb0, 0x02, 0xbd, 0x1d,
	0x30, 0x5f, 0x11, 0x4e, 0x85, 0xa2, 0x26, 0xa4, 0x2b, 0xd7, 0x82, 0x11, 0xfd, 0x68, 0xc0, 0x2f,
	0xb4, 0x64, 0x95, 0x21, 0xd4, 0x18, 0x90, 0x4c, 0x60, 0x62, 0xe9, 0xfd, 0x53, 0x81, 0x8a, 0x4c,
	0x7a, 0x43, 0xa5, 0x3e, 0x04, 0x8b, 0xc8, 0x6a, 0xb9, 0xe5, 0xeb, 0x8b, 0xa8, 0x61, 0xf4, 0x15,
	0x94, 0xf9, 0x28, 0x76, 0xcd, 0x55, 0x63, 0xad, 0x51, 0x24, 0xf1, 0x7a, 0x67, 0x14, 0x53, 0x2c,
	0x70, 0xf1, 0x9d, 0x21, 0x39, 0xa3, 0xa1, 0x5b, 0x91, 0x3f, 0xa2, 0x0c, 0xf4, 0x25, 0x98, 0x01,
	0xe1, 0xd4, 0xb5, 0x64, 0xfd, 0x16, 0xf2, 0x68, 0x71, 0x5c, 0x2c, 0xa1, 0x42, 0xcb, 0xa8, 0x4e,
	0xb5, 0x8c, 0x5c, 0x12, 0xb5, 0x49, 0x49, 0x3c, 0x80, 0x8a, 0xba, 0x47, 0xfb, 0x5a, 0x66, 0x2b,
	0xf0, 0x13, 0x85, 0x23, 0x3e, 0x9a, 0xa5, 0x5d, 0x2d, 0x9a, 0x89, 0x8f, 0x66, 0x69, 0x17, 0x4b,
	0x48, 0x34, 0xcc, 0x84, 0x5e, 0x32, 0xfa, 0x56, 0x8a, 0xc5, 0xc6, 0xda, 0x12, 0x3a, 0xea, 0x46,
	0xf1, 0x28, 0x61, 0xbd, 0x0b, 0xae, 0xf5, 0x30, 0xd6, 0xd1, 0x6e, 0x86, 0xe0, 0xb1, 0xd3, 0x27,
	0xab, 0x60, 0xad, 0xa8, 0x82, 0x22, 0xb7, 0x72, 0x11, 0x3c, 0x87, 0x7a, 0x4a, 0x42, 0x9a, 0xa9,
	0xd4, 0x29, 0x14, 0xa1, 0x4d, 0x42, 0xaa, 0x15, 0x0a, 0x69, 0xbe, 0x16, 0x51, 0xb2, 0x86, 0x1f,
	0x17, 0x00, 0x48, 0x3f, 0xc5, 0xfe, 0x6d, 0x30, 0x05, 0x25, 0x90, 0x0d, 0x95, 0xe6, 0xe1, 0xce,
	0xeb, 0x23, 0xc7, 0x40, 0x00, 0x56, 0xbb, 0x75, 0x7c, 0x70, 0xb8, 0xe7, 0x94, 0xd0, 0x22, 0xd4,
	0x77, 0x4f, 0x8e, 0x4e, 0x5b, 0x87, 0xcd, 0x4e, 0xeb, 0xe4, 0xd8, 0x29, 0x23, 0x0b, 0x4a, 0x7b,
	0xa7, 0x8e, 0xe9, 0xfd, 0x6d, 0x42, 0x45, 0x35, 0xfc, 0x9b, 0xb1, 0x37, 0xd7, 0x55, 0x59, 0x7e,
	0xd7, 0xf5, 0xba, 0x9a, 0xe0, 0xb8, 0xf9, 0x61, 0x8e, 0xdf, 0x06, 0x6b, 0x30, 0xec, 0x9f, 0xd1,
	0x44, 0xb2, 0x77, 0x09, 0x6b, 0x0b, 0xdd, 0x83, 0xba, 0xb8, 0x6e, 0x5f, 0x83, 0x96, 0x22, 0xa7,
	0xd8, 0x3a, 0x56, 0x0e, 0x77, 0xa1, 0x16, 0x0c, 0x13, 0xd9, 0xff, 0x35, 0x75, 0x73, 0xbb, 0x40,
	0xec, 0xda, 0x14, 0xb1, 0xef, 0x42, 0x8d, 0xbe, 0x8b, 0x43, 0xd6, 0x65, 0xdc, 0xb5, 0x65, 0x37,
	0xca, 0xed, 0x4f, 0x24, 0x6e, 0x81, 0x4b, 0xf5, 0x9b, 0x72, 0xe9, 0x6b, 0x30, 0xcf, 0x59, 0x48,
	0xa7, 0x1a, 0x7f, 0x73, 0x18, 0xb0, 0x68, 0x9f, 0x85, 0x14, 0x4b, 0x1c, 0x6d, 0x88, 0x61, 0x2c,
	0x7f, 0x8d, 0xb3, 0x4b, 0xaa, 0xf9, 0x5d, 0x7c, 0x9e, 0x4c, 0xba, 0x14, 0xb9, 0xd7, 0xb8, 0x19,
	0xf7, 0xbe, 0x85, 0x6a, 0xac, 0xe5, 0xb5, 0x38, 0xf3, 0x93, 0x32, 0x17, 0xef, 0x4f, 0x03, 0x2a,
	0x92, 0x8e, 0x68, 0x05, 0xaa, 0xe2, 0x3b, 0xfd, 0x9c, 0x3f, 0x96, 0x30, 0x5b, 0x01, 0x7a, 0x08,
	0x66, 0xca, 0xde, 0x2b, 0x0a, 0x35, 0x8a, 0x2c, 0x5e, 0x6f, 0xb3, 0xf7, 0x14, 0x4b, 0x07, 0xd1,
	0x6c, 0xde, 0xb2, 0x80, 0x5f, 0xe8, 0xce, 0xaa, 0x0c, 0x41, 0x8f, 0x0b, 0x2a, 0x25, 0x6d, 0x2a,
	0x7a, 0x28, 0xcb, 0xdb, 0x02, 0x53, 0xc4, 0xa2, 0x3a, 0x54, 0x5f, 0xed, 0xed, 0x37, 0x5f, 0x1f,
	0x76, 0x9c, 0x39, 0x41, 0xfd, 0xf6, 0x51, 0xf3, 0xf0, 0xd0, 0x31, 0xc4, 0xf2, 0xb0, 0x89, 0x0f,
	0x04, 0xf3, 0x01, 0xac, 0x37, 0x6a, 0x5d, 0x16, 0x23, 0x62, 0x2c, 0x1f, 0x41, 0x65, 0x26, 0xac,
	0xa9, 0x11, 0xa1, 0x3b, 0x99, 0x04, 0xbd, 0xdf, 0x0c, 0xb0, 0xf3, 0x39, 0x2c, 0x24, 0xc1, 0xe9,
	0x3b, 0xae, 0x1f, 0x8e, 0x72, 0x7d, 0x65, 0xdc, 0x97, 0x3e, 0x32, 0xee, 0xa7, 0x07, 0x5c, 0xb9,
	0x70, 0x33, 0xb3, 0x07, 0xdc, 0x1b, 0x30, 0x5f, 0xe9, 0x16, 0xa8, 0x65, 0x61, 0x5c, 0xd1, 0xcc,
	0x0c, 0xb9, 0xaa, 0x27, 0x5f, 0xf9, 0x43, 0x2f, 0xda, 0x5f, 0xc0, 0xce, 0x5b, 0x24, 0x7a, 0xa4,
	0xc6, 0x8e, 0x21, 0x6f, 0x6c, 0x65, 0xba, 0x87, 0x4e, 0x8c, 0x9e, 0xac, 0x1a, 0xa5, 0x71, 0x35,
	0xbc, 0x65, 0xdd, 0x88, 0x2a, 0x60, 0x9c, 0x3a, 0x73, 0xe2, 0xcf, 0xae, 0x63, 0x78, 0xff, 0x19,
	0x50, 0x9f, 0x50, 0x01, 0x7a, 0x2c, 0xde, 0x7b, 0xe2, 0xd1, 0xcd, 0x68, 0xea, 0x93, 0x30, 0x8c,
	0xde, 0xd2, 0x40, 0xa7, 0x71, 0x72, 0xa0, 0xa9, 0xf6, 0xd1, 0x13, 0xb8, 0x35, 0x76, 0x3e, 0x8f,
	0x92, 0x33, 0x16, 0x04, 0x74, 0x20, 0x99, 0x62, 0x63, 0x94, 0x43, 0xfb, 0x19, 0x82, 0x1e, 0x4f,
	0x4e, 0xce, 0x3b, 0xd7, 0xc9, 0x70, 0xe2, 0x10, 0xf7, 0x61, 0xa1, 0x4b, 0x38, 0x09, 0xa3, 0xde,
	0x90, 0xfa, 0x29, 0x4f, 0xe4, 0x53, 0xd7, 0xc6, 0xf3, 0xf9, 0x66, 0x9b, 0x27, 0xde, 0x67, 0xfa,
	0x54, 0x0b, 0x60, 0xb7, 0x3b, 0x78, 0xaf, 0x79, 0xd4, 0x3a, 0x3e, 0x70, 0xe6, 0xbc, 0xdf, 0x0d,
	0x80, 0xb1, 0x94, 0x8a, 0x6d, 0xc0, 0xb8, 0x69, 0x1b, 0xb8, 0x0f, 0x15, 0xf9, 0xdf, 0x82, 0xac,
	0xc0, 0xd4, 0xb4, 0x56, 0x18, 0xba, 0x07, 0x65, 0x3a, 0x08, 0x74, 0xdf, 0x2d, 0xb8, 0x08, 0xc4,
	0x5b, 0x07, 0x18, 0xb7, 0x27, 0xd1, 0xce, 0xb3, 0x6b, 0xb4, 0xd5, 0x41, 0x1b, 0x50, 0x62, 0x59,
	0x91, 0x4b, 0x2c, 0xf0, 0xfe, 0x28, 0x81, 0x9d, 0xab, 0x7a, 0xb6, 0x84, 0x9f, 0x82, 0x75, 0x1e,
	0x25, 0x7d, 0xc2, 0xb5, 0x88, 0xef, 0x4c, 0xb7, 0x84, 0xf5, 0x7d, 0xe9, 0x80, 0xb5, 0xa3, 0xf7,
	0xaf, 0x01, 0x96, 0xda, 0x42, 0x4b, 0xb0, 0x70, 0x72, 0x70, 0xe0, 0xff, 0x74, 0x82, 0x77, 0x5a,
	0x6d, 0xff, 0xc5, 0xb6, 0x33, 0x87, 0x10, 0x34, 0x26, 0xb6, 0x9e, 0x6e, 0x6f, 0x38, 0x46, 0x61,
	0xef, 0xd9, 0xe6, 0x86, 0x53, 0x12, 0xe2, 0x3e, 0x3a, 0x7d, 0xe6, 0x6f, 0x6e, 0x6d, 0x3b, 0xe5,
	0xcc, 0x10, 0x88, 0x99, 0x19, 0x22, 0xb4, 0x22, 0x04, 0x2e, 0x8c, 0x17, 0xdb, 0x8e, 0x25, 0xc6,
	0x9c, 0x06, 0xfc, 0xbd, 0xe3, 0x5d, 0xa7, 0x2a, 0xc0, 0x93, 0xce, 0x0f, 0x7b, 0x78, 0xd3, 0xa9,
	0xe5, 0xeb, 0x67, 0x8e, 0x2d, 0x32, 0x34, 0x9b, 0xbb, 0x32, 0x03, 0x64, 0x86, 0xc8, 0x5d, 0xcf,
	0xbd, 0x9e, 0x3b, 0xf3, 0xf9, 0x7a, 0xcb, 0x59, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x54, 0x99,
	0x9b, 0x70, 0x5c, 0x0e, 0x00, 0x00,
}
