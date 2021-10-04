package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const plex_server = "192.168.9.10:32400"
const plex_token = ""

// Plex_token string := os.Getenv(plex_token)

type Library struct {
	XMLName         xml.Name `xml:"MediaContainer"`
	Text            string   `xml:",chardata"`
	Size            string   `xml:"size,attr"`
	AllowSync       string   `xml:"allowSync,attr"`
	Identifier      string   `xml:"identifier,attr"`
	MediaTagPrefix  string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string   `xml:"mediaTagVersion,attr"`
	Title1          string   `xml:"title1,attr"`
	Directory       []struct {
		Text             string `xml:",chardata"`
		AllowSync        string `xml:"allowSync,attr"`
		Art              string `xml:"art,attr"`
		Composite        string `xml:"composite,attr"`
		Filters          string `xml:"filters,attr"`
		Refreshing       string `xml:"refreshing,attr"`
		Thumb            string `xml:"thumb,attr"`
		Key              string `xml:"key,attr"`
		Type             string `xml:"type,attr"`
		Title            string `xml:"title,attr"`
		Agent            string `xml:"agent,attr"`
		Scanner          string `xml:"scanner,attr"`
		Language         string `xml:"language,attr"`
		Uuid             string `xml:"uuid,attr"`
		UpdatedAt        string `xml:"updatedAt,attr"`
		CreatedAt        string `xml:"createdAt,attr"`
		ScannedAt        string `xml:"scannedAt,attr"`
		Content          string `xml:"content,attr"`
		Directory        string `xml:"directory,attr"`
		ContentChangedAt string `xml:"contentChangedAt,attr"`
		Hidden           string `xml:"hidden,attr"`
		Location         struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Path string `xml:"path,attr"`
		} `xml:"Location"`
	} `xml:"Directory"`
}

type Media struct {
	XMLName             xml.Name `xml:"MediaContainer"`
	Text                string   `xml:",chardata"`
	Size                string   `xml:"size,attr"`
	AllowSync           string   `xml:"allowSync,attr"`
	Art                 string   `xml:"art,attr"`
	Identifier          string   `xml:"identifier,attr"`
	LibrarySectionID    string   `xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `xml:"mediaTagVersion,attr"`
	Thumb               string   `xml:"thumb,attr"`
	Title1              string   `xml:"title1,attr"`
	Title2              string   `xml:"title2,attr"`
	ViewGroup           string   `xml:"viewGroup,attr"`
	ViewMode            string   `xml:"viewMode,attr"`
	Video               []struct {
		Text                  string `xml:",chardata"`
		RatingKey             string `xml:"ratingKey,attr"`
		Key                   string `xml:"key,attr"`
		Guid                  string `xml:"guid,attr"`
		Studio                string `xml:"studio,attr"`
		Type                  string `xml:"type,attr"`
		Title                 string `xml:"title,attr"`
		ContentRating         string `xml:"contentRating,attr"`
		Summary               string `xml:"summary,attr"`
		AudienceRating        string `xml:"audienceRating,attr"`
		Year                  string `xml:"year,attr"`
		Tagline               string `xml:"tagline,attr"`
		Thumb                 string `xml:"thumb,attr"`
		Art                   string `xml:"art,attr"`
		Duration              string `xml:"duration,attr"`
		OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
		AddedAt               string `xml:"addedAt,attr"`
		UpdatedAt             string `xml:"updatedAt,attr"`
		AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
		PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
		ViewCount             string `xml:"viewCount,attr"`
		LastViewedAt          string `xml:"lastViewedAt,attr"`
		ChapterSource         string `xml:"chapterSource,attr"`
		TitleSort             string `xml:"titleSort,attr"`
		SkipCount             string `xml:"skipCount,attr"`
		ViewOffset            string `xml:"viewOffset,attr"`
		Media                 []struct {
			Text                  string `xml:",chardata"`
			ID                    string `xml:"id,attr"`
			Duration              string `xml:"duration,attr"`
			Bitrate               string `xml:"bitrate,attr"`
			Width                 string `xml:"width,attr"`
			Height                string `xml:"height,attr"`
			AspectRatio           string `xml:"aspectRatio,attr"`
			AudioChannels         string `xml:"audioChannels,attr"`
			AudioCodec            string `xml:"audioCodec,attr"`
			VideoCodec            string `xml:"videoCodec,attr"`
			VideoResolution       string `xml:"videoResolution,attr"`
			Container             string `xml:"container,attr"`
			VideoFrameRate        string `xml:"videoFrameRate,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			Part                  struct {
				Text                  string `xml:",chardata"`
				ID                    string `xml:"id,attr"`
				Key                   string `xml:"key,attr"`
				Duration              string `xml:"duration,attr"`
				File                  string `xml:"file,attr"`
				Size                  string `xml:"size,attr"`
				Container             string `xml:"container,attr"`
				VideoProfile          string `xml:"videoProfile,attr"`
				AudioProfile          string `xml:"audioProfile,attr"`
				Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
				OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
				HasThumbnail          string `xml:"hasThumbnail,attr"`
			} `xml:"Part"`
		} `xml:"Media"`
		Genre []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Genre"`
		Director []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Director"`
		Writer []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Writer"`
		Country []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Country"`
		Role []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Role"`
	} `xml:"Video"`
}

