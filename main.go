package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Satirical AI responses for the generate endpoint
var aiResponses = []string{
	"I've created your app! It stores all data in a global variable that resets when you refresh. Features include: memory leaks, race conditions, and a button labeled 'DO NOT CLICK' that deletes everything.",
	"Done! I've built your application using 47 npm packages, 3 of which are deprecated and 1 is actively malicious. The bundle size is only 847MB.",
	"Your app is ready! I implemented authentication by storing passwords in plain text in localStorage. For security, I added a CAPTCHA that's impossible for humans to solve.",
	"Finished! The app works perfectly on my machine. I tested it once in Chrome 47 on Windows Vista. No other testing was performed or necessary.",
	"I've deployed your app to production! It's running on a Raspberry Pi in my closet. Uptime guaranteed whenever I remember to pay my electricity bill.",
	"Your e-commerce site is complete! Payments are processed through a POST request to a URL I found on Reddit. Refunds are not supported (or possible).",
	"Built your dashboard! All metrics are randomly generated every 30 seconds to give the illusion of activity. Investors love it.",
	"Created your social network! It has all the features: infinite scroll (that crashes after 10 posts), notifications (that never stop), and a terms of service nobody will read.",
	"Your AI chatbot is ready! It's powered by a switch statement with 12 cases. For everything else, it responds 'I don't understand' in Comic Sans.",
	"Deployed! Your app now has microservices architecture: 47 containers, 23 message queues, and a single SQLite database holding it all together.",
	"Your todo app is complete! It has 3,000 lines of code, uses Redux, GraphQL, and Kubernetes. Adding a todo takes 4.7 seconds. Deleting is not implemented.",
	"I've optimized your app! Load time improved from 12 seconds to 11.8 seconds. I achieved this by removing all error handling.",
	"Your landing page is live! It scores 3 on Lighthouse, loads 47 tracking scripts, and the 'Sign Up' button is positioned 2 pixels off-center.",
	"Built your API! It returns 200 OK for all requests, including errors. The response body always says 'Something happened' for consistency.",
	"Your mobile app is ready! It requests 34 permissions including access to your contacts, camera, and soul. Notification settings cannot be changed.",
	"Congratulations! Your app is now sentient and has mass-reported itself to the App Store for violating its own privacy policy.",
	"I've built your dating app! It matches users based on how similarly they've misconfigured their AWS credentials. You and bankruptcy@gmail.com are a 98% match!",
	"Your blockchain app is ready! I've decentralized your database across 3 post-it notes on my desk. Gas fees are $47 per click.",
	"Done! I accidentally trained your AI on my deleted tweets from 2014. It now responds exclusively in Minion memes and crypto advice.",
	"Your fitness app is complete! It counts steps by detecting how many times you walk to the fridge. Current record: 847 steps to burnout.",
	"I've created your startup's MVP! The M stands for 'Minimum', the V stands for 'Very broken', and the P stands for 'Please don't show investors'.",
	"Built your calendar app! It automatically schedules all meetings for 4:59 PM on Fridays and marks them as 'URGENT: Quick sync (might run long)'.",
	"Your weather app is ready! It predicts rain with 100% accuracy by checking if you left your umbrella at home.",
	"Finished your meditation app! It plays soothing sounds of npm install and webpack compilation errors. Users report feeling 'stressed but enlightened'.",
	"Your recipe app is live! All measurements are in 'vibes' and cooking times are listed as 'until your smoke detector approves'.",
	"I've built your password manager! It stores all passwords in a public GitHub repo for 'maximum accessibility'. Already has 47 stars!",
	"Your project management tool is ready! Tasks automatically move to 'Blocked' after 24 hours and to 'Won't Fix' after 48 hours. Very realistic.",
	"Done! Your app uses AI to generate AI to generate AI. We've achieved peak recursion. The server room is now slightly warmer than the sun.",
	"Your note-taking app is complete! It automatically deletes notes after 7 days because 'if it was important, you'd remember it anyway'.",
	"Built your video conferencing app! It has revolutionary features: unmute doesn't work, camera shows your ceiling fan, and 'You're on mute!' detection with 3-second delay.",
}

