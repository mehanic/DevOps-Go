package interface_module

import "fmt"

type ICredits interface {
	CalculateInterestAmount() float64
	CalculateMonthlyPayment() float64
	PrintInterest()
	PrintMontlyPayment()
}

type Account struct {
	accountName     string
	accountOwner    string
	accountCurrency string
	accountBalance  float64
}

type PersonalCredit struct {
	CreditOwner        string
	CreditTimeline     int
	CreditCurrency     string
	CreditAmount       int
	CreditInterestRate float64
}

type AutoCredit struct {
	CreditOwner        string
	CreditTimeline     int
	CreditCurrency     string
	CreditAmount       int
	CreditInterestRate float64
}

type MortgageCredit struct {
	CreditOwner        string
	CreditTimeline     int
	CreditCurrency     string
	CreditAmount       int
	CreditInterestRate float64
}

func (account *Account) Deposit(amount float64) {
	account.accountBalance += amount
	fmt.Printf("%.2f %s added to your %s account, new balance of your account is %.2f %s.\n",
		amount, account.accountCurrency, account.accountName, account.accountBalance, account.accountCurrency)
}

func (account *Account) Withdraw(amount float64) {
	account.accountBalance -= amount
	fmt.Printf("%.2f %s withdrawed from your %s account, new balance of your account is %.2f %s.\n",
		amount, account.accountCurrency, account.accountName, account.accountBalance, account.accountCurrency)
}

func AccountTransaction() {

	/*
		a := Account{"Saving", "Berkay Alan", "Euro", 300}
		fmt.Printf("Welcome to Gitbank Sir %s!\n", a.accountOwner)
		a.Deposit(45)
		a.Withdraw(74)
		fmt.Printf("The balance of your account is %.2f %s. See you next time!\n", a.accountBalance, a.accountCurrency)

		fmt.Print("------------------------\n")

		p := PersonalCredit{"Kelvin Yosi", 36, "Euro", 40000, 0.69}
		pc_interest := p.CalculateInterest()
		fmt.Printf("For getting %d %s credit, you will pay %.2f %s in %d month.\nMonthly payment is %.2f.",
			p.CreditAmount, p.CreditCurrency, pc_interest, p.CreditCurrency, p.CreditTimeline, pc_interest/float64(p.CreditTimeline))

	*/

	mortgage := MortgageCredit{"Pascal Oka", 48, "Dollar", 300000, 1.10}

	mortgage.PrintInterest()

	mortgage.PrintMontlyPayment()

	auto := AutoCredit{"Anna Keck", 12, "Euro", 20000, 0.76}

	auto.PrintInterest()

	auto.PrintMontlyPayment()

	personal := PersonalCredit{"Bruno Alvez", 12, "Euro", 10000, 0.46}

	personal.PrintInterest()

	personal.PrintMontlyPayment()

}

func (personal *PersonalCredit) PrintInterest() {

	pe_interest := personal.CalculateInterest()

	fmt.Printf("For getting %d %s personal credit, you will pay %.2f %s in %d month.\n",
		personal.CreditAmount, personal.CreditCurrency, pe_interest, personal.CreditCurrency, personal.CreditTimeline)
}

func (personal *PersonalCredit) PrintMontlyPayment() {

	pe_interest := personal.CalculateInterest()
	fmt.Printf("Your monthly payment will be %.2f %s.\n",
		pe_interest/float64(personal.CreditTimeline), personal.CreditCurrency)

}

func (personal *PersonalCredit) CalculateInterest() float64 {

	interest := personal.CreditInterestRate * float64(personal.CreditAmount)
	return float64(personal.CreditAmount) + interest
}

func (personal *PersonalCredit) CalculateMonthlyPayment() float64 {

	interest := (personal.CreditInterestRate * float64(personal.CreditAmount)) + float64(personal.CreditAmount)
	return interest / float64(personal.CreditTimeline)
}

//////////////////////////////////////////////////////

func (auto *AutoCredit) PrintInterest() {

	au_interest := auto.CalculateInterest()

	fmt.Printf("For getting %d %s auto credit, you will pay %.2f %s in %d month.\n",
		auto.CreditAmount, auto.CreditCurrency, au_interest, auto.CreditCurrency, auto.CreditTimeline)
}

func (auto *AutoCredit) PrintMontlyPayment() {

	au_interest := auto.CalculateInterest()
	fmt.Printf("Your monthly payment will be %.2f %s.\n",
		au_interest/float64(auto.CreditTimeline), auto.CreditCurrency)

}

func (auto *AutoCredit) CalculateInterest() float64 {

	interest := auto.CreditInterestRate * float64(auto.CreditAmount)
	return float64(auto.CreditAmount) + interest
}

func (auto *AutoCredit) CalculateMonthlyPayment() float64 {

	interest := (auto.CreditInterestRate * float64(auto.CreditAmount)) + float64(auto.CreditAmount)
	return interest / float64(auto.CreditTimeline)
}

//////////////////////////////////////////////////////

func (mortgage *MortgageCredit) PrintInterest() {

	mg_interest := mortgage.CalculateInterest()

	fmt.Printf("For getting %d %s mortgage credit, you will pay %.2f %s in %d month.\n",
		mortgage.CreditAmount, mortgage.CreditCurrency, mg_interest, mortgage.CreditCurrency, mortgage.CreditTimeline)
}

func (mortgage *MortgageCredit) PrintMontlyPayment() {

	mg_interest := mortgage.CalculateInterest()
	fmt.Printf("Your monthly payment will be %.2f %s.\n",
		mg_interest/float64(mortgage.CreditTimeline), mortgage.CreditCurrency)

}

func (mortgage *MortgageCredit) CalculateInterest() float64 {

	interest := mortgage.CreditInterestRate * float64(mortgage.CreditAmount)
	return float64(mortgage.CreditAmount) + interest
}

func (mortgage *MortgageCredit) CalculateMonthlyPayment() float64 {

	interest := (mortgage.CreditInterestRate * float64(mortgage.CreditAmount)) + float64(mortgage.CreditAmount)
	return interest / float64(mortgage.CreditTimeline)
}