var video Media

func process_video() {

	// for _, video := range media.Video {
	// 	fmt.Println(" ", video.Title, "Key:", video.Key)

	// 	var url_video = "http://" + plex_server + video.Key + "/?X-Plex-Token=" + plex_token

	// 	response, err := http.Get(url_video)

	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		os.Exit(1)
	// 	}

	// 	responseData, err := ioutil.ReadAll(response.Body)

	// 	var video Media
	// 	xml.Unmarshal(responseData, &video)

	fmt.Println(video.Text)

	// }

}

func main() {

	var url_libraries = "http://" + plex_server + "/library/sections/all?X-Plex-Token=" + plex_token

	// fmt.Println(url_libraries)

	// fmt.Println("Connecting to Plex")
	response, err := http.Get(url_libraries)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var library Library
	xml.Unmarshal(responseData, &library)

	// for each library
	for _, library := range library.Directory {

		fmt.Println("Library:", library.Title, ", Type:", library.Type, ", Key:", library.Key)

		if library.Type == "movie" {
			// fmt.Println("This is a movie")
			var url_movie = "http://" + plex_server + "/library/sections/" + library.Key + "/all?duplicate=1&X-Plex-Token=" + plex_token
			// fmt.Println(url_movie)
			response, err := http.Get(url_movie)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Fatal(err)
			}

			var media Media
			xml.Unmarshal(responseData, &media)

			for _, video := range media.Video {
				fmt.Println("  Video Title:", video.Title, "Key: ", video.Key)
				// fmt.Println("    ", video.Media)

				score := make(map[string]int)
				// width := 0
				for _, media := range video.Media {

					fmt.Println("     ID:", media.ID, "Size:", media.Part.Size, ", Width:", media.Width, ", Codec:", media.VideoCodec)

					// need sorting logic here (or above?)

					if media.VideoCodec == "hevc" {
						score[media.ID] = score[media.ID] + 1000
					} else {
						score[media.ID] = score[media.ID] + 0
					}

					// if media.Width > width {
					// 	score[media.ID] = score[media.ID] + 100
					// 	width = media.Width
					// }

					//					fmt.Println(score)

				}
				fmt.Println(score)
				// process_video(video)
			}

		}

		if library.Type == "show" {
			//fmt.Println("This is a show")
			var url_show = "http://" + plex_server + "/library/sections/" + library.Key + "/search?type=4&duplicate=1&X-Plex-Token=" + plex_token
			// fmt.Println(url_show)

			response, err := http.Get(url_show)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Fatal(err)
			}

			var media Media
			xml.Unmarshal(responseData, &media)

			for _, video := range media.Video {
				fmt.Println("  Video Title:", video.Title, "Key: ", video.Key)
				// fmt.Println("    ", video.Media)

				for _, media := range video.Media {

					fmt.Println("     ID:", media.ID, "Size:", media.Part.Size, ", Width:", media.Width, ", Codec:", media.VideoCodec)

					// need sorting logic here (or above?)

				}

				// process_video(video)
			}

		}

	}

}
