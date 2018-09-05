package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templ = template.Must(template.New("index").Parse(templateStr))
var dashboard = template.Must(template.New("index").Parse(templateDashboardStr))
var addr = flag.String("addr", ":1234", "http service address")

type Configuration struct {
	Port                      string `json:"Port"`
	FirebaseAPIKey            string `json:"FIREBASE_API_KEY"`
	FirebaseAuthDomain        string `json:"FIREBASE_AUTH_DOMAIN"`
	FirebaseDatabaseURL       string `json:"FIREBASE_DATABASE_URL"`
	FirebaseProjectID         string `json:"FIREBASE_PROJECT_ID"`
	FirebaseStorageBucket     string `json:"FIREBASE_STORAGEBUCKET"`
	FirebaseMessagingSenderID string `json:"FIREBASE_MESSAGINGSENDERID"`
}

var configuration = Configuration{}

func main() {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	fmt.Print(`Server listening on 0.0.0.0 port ` + configuration.Port + `
Initializing Firebase Configuration
------------------------------------------------------------------------------
FirebaseAPIKey: ` + configuration.FirebaseAPIKey + `
FirebaseAuthDomain: ` + configuration.FirebaseAuthDomain + `
FirebaseDatabaseURL: ` + configuration.FirebaseDatabaseURL + `
FirebaseProjectID: ` + configuration.FirebaseProjectID + `
FirebaseStorageBucket: ` + configuration.FirebaseStorageBucket + `
FirebaseMessagingSenderID: ` + configuration.FirebaseMessagingSenderID + `
------------------------------------------------------------------------------
Press CTRL+C to exit
`)
	if err != nil {
		log.Fatal(err)
		return
	}

	flag.Parse()
	http.HandleFunc("/asset/images/digitize_1024.png", serveImage)
	http.HandleFunc("/asset/css/api.css", serveCSS)
	http.HandleFunc("/", addHandler)
	http.HandleFunc("/dashboard", addHandlerDashboard)
	if err := http.ListenAndServe(configuration.Port, nil); err != nil {

		log.Fatal("ListenAndServe: ", err)
	}
}

func addHandlerDashboard(w http.ResponseWriter, r *http.Request) {
	dashboard.Execute(w, r)
	log.Print("ok")
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "asset/images/digitize_1024.png")
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "asset/css/api.css")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, r)
	log.Print("ok")
}
func setHeader(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
	a := string(b)
	log.Println(a)
	log.Print(b)
	log.Print("Handle is ok")
}

const templateDashboardStr = `<html>
<head>
	<title>Login Coins Base</title>
</head>
<body>
Dashboard
</body>
</html>`

const templateStr = `<html>
<head>
	<title>Authentication</title>
	<script src="https://www.gstatic.com/firebasejs/5.4.1/firebase.js"></script>
	<script src="https://cdn.firebase.com/libs/firebaseui/3.4.0/firebaseui.js"></script>
	<script src="https://www.gstatic.com/firebasejs/ui/3.4.0/firebase-ui-auth__th.js"></script>

	<link type="text/css" rel="stylesheet" href="https://cdn.firebase.com/libs/firebaseui/3.4.0/firebaseui.css" />
	<link type="text/css" rel="stylesheet" href="https://www.gstatic.com/firebasejs/ui/3.4.0/firebase-ui-auth.css" />

	<script>
	  // Initialize Firebase
	  var config = {
		apiKey: "",
		authDomain: "coinsbase-435a8.firebaseapp.com",
		databaseURL: "https://coinsbase-435a8.firebaseio.com",
		projectId: "coinsbase-435a8",
		storageBucket: "coinsbase-435a8.appspot.com",
		messagingSenderId: "469523822327"
	  };
	  firebase.initializeApp(config);
	</script>
    <script type="text/javascript">
      // FirebaseUI config.
      var uiConfig = {
        signInSuccessUrl: 'http://localhost:1234/dashboard',
        signInOptions: [
          // Leave the lines as is for the providers you want to offer your users.
          firebase.auth.GoogleAuthProvider.PROVIDER_ID,
          firebase.auth.FacebookAuthProvider.PROVIDER_ID,
        //   firebase.auth.TwitterAuthProvider.PROVIDER_ID,
          firebase.auth.GithubAuthProvider.PROVIDER_ID,
          firebase.auth.EmailAuthProvider.PROVIDER_ID,
          firebase.auth.PhoneAuthProvider.PROVIDER_ID,
        ],
        // Terms of service url/callback.
        tosUrl: 'http://localhost:1234/',
        // Privacy policy url/callback.
        privacyPolicyUrl: function() {
          window.location.assign('http://localhost:1234/');
        }
      };

      // Initialize the FirebaseUI Widget using Firebase.
      var ui = new firebaseui.auth.AuthUI(firebase.auth());
      // The start method will wait until the DOM is loaded.
      ui.start('#firebaseui-auth-container', uiConfig);
    </script>
</head>
<body>
<h1>
Hello
<div id="firebaseui-auth-container"></div>
</h1>
</body>
</html>`
