This Go program demonstrates how to calculate the **future value** of an investment considering the **expected return rate** and **inflation rate** over a period of time. Let’s break down the code and explain the rules behind it.

### Code Breakdown:

1. **Constants and Variables:**
   ```go
   const inflationRate = 6.5
   ```
   - **Constant:** `inflationRate` is declared as a constant with the value `6.5`. In Go, constants are values that cannot change once they are set. Constants are evaluated at compile-time, and the `inflationRate` is a fixed value used in future calculations.
   
   ```go
   investmentAmount := 1000.0
   expectedReturnRate := 5.5
   years := 10.0
   ```
   - **Variables:**
     - `investmentAmount`: This is initialized to `1000.0` (float). This represents the initial amount of money being invested.
     - `expectedReturnRate`: This is initialized to `5.5`. This represents the expected annual return rate on the investment (5.5% per year).
     - `years`: This is initialized to `10.0`. It represents the number of years over which the investment will grow.
   
     Go uses **short variable declaration** (i.e., `:=`) here to automatically determine the type of these variables based on the assigned values. You don’t need to explicitly declare their types when using `:=` (although this is only allowed inside a function).

2. **Calculating Future Value (without considering inflation):**
   ```go
   futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
   ```
   - This line calculates the **future value** of the investment using the **compound interest formula**:
     \[
     \text{Future Value} = P \times (1 + r)^{n}
     \]
     where:
     - \( P \) = Initial investment (`investmentAmount`)
     - \( r \) = Expected return rate (as a decimal)
     - \( n \) = Number of years
     
   - `math.Pow(1 + expectedReturnRate/100, years)` raises `1 + expectedReturnRate/100` to the power of `years`. This is equivalent to applying compound interest for `years` number of years.
   
   **Example Calculation:**
   \[
   \text{Future Value} = 1000 \times (1 + 5.5/100)^{10} \approx 1708.14
   \]
   This is the value of the investment after 10 years, assuming no inflation.

3. **Calculating Future Real Value (adjusted for inflation):**
   ```go
   futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
   ```
   - This line adjusts the **future value** for inflation. It divides the future value by the inflation adjustment factor, which is also calculated using the **compound interest formula** but with the inflation rate:
     \[
     \text{Future Real Value} = \frac{\text{Future Value}}{(1 + \text{Inflation Rate})^n}
     \]
     
   - `math.Pow(1 + inflationRate/100, years)` calculates the inflation adjustment factor for the given number of years.
   - Dividing `futureValue` by this factor gives the **real value** of the investment, i.e., its value adjusted for inflation.

   **Example Calculation:**
   \[
   \text{Future Real Value} = \frac{1708.14}{(1 + 6.5/100)^{10}} \approx 909.97
   \]
   This represents the **purchasing power** of the future value in today’s money after accounting for inflation.

4. **Output:**
   ```go
   fmt.Println(futureValue)
   fmt.Println(futureRealValue)
   ```
   - The program prints both the **future value** and the **future real value**.
   
   The output will be:
   ```
   1708.1444583535929
   909.9730253950707
   ```
   - `1708.1444583535929`: This is the **future value** of the investment after 10 years, considering the expected return rate but not accounting for inflation.
   - `909.9730253950707`: This is the **future real value** of the investment after adjusting for inflation. It represents the actual value of the investment in today's terms.

### Key Concepts:

1. **Constant vs Variable:**
   - Constants (`inflationRate`) have fixed values that cannot be changed during the program execution.
   - Variables (`investmentAmount`, `expectedReturnRate`, `years`) hold values that can be modified during runtime.

2. **Short Variable Declaration (`:=`):**
   - Go allows the short variable declaration (`:=`) to automatically infer the type of a variable based on the value assigned to it. This makes Go code more concise and easier to write.

3. **Compound Interest Formula:**
   - The program uses the compound interest formula to calculate the future value of the investment over time. It accounts for both expected returns and inflation.

4. **Inflation Adjustment:**
   - Inflation reduces the purchasing power of money over time. The program calculates the future value of the investment in today’s terms by adjusting for inflation.

### Output Explanation:

- **Future Value:** This is the nominal future value of the investment, assuming the expected return rate, without considering inflation.
- **Future Real Value:** This adjusts the future value by inflation to give the "real" value of the investment in today's money.

### Conclusion:
This program demonstrates how to calculate the future value of an investment and adjust it for inflation to better understand its real worth in the future. It also showcases the use of constants, variables, and mathematical functions in Go.