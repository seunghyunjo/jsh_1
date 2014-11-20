package hello
import (

"net/http"
"text/template"
"appengine"
"appengine/datastore"



)



type User struct {
	Account	string	
	Password string	
	Name string		
}


const loginHTML = `
<html>
<body>
<h2>Please Log In</h2>
	<p>
	Please enter your id and password to log in to this site.
	</p>
	<form method="post" action="/">
	Account: <input type="text" name="account"/> <br>
	Password: <input type="password" name="password"/> <br/>
	<input type="submit" value="Submit" />
	<input type="submit" value="Cancel"
		onclick="window.location='/'; return false;"/>
	<input type="submit" value="New Account"
		onclick="window.location='/apply'; return false;"/>
	</form>
	</p>
</body>
</html>
`
const memberHTML = `
<html>
<body>
<h1>Members</h1>
	{% for %}
	<pre>{{html .}}</pre>
	{% end for %}
</body>
</html>
`
const chatHTML = `
<html>
<body>
<h2>App Engine: About</h2>
   <h1>App Engine Chat</h1>
	<p>
	<form method="post" action="/chat">
	<input type="text" name="message" size="60"/>
	<input type="submit" name="Chat"/>
	</form>
	</p>
	<pre>{{html .}}</pre>
</body>
</html>
`
const applyHTML = `
<html>
<body>
<h1>New Account Request</h1>
	<p>
	Please enter your information below:
	</p>
	<form method="post" action="/apply">
	Name: <input type="text" name="name"/> <br/>
	Account: <input type="text" name="account"/> <br>
	Password: <input type="password" name="password"/> <br/>
	<input type="submit" value="Submit"/>
	<input type="submit" value="Cancel"
		onclick="window.location='/'; return false;"/>
	</form>
	</p>
</body>
</html>
`
var loginTemplate = template.Must(template.New("login").Parse(loginHTML))
var chatTemplate = template.Must(template.New("sign").Parse(chatHTML))
var applyTemplate = template.Must(template.New("apply").Parse(applyHTML))
var memberTemplate = template.Must(template.New("member").Parse(memberHTML))
func init() {
http.HandleFunc("/", login)
http.HandleFunc("/chat", chat)
http.HandleFunc("/apply", apply)

}



func login(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "text/html")
switch r.Method{
case "GET":
loginTemplate.Execute(w,r.FormValue("name"))
case "POST":

a :=r.FormValue("account")
p :=r.FormValue("password")



que:=datastore.NewQuery("User").Filter("Account =",a).Filter("Password =",p).Limit(1)
if que!=nil {
	chatTemplate.Execute(w,r.FormValue("message"))
}else {
	loginTemplate.Execute(w,r.FormValue("account"))
}
}
}
func chat(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "text/html")
err := chatTemplate.Execute(w, r.FormValue("message"))
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}

func apply(w http.ResponseWriter, r *http.Request) {

//w.Header().Set("Content-Type", "text/html")

switch r.Method{
case "GET":
applyTemplate.Execute(w, r.FormValue("name"))
case "POST":
c := appengine.NewContext(r)

a :=r.FormValue("account")

result := datastore.NewQuery("User").Filter("Account =", a).Limit(1)

if result!=nil {
	applyTemplate.Execute(w,r.FormValue("name"))
	
}

user := User {
	Account:r.FormValue("account"),
	Password:r.FormValue("password"),
	Name:r.FormValue("name"),
}

datastore.Put(c, datastore.NewIncompleteKey(c, "User",nil), &user)
//http.Redirect(w, r, "/", http.StatusFound)

chatTemplate.Execute(w, r.FormValue("message"))	

}
}


