package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var templ = template.Must(template.New("index").Parse(templateStr))
var dashboard = template.Must(template.New("index").Parse(templateDashboardStr))
var addr = flag.String("addr", ":1234", "http service address")
var dbconnection = ""

//Configuration struct
type Configuration struct {
	Port             string `json:"Port"`
	DatabaseAddress  string `json:"DATABASE_ADDRESS"`
	DatabaseName     string `json:"DATABASE_NAME"`
	DatabaseUser     string `json:"DATABASE_USER"`
	DatabasePassword string `json:"DATABASE_PASSWORD"`
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
PORT: ` + configuration.Port + `
DatabaseAddress: ` + configuration.DatabaseAddress + `
DatabaseName: ` + configuration.DatabaseName + `
DatabaseUser: ` + configuration.DatabaseUser + `
DatabasePassword: ` + configuration.DatabasePassword + `
------------------------------------------------------------------------------
Press CTRL+C to exit
`)
	dbconnection = configuration.DatabaseUser + ":" + configuration.DatabasePassword + "@tcp(" + configuration.DatabaseAddress + ")/" + configuration.DatabaseName
	testConnection(configuration)
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
func testConnection(configuration Configuration) {
	db, err := sql.Open("mysql", dbconnection)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("SELECT * FROM `members`")

	if err != nil {
		creatMemberTable()
		log.Println("Create Table Members Completed!")
	}

}

func creatMemberTable() {
	db, err := sql.Open("mysql", dbconnection)
	_, err = db.Exec("CREATE TABLE `members` (`mid` int(11) NOT NULL,`telegram_uid` varchar(20) NOT NULL,`first_name` varchar(255) NOT NULL,`lastname` varchar(255) NOT NULL,`is_bot` tinyint(1) NOT NULL,`chat_type` varchar(100) NOT NULL,`created` int(10) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8;")

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("ALTER TABLE `members` ADD UNIQUE KEY `mid` (`mid`);")

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("ALTER TABLE `members` MODIFY `mid` int(11) NOT NULL AUTO_INCREMENT;")

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec("COMMIT;")

	if err != nil {
		log.Fatal(err)
		return
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

const templateStr = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Telegram: Contact @aofieebot</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    
<meta property="og:title" content="AofieeBot">
<meta property="og:image" content="https://cdn5.telesco.pe/file/UAqx08l5ldNx7VrZ3pcthbUCq7xD_m1iCR8vzVI90lrIOk8mm1ODqBM-B1_A7omr-dGfi8xeJlhsNq6eNzJ0Sx4Ze_olyF86V0t4pSP1OnVQ3pV-408YEJePtEXzs8EnalFfVFTyPZJWacCH5lKamtBOCVidM-Jv89up7E0Pif5PNaJEvyRG8IMKA6gCqERpzuQThMGPQ_vZeyNo-4XRP_XZkvNnPTQMDf6kJR-XBbT-5dQbHl-8hzSzuPQAhYQz7wydRAcKpSe7YmRRrVkF_76n921vTNaqGQpyIw43QHDyCSiEJY1wXM8omliN2KhLo-i8bofz2Rw8r4Jfmz0p_Q.jpg">
<meta property="og:site_name" content="Telegram">
<meta property="og:description" content="My name is Aofiee. I like your son.">

<meta property="twitter:title" content="AofieeBot">
<meta property="twitter:image" content="https://cdn5.telesco.pe/file/UAqx08l5ldNx7VrZ3pcthbUCq7xD_m1iCR8vzVI90lrIOk8mm1ODqBM-B1_A7omr-dGfi8xeJlhsNq6eNzJ0Sx4Ze_olyF86V0t4pSP1OnVQ3pV-408YEJePtEXzs8EnalFfVFTyPZJWacCH5lKamtBOCVidM-Jv89up7E0Pif5PNaJEvyRG8IMKA6gCqERpzuQThMGPQ_vZeyNo-4XRP_XZkvNnPTQMDf6kJR-XBbT-5dQbHl-8hzSzuPQAhYQz7wydRAcKpSe7YmRRrVkF_76n921vTNaqGQpyIw43QHDyCSiEJY1wXM8omliN2KhLo-i8bofz2Rw8r4Jfmz0p_Q.jpg">
<meta property="twitter:site" content="@Telegram">

<meta property="al:ios:app_store_id" content="686449807">
<meta property="al:ios:app_name" content="Telegram Messenger">
<meta property="al:ios:url" content="tg://resolve?domain=aofieebot">

<meta property="al:android:url" content="tg://resolve?domain=aofieebot">
<meta property="al:android:app_name" content="Telegram">
<meta property="al:android:package" content="org.telegram.messenger">

<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@Telegram">
<meta name="twitter:description" content="My name is Aofiee. I like your son.
">
<meta name="twitter:app:name:iphone" content="Telegram Messenger">
<meta name="twitter:app:id:iphone" content="686449807">
<meta name="twitter:app:url:iphone" content="tg://resolve?domain=aofieebot">
<meta name="twitter:app:name:ipad" content="Telegram Messenger">
<meta name="twitter:app:id:ipad" content="686449807">
<meta name="twitter:app:url:ipad" content="tg://resolve?domain=aofieebot">
<meta name="twitter:app:name:googleplay" content="Telegram">
<meta name="twitter:app:id:googleplay" content="org.telegram.messenger">
<meta name="twitter:app:url:googleplay" content="https://t.me/aofieebot">

<meta name="apple-itunes-app" content="app-id=686449807, app-argument: tg://resolve?domain=aofieebot">
    <link rel="shortcut icon" href="//telegram.org/favicon.ico?3" type="image/x-icon" />
    <link href="https://fonts.googleapis.com/css?family=Roboto:400,700" rel="stylesheet" type="text/css">
    <!--link href="/css/myriad.css" rel="stylesheet"-->
    <link href="//telegram.org/css/bootstrap.min.css?2" rel="stylesheet">
    <link href="//telegram.org/css/telegram.css?149" rel="stylesheet" media="screen">
  </head>
  <body>

    <div class="tgme_page_wrap">
      <div class="tgme_head_wrap">
        <div class="tgme_head">
          <a href="//telegram.org/" class="tgme_head_brand">
            <i class="tgme_logo"></i>
          </a>
        </div>
      </div>
      <a class="tgme_head_dl_button" href="//telegram.org/dl?tme=96e0898a29536859b5_18360203280349133312">
        Don't have <strong>Telegram</strong> yet? Try it now!<i class="tgme_icon_arrow"></i>
      </a>
      <div class="tgme_page">
        <div class="tgme_page_photo">
  <a href="tg://resolve?domain=aofieebot"><img class="tgme_page_photo_image" src="https://cdn5.telesco.pe/file/UAqx08l5ldNx7VrZ3pcthbUCq7xD_m1iCR8vzVI90lrIOk8mm1ODqBM-B1_A7omr-dGfi8xeJlhsNq6eNzJ0Sx4Ze_olyF86V0t4pSP1OnVQ3pV-408YEJePtEXzs8EnalFfVFTyPZJWacCH5lKamtBOCVidM-Jv89up7E0Pif5PNaJEvyRG8IMKA6gCqERpzuQThMGPQ_vZeyNo-4XRP_XZkvNnPTQMDf6kJR-XBbT-5dQbHl-8hzSzuPQAhYQz7wydRAcKpSe7YmRRrVkF_76n921vTNaqGQpyIw43QHDyCSiEJY1wXM8omliN2KhLo-i8bofz2Rw8r4Jfmz0p_Q.jpg"></a>
</div>
<div class="tgme_page_title">AofieeBot</div>
<div class="tgme_page_extra">
  @aofieebot
</div>
<div class="tgme_page_description">My name is Aofiee. I like your son.</div>
<div class="tgme_page_action">
  <a class="tgme_action_button_new" href="tg://resolve?domain=aofieebot">Send Message</a>
</div>
<!-- WEBOGRAM_BTN -->
<div class="tgme_page_additional">
  If you have <strong>Telegram</strong>, you can contact <br><strong>AofieeBot</strong> right away.
</div>
      </div>
    </div>

    <div id="tgme_frame_cont"></div>

    <script type="text/javascript">

var protoUrl = "tg:\/\/resolve?domain=aofieebot";
if (false) {
  var iframeContEl = document.getElementById('tgme_frame_cont') || document.body;
  var iframeEl = document.createElement('iframe');
  iframeContEl.appendChild(iframeEl);
  var pageHidden = false;
  window.addEventListener('pagehide', function () {
    pageHidden = true;
  }, false);
  window.addEventListener('blur', function () {
    pageHidden = true;
  }, false);
  if (iframeEl !== null) {
    iframeEl.src = protoUrl;
  }
  !false && setTimeout(function() {
    if (!pageHidden) {
      window.location = protoUrl;
    }
  }, 2000);
}
else if (protoUrl) {
  setTimeout(function() {
    window.location = protoUrl;
  }, 100);
}


    </script>
    <script>(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-45099287-3', 'auto', {'sampleRate': 5});
ga('set', 'anonymizeIp', true);
ga('send', 'pageview');</script>
  </body>
</html>
<!-- page generated in 9.22ms -->
`