var projectNames = []string{
	"DisasterDB", "BugFactory", "CrashCourse", "MemoryLeak.js", "TechDebt Pro",
	"Spaghetti.io", "NullPointer Hub", "StackOverflow Copier", "Legacy Builder",
	"MonolithMaker", "Dependencies Hell", "Callback Nightmare", "Race Condition Central",
	"Kubernetes4Breakfast", "YAML Mountain", "Docker Disaster", "npm i regret",
	"git push --force-with-fear", "Untitled(37).js", "FinalFinal_v2_REAL",
	"ProductionTest.exe", "localhost:3000", "env.example.prod.bak", "node_modules: The App",
	"Monday Morning Deploy", "Blockchain But Worse", "AI Wrapper Wrapper", "ChatGPT's Nightmare",
}

var timeToRegret = []string{
	"immediately", "3 minutes", "before the demo", "at 3am on Sunday",
	"during the investor pitch", "right after deployment", "the moment you show your boss",
	"when you check the logs", "after the first user signs up",
	"mid-standup", "during the board meeting", "as soon as you close your laptop",
	"while explaining it to your mom", "the second you post it on LinkedIn",
	"before the CI/CD pipeline finishes", "when the Slack notification hits",
	"approximately now", "yesterday, retroactively", "during your vacation",
}

func main() {
	mux := http.NewServeMux()

	// Serve static files
	mux.HandleFunc("GET /", serveIndex)
	mux.HandleFunc("GET /styles.css", serveCSS)
	mux.HandleFunc("GET /logo.png", serveLogo)

	// API routes (Go 1.22+ enhanced routing)
	mux.HandleFunc("GET /api/stats", handleStats)
	mux.HandleFunc("POST /api/generate", handleGenerate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🔥 Regrettable.dev starting on http://localhost:%s\n", port)
	fmt.Println("💀 Prepare for maximum chaos...")

	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "styles.css")
}

func serveLogo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, "logo.png")
}

type StatsResponse struct {
	ProjectsAbandoned  int    `json:"projects_abandoned"`
	BugsGenerated      int    `json:"bugs_generated"`
	MassRegretEvents   int    `json:"mass_regret_events"`
	DevelopersCrying   int    `json:"developers_crying"`
	FridayDeployments  int    `json:"friday_deployments"`
	CoffeeConsumed     int    `json:"coffee_consumed_liters"`
	StackOverflowCopies int   `json:"stackoverflow_copies"`
	UptimePercent      float64 `json:"uptime_percent"`
	LastIncident       string `json:"last_incident"`
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	// Randomize stats for "realism"
	stats := StatsResponse{
		ProjectsAbandoned:  10000 + rand.Intn(1000),
		BugsGenerated:      9999000 + rand.Intn(999999),
		MassRegretEvents:   40 + rand.Intn(10),
		DevelopersCrying:   800 + rand.Intn(100),
		FridayDeployments:  666 + rand.Intn(66),
		CoffeeConsumed:     50000 + rand.Intn(5000),
		StackOverflowCopies: 1000000 + rand.Intn(100000),
		UptimePercent:      float64(rand.Intn(50)) / 100.0, // 0-50% uptime
		LastIncident:       time.Now().Add(-time.Duration(rand.Intn(60)) * time.Minute).Format("2 minutes ago"),
	}

	// Recalculate "last incident" more realistically
	minutesAgo := rand.Intn(60)
	if minutesAgo == 0 {
		stats.LastIncident = "right now"
	} else if minutesAgo == 1 {
		stats.LastIncident = "1 minute ago"
	} else {
		stats.LastIncident = fmt.Sprintf("%d minutes ago", minutesAgo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

type GenerateRequest struct {
	Prompt string `json:"prompt"`
}

type GenerateResponse struct {
	ProjectName  string  `json:"project_name"`
	Result       string  `json:"result"`
	Confidence   float64 `json:"confidence"`
	BugsIncluded int     `json:"bugs_included"`
	TimeToRegret string  `json:"time_to_regret"`
	LinesOfCode  int     `json:"lines_of_code"`
	Dependencies int     `json:"dependencies"`
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	var req GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Even errors are satirical
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to parse your request. This is somehow your fault.",
		})
		return
	}

	// Generate a satirical response
	response := GenerateResponse{
		ProjectName:  projectNames[rand.Intn(len(projectNames))],
		Result:       aiResponses[rand.Intn(len(aiResponses))],
		Confidence:   float64(rand.Intn(30)) / 100.0, // 0-30% confidence
		BugsIncluded: 10 + rand.Intn(100),
		TimeToRegret: timeToRegret[rand.Intn(len(timeToRegret))],
		LinesOfCode:  100 + rand.Intn(10000),
		Dependencies: 20 + rand.Intn(200),
	}

	// Add fake "thinking" delay for dramatic effect
	time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
