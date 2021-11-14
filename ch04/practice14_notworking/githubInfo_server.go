package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe("localhost:8010", nil))
}

//Handlerを書く
func requestHandler(w http.ResponseWriter, r *http.Request) {
	var data githubData
	if issues, err := getIssues(); err == nil {
		data.Issues = issues
	}
	if milestones, err := getMilestones(); err == nil {
		data.Milestones = milestones
	}
	if users, err := getMembers(); err == nil {
		data.Users = users
	}
	render(w, data)
}

func render(w io.Writer, data githubData) {
	if err := htmlTemplate.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

//以下、必要そうな構造体の定義
//issue, milestone, userのそれぞれの構造体配列をまとめた構造体
type githubData struct {
	Issues     []*Issue
	Milestones []*Milestone
	Users      []*User //Userはもう一つ構造体ある
}

type User struct {
	Login     string
	Id        int
	NodeId    string `json:"node_id"`
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

//dataの中身
type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"` //JSON表記でsnakecaseのものはタグをつけて変換
	Title   string
	State   string
	User    *User
}

//とりあえずissueと同じ内容。
type Milestone struct {
	Number  int
	HTMLURL string `json:"html_url"` //JSON表記でsnakecaseのものはタグをつけて変換
	Title   string
	State   string
	User    *User
}

//テンプレートを用意、とりあえずissuesをそのまま
var htmlTemplate = template.Must(template.New("htmlTemplate").Parse(`
	<h1>issues</h1>
	<table>
		<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
		</tr>
		{{range .Issues}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>

		<h1>milestones</h1>
		<table>
			<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
			</tr>
			{{range .Milestones}}
			<tr>
				<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
				<td>{{.State}}</td>
				<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
				<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
			</tr>
			{{end}}
			</table>

		<h1>members</h1>
		<table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>User</th>
		</tr>
		{{range .Users}}
		<tr>
			<td><img src='{{.AvatarURL}}' height='64px' width='64px'></td>
			<td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
		</tr>
		{{end}}
`))

//以下でデータを取得
func getIssues() ([]*Issue, error) {
	resp, err := http.Get("https://api.github.com/repos/golang/go/issues")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getIssues: %v\n", err)
		return nil, err
	}
	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func getMilestones() ([]*Milestone, error) {
	resp, err := http.Get("https://api.github.com/repos/golang/go/milestones")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getMilestones: %v\n", err)
		return nil, err
	}
	var result []*Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func getMembers() ([]*User, error) {
	resp, err := http.Get("https://api.github.com/orgs/golang/members")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getMembers: %v\n", err)
		return nil, err
	}
	var result []*User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
