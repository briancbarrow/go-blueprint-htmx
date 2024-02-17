package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

type GetTokenRequestBody struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func HandleGetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleGetToken")
	code := r.URL.Query().Get("code")
	repoName := "temp-repo"

	// randomState, err := generateRandomString(16)
	// if err != nil {
	// 	http.Error(w, "Error generating random state", http.StatusInternalServerError)
	// 	return
	// }
	// randomState += ":" + repoName

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			os.Getenv("GA_CLIENT_ID"),
			os.Getenv("GA_CLIENT_SECRET"), code),
		// randomSt),
		bytes.NewBuffer([]byte{}))

	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HandleGetToken", "After post request to get access token")
	// parse response parameters
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	fmt.Println("body", string(body))

	responseValues, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
		return
	}

	accessToken := responseValues.Get("access_token")
	fmt.Println("HandleGetToken", "accessToken: ", accessToken)

	// repoName = responseValues.Get("state")

	gitUrl := createRepo(accessToken, repoName)
	createGitInit(repoName, gitUrl)

}

type RepoResponse struct {
	GitUrl   string `json:"git_url"`
	CloneUrl string `json:"clone_url"`
}

func createRepo(accessToken, repoName string) string {
	fmt.Println("createRepo", "Creating repo...")

	data := []byte(fmt.Sprintf(`{"name": %s,"description":"This is a test repo!","homepage":"https://github.com","private":false,"is_template":true}`, repoName))
	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	var repoResponse RepoResponse
	err = json.NewDecoder(resp.Body).Decode(&repoResponse)
	if err != nil {
		log.Fatal("Error decoding response: ", err)
	}

	defer resp.Body.Close()
	fmt.Println("createRepo", "Repo created successfully", repoResponse.GitUrl)
	return repoResponse.GitUrl
}

func createGitInit(repoName, gitUrl string) {
	fmt.Println("createGitInit", "Creating git init...")
	// git init
	// git remote add origin
	os.Chdir("~/code")
	fmt.Println("createGitInit", "Chdir to ~/code")
	exec.Command("go-blueprint", "-d", "mysql", "-f", "chi", "-n", repoName).Run()
	fmt.Println("createGitInit", "go-blueprint executed")
	repoPath := fmt.Sprintf("~/code/%s", repoName)
	// exec.Command("mkdir", repoPath).Run()
	os.Chdir(repoPath)
	fmt.Println("createGitInit", "Chdir to repoPath")
	exec.Command("git", "init").Run()
	fmt.Println("createGitInit", "Git init executed")
	exec.Command("git", "remote", "add", "origin", gitUrl).Run()
	fmt.Println("createGitInit", "Git remote add origin executed")
	exec.Command("git", "add", ".").Run()
	fmt.Println("createGitInit", "Git add executed")
	exec.Command("git", "commit", "-m", "Initial commit from go-blueprint").Run()
	fmt.Println("createGitInit", "Git commit executed")
	exec.Command("git", "push", "-u", "origin", "main").Run()
	fmt.Println("createGitInit", "Git push executed")
	fmt.Println("createGitInit", "Git init created successfully")
}

// func generateRandomString(n int) (string, error) {
// 	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// 	bytes := make([]byte, n)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	for i, b := range bytes {
// 		bytes[i] = letters[b%byte(len(letters))]
// 	}

// 	return string(bytes), nil
// }
