package main

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Accepted format
// 31/12/2021, 13:59 - <Person who sent the msg>: <amount> <spent on>

var fileName string
func main(){
	fileName = "expenses2021.txt"
	ctx := context.Background()
	monthlyExpenses := processFile(ctx, fileName)
	annualExpense := 0
	for month, amount := range monthlyExpenses{
		log.Printf("spent %d in %s", amount, month)
		annualExpense = annualExpense+amount
	}
	log.Printf("spend %d in year ", annualExpense)
	/*input := "09/02/2021, 20:45 - Randeep: 238 sugar free"
	month, spenderName, amount, description, err := regexer(input)
	if err != nil {
		log.Println("error for " , input)
	}
	log.Printf("%s has spent %d in the month of %s for %s", spenderName, amount, month, description)*/
	log.Println(ctx, "completed")
}


func processFile(ctx context.Context, fileNameStr string) map[string]int{
	file, err := os.Open(fileNameStr)
	if err != nil {
		log.Fatal(ctx, err)
	}
	defer file.Close()
	monthExpenseMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	var errLineCount int
	var totalLineCount int
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		//txtlines = append(txtlines, scanner.Text())
		currLine := scanner.Text()
		if strings.Contains(currLine, "deleted"){
			log.Println("this line was deleted ")
			continue
		}
		monthYear, spenderName, amount, description, err := regexer(currLine)
		if err != nil {
			log.Printf("error %s for " , err.Error(), scanner.Text())
			errLineCount = errLineCount +1
			continue
		}
		log.Printf("%s has spent %d in the month of %s for %s", spenderName, amount, monthYear, description)
		if totalAmount, ok := monthExpenseMap[monthYear]; !ok {
			monthExpenseMap[monthYear] = amount
		} else {
			monthExpenseMap[monthYear] = amount+totalAmount
		}
		totalLineCount = totalLineCount+1
	}
	//log.Printf(" number of lines %d ", len(txtlines))

	if err := scanner.Err(); err != nil {
		log.Fatal(ctx, err)
	}
	log.Printf("errors in %d lines ", errLineCount)
	return monthExpenseMap
}
/*func stringSplitter(str string) (int, []string, error){

}*/

func regexer(str string) (string, string, int, string, error){
	regexpPattern := "\\d{2}\\/(\\d{2})\\/(\\d{4}),\\s+\\d{2}:\\d{2}\\s+-\\s+(\\w+):\\s+(\\d+)\\s+(.*)"
	var month string
	var spenderName string
	var amount int
	var description string
	var year string
	monthYear := month+"_"+year

	re, err := regexp.Compile(regexpPattern)
	if err != nil {
		log.Println("error ", err)
		return month, spenderName, amount, description, err
	}
	output := re.FindAllStringSubmatch(str, -1)
	if len(output)== 0{
		return month, spenderName, amount, description, errors.New("err in line parsing")
	}
	if len(output[0])<5 {
		return month, spenderName, amount, description, errors.New("err in line parsing")
	}

	month   = monthConvertor(output[0][1])
	year    = output[0][2]
	spenderName = output[0][3]
	amount, err = strconv.Atoi(output[0][4])
	if err != nil {
		return monthYear, spenderName, amount, description, err
	}
	monthYear = month+"_"+year
	description = output[0][5]
	return monthYear, spenderName, amount, description, err
}

func monthConvertor(str string) string {
	jan := "January"
	feb := "February"
	mar := "March"
	apr := "April"
	may := "May"
	jun := "June"
	jul := "July"
	aug := "August"
	sep := "September"
	oct := "October"
	nov := "November"
	dec := "December"
	switch str {
	case "01":
		return jan
	case "02":
		return feb
	case "03":
		return mar
	case "04":
		return apr
	case "05":
		return may
	case "06":
		return jun
	case "07":
		return jul
	case "08":
		return aug
	case "09":
		return sep
	case "10":
		return oct
	case "11":
		return nov
	case "12":
		return dec
	}
	return "month invalid"
}