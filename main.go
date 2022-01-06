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

/* The beginning
 */
func main() {
	fileName = "expenses2021.txt"
	ctx := context.Background()
	expenseData := processFile(ctx, fileName)
	WriteToDB(expenseData)
	log.Println("completed")
}

/*
	monthExpenseMap := make(map[string]int)
	spenderExpenseMap := make(map[string]int)
	//log.Printf("%s has spent %d in the month of %s for %s", spenderName, amount, monthYear, description)
	if totalAmount, ok := monthExpenseMap[monthYear]; !ok {
		monthExpenseMap[monthYear] = amount
	} else {
		monthExpenseMap[monthYear] = amount + totalAmount
	}
	if spenderTotal, ok := spenderExpenseMap[spenderName]; !ok {
		spenderExpenseMap[spenderName] = amount
	} else {
		spenderExpenseMap[spenderName] = amount + spenderTotal
	}

	monthlyExpenses, spenderExpenses := processFile(ctx, fileName)
	annualExpense := 0
	for month, amount := range monthlyExpenses {
		log.Printf("spent %d in %s", amount, month)
		annualExpense = annualExpense + amount
	}
	log.Printf("spend %d in year \n", annualExpense)
	for name, amount := range spenderExpenses {
		log.Printf("%s has spent %d \n", name, amount)
		annualExpense = annualExpense + amount
	}
 */

type ExpenseData struct {
	month       string
	year        string
	spenderName string
	amount      float64
	description  string
}

/*
Read the input file, process to a data structure and returns
*/
func processFile(ctx context.Context, fileNameStr string) []ExpenseData {
	file, err := os.Open(fileNameStr)
	if err != nil {
		log.Fatal(ctx, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var errLineCount int
	var totalLineCount int
	var expenseData []ExpenseData
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		//txtlines = append(txtlines, scanner.Text())
		currLine := scanner.Text()
		if strings.Contains(currLine, "deleted") {
			//log.Println("this line was deleted ")
			continue
		}
		expense, err := regexer(currLine)
		if err != nil {
			log.Printf("error %s for ", err.Error(), scanner.Text())
			errLineCount = errLineCount + 1
			continue
		}
		expenseData = append(expenseData, expense)
		totalLineCount = totalLineCount + 1
	}
	//log.Printf(" number of lines %d ", len(txtlines))

	if err := scanner.Err(); err != nil {
		log.Fatal(ctx, err)
	}
	log.Printf("errors in %d lines ", errLineCount)
	return expenseData
}

func WriteToDB(data []ExpenseData) {
	
}

func regexer(str string) (ExpenseData, error) {
	regexpPattern := "\\d{2}\\/(\\d{2})\\/(\\d{4}),\\s+\\d{2}:\\d{2}\\s+-\\s+(\\w+):\\s+(\\d+)\\s+(.*)"
	var month string
	var spenderName string
	var amount float64
	var description string
	var year string
	var expense ExpenseData
	re, err := regexp.Compile(regexpPattern)
	if err != nil {
		log.Println("error ", err)
		return expense, err
	}
	output := re.FindAllStringSubmatch(str, -1)
	if len(output) == 0 {
		return expense, errors.New("err in line parsing")
	}
	if len(output[0]) < 5 {
		return expense, errors.New("err in line parsing")
	}

	month = output[0][1]
	year = output[0][2]
	spenderName = output[0][3]
	amount, err = strconv.ParseFloat(output[0][4], 64)
	if err != nil {
		return expense, err
	}
	description = output[0][5]
	expense.month = month
	expense.year = year
	expense.spenderName = spenderName
	expense.amount = amount
	expense.description = description
	return expense, err
}

//monthStr = monthConvertor(month)
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
