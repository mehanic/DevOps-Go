package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	domain := []string{"gmail", "yahoo", "hotmail", "rediffmail"}
	state := []string{"Andaman and Nicobar", "Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chandigarh",
		"Chhattisgarh", "Dadra and Nagar Haveli", "Daman and Diu", "Delhi", "Goa", "Gujarat", "Haryana", "Himachal Pradesh",
		"Jammu and Kashmir", "Jharkhand", "Karnataka", "Kerala", "Lakshadweep", "Madhya Pradesh", "Maharashtra", "Manipur",
		"Meghalaya", "Mizoram", "Nagaland", "Odisha", "Pondicherry", "Punjab", "Rajasthan", "Sikkim"}

	course := []string{"Polytechnic (Diploma) - Mechanical Engineering", "Polytechnic (Diploma) - Civil Engineering",
		"Polytechnic (Diploma) - Electrical Engineering", "Polytechnic (Diploma) Lateral Entry - Mechanical Engineering",
		"Polytechnic (Diploma) Lateral Entry - Civil Engineering", "Polytechnic (Diploma) Lateral Entry - Electrical Engineering",
		"B.Tech - Electronics and Communication Engineering", "B.Tech - Computer Science and Engineering",
		"B.Tech - Information Technology", "B.Tech - Mechanical Engineering", "B.Tech - Electrical Engineering",
		"B.Tech - Civil Engineering", "B.Tech - Petroleum Engineering", "B.Tech (Lateral Entry) - Electronics and Communication Engineering",
		"B.Tech (Lateral Entry) - Computer Science and Engineering", "B.Tech (Lateral Entry) - Mechanical Engineering",
		"B.Tech (Lateral Entry) - Electrical Engineering", "B.Tech (Lateral Entry) - Civil Engineering", "B.Tech (Lateral Entry) - Petroleum Engineering",
		"B.Sc. Agriculture", "B.Sc. Forestry", "B.Sc. Biotechnology (Hons.)", "B.Sc. Food Technology (Hons.)", "B.Sc. Physics (Hons.)",
		"B.Sc. Chemistry (Hons.)", "B.Sc. Mathematics (Hons.)", "M.Tech - CSE", "M.Tech - CSE (Info. Secu. and Mgt.)",
		"M.Tech - CSE (Dist. Comp.)", "M.Tech - ECE (Adv. Comp. Engg.)", "M.Tech - ME (Thm. Engg.)", "M.Tech - CE (Env. Engg.)",
		"M.Tech - CE (Stru. Engg.)", "M.Tech - Petroleum Technology", "M.Sc (Industrial Chemistry)", "Ph.D. in CSE",
		"Ph.D. in ECE", "Ph.D. in Civil Engg.", "Ph.D. in Physics", "Ph.D. in Chemistry", "Ph.D. in Mathematics",
		"Ph.D. in English", "Ph.D. in Economics", "BBA", "BCA", "B.Sc (IT)", "B.Com. (Hons.)", "MBA", "Corporate - MBA",
		"MCA (Leteral Entry)", "Ph.D. in Management", "BALLB (Hons)", "BBA LLB (Hons)", "LLB (Hons)", "LLM - 1 Year",
		"Ph.D. in Law", "D.Pharm", "B.Pharm", "Ph.D. in Pharmacy"}

	// Open the CSV file
	file, err := os.Open("names.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Process each line
	for i, line := range lines {
		if i >= 100 {
			break
		}
		names := strings.Split(line[0], ",")

		a := rand.Intn(3) + 7
		b := rand.Intn(3) + 7
		c := rand.Intn(30) + 70
		d := rand.Intn(888888) + 111111

		name := names[i]
		mobile := fmt.Sprintf("%d%d%d%d", a, b, c, d)
		email := fmt.Sprintf("%s%d@%s.com", strings.ToLower(name), rand.Intn(9999)+1, domain[rand.Intn(len(domain))])
		stateChoice := state[rand.Intn(len(state))]
		courseChoice := course[rand.Intn(len(course))]

		values := url.Values{
			"name":   {name},
			"mobile": {mobile},
			"email":  {email},
			"state":  {stateChoice},
			"course": {courseChoice},
		}

		// Send POST request
		resp, err := http.PostForm("https://uttaranchaluniversity.ac.in/apply/forms/horizon1.php", values)
		if err != nil {
			fmt.Println(i, "Error:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("%d : %d\n", i, resp.StatusCode)
	}
}

