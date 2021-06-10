package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("personsmall.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	
// //2.1
// // fmt.Println("---GroupPeopleByCity ----")
// 	peopleByCity := GroupPeopleByCity(people)
// 	for key, value := range peopleByCity {
// 		fmt.Println(key)
// 		for _, person := range value {
// 			fmt.Println("  ", (&person).Name)
// 		}
// 	}
// // 2.2
// fmt.Println("---GroupPeopleByJob ----")
// 	peopleByJob := GroupPeopleByJob(people)
// 	for key, value := range peopleByJob {
// 		fmt.Println(key+"-",value)
// 	}

// // 2.3 
// fmt.Println("---Top5JobsByNumer ----")
// 	top5Jobs := Top5JobsByNumber(peopleByJob)
// 	for _, job := range top5Jobs {
// 		fmt.Printf("%s : %d", job.Key, job.Value)
// 		fmt.Println("")
// 	}
// // 2.4
// fmt.Println("-------Top5CitiesByNumber ------")
// top5Cities := Top5CitiesByNumber(people)
// 	for _, city := range top5Cities {
// 		fmt.Printf("%s :%v", city.Key, city.Value)
// 		fmt.Println("")
// 	}
// fmt.Println(top5Cities)
// // 2.5
// fmt.Println("-------TopJobByNumerInEachCity ------")
// resultMap := make(map[string]map[string]int)
// jobByCity := GroupJobByCity(people)
// for key, value := range jobByCity {
// 	result := CountNumberEachJob(value)
// 	resultMap[key] = result
// }
// fmt.Println(jobByCity)
// topJobEachCity := TopJobByNumberInEachCity(resultMap)
// for key, value := range topJobEachCity {
// 	fmt.Println("-", key)
// 	for k, v := range value {
// 		fmt.Printf("%s : %d", k, v)
// 		fmt.Println("")
// 	}
// }
// // 2.6 
// fmt.Println("-----------------AverageSalaryByJob-------------------")
// numberEachJob := GroupPeopleByJob(people)
// salaryEachJob := SalaryEachJob(people)
// averageSalaryByJob := AverageSalaryByJob(numberEachJob, salaryEachJob)
// for key, value := range averageSalaryByJob {
// 	fmt.Printf("%s : %d", key, value)
// 	fmt.Println("")
// }
// // 2.7 
// fmt.Println("-----------------------Five Cities Has Top Average Salary-----------------")
// numberEachCity := CountPersonByCity(people)
// salaryEachCity := SalaryEachCity(people)
// averageSalaryByCity := FiveCitiesHasTopAverageSalary(numberEachCity, salaryEachCity)
// for _, a := range averageSalaryByCity {
// 	fmt.Printf("%s : %d", a.Key, a.Value)
// 	fmt.Println("")
// }
// 2.8
fmt.Println("-------------Five Cities Has Top Salary For Developer-------------- ")
numberDeveloperEachCity := CountDeveloperByCity(people)
salaryDeveloperEachCity := SalaryDeveloperByCity(people)
averageSalaryDeveloperByCity := FiveCitiesHasTopSalaryForDeveloper(numberDeveloperEachCity, salaryDeveloperEachCity)
for _, a := range averageSalaryDeveloperByCity {
	fmt.Printf("%s : %d", a.Key, a.Value)
	fmt.Println("")
}
// // 2.9 
// fmt.Println("---------------AverageAgePerJob-------------------- ")
// sumAgeEachJob := SumAgeEachJob(people)
// numberPeopleEachJob := GroupPeopleByJob(people)
// averageAgePerJob := AverageAgePerJob(numberPeopleEachJob, sumAgeEachJob)
// for key, value := range averageAgePerJob {
// 	fmt.Printf("%s : %d", key, value)
// 	fmt.Println("")
// }
// // 2.10 
// fmt.Println("-----------AverageAgePerCity--------------- ")
// sumAgeEachCity := SumAgeEachCity(people)
// numberPeopleEachCity := NumberPeopleEachCity(people)
// averageAgePerCity := AverageAgePerCity(numberPeopleEachCity, sumAgeEachCity)
// for key, value := range averageAgePerCity {
// 	fmt.Printf("%s : %d", key, value)
// 	fmt.Println("")
// }
}



