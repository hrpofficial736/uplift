package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
)



type User struct {
	Login string `json:"login"`
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
}



type Repo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
}


type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}


func FetchRepoInfo (url string) []Content {
	request, err := http.NewRequest("GET", url, nil);
	if (err != nil) {
		log.Fatal(err);
	}

	token := config.ConfigLoad().GithubAccessToken;
	request.Header.Set("Authorization", "token " + token);
	request.Header.Set("Accept", "application/vnd.github+json");


	client := &http.Client{};
	response, err := client.Do(request);

	if err != nil {
		log.Fatal(response);
		log.Fatalf("Error while making request: %s\n" , err);
	}

	fmt.Println(response);
	defer response.Body.Close();

	// var user User;

	// if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
	// 	log.Fatalf("Error while converting to json: %s\n",err);
	// }

var contents []Content;
	if err := json.NewDecoder(response.Body).Decode(&contents); err != nil {
		log.Fatal(err)
	}
	return contents;
}