# Billing Engine for Loan Management

## Overview
The billing engine is designed to manage and track loans, providing features such as loan schedules, outstanding balances, and delinquency status. This system supports a 50-week loan with a flat interest rate of 10% per annum.

## Features
1. **Loan Schedule**: Provides the repayment schedule for a given loan.
2. **Outstanding Amount**: Tracks the outstanding balance of the loan.
3. **Delinquency Status**: Determines if a borrower is delinquent based on missed payments.

## Loan Details
- **Loan Amount**: Rp 5,000,000
- **Interest Rate**: 10% per annum
- **Repayment Duration**: 50 weeks

### Repayment Schedule
For a loan amount of Rp 5,000,000 over 50 weeks at a 10% annual interest rate, the repayment schedule is as follows:
- **W1**: Rp 110,000
- **W2**: Rp 110,000
- **W3**: Rp 110,000
- ...
- **W50**: Rp 110,000

### Payment Rules
- The borrower repays the exact amount due each week or does not pay at all.
- Outstanding balance decreases with each repayment.
- At the end of the loan period, the outstanding balance should be zero.

### Delinquency
- A borrower is considered delinquent if they miss 2 consecutive repayments.
- If repayments are missed, the borrower must make up for the missed payments in addition to the regular repayments.

## We are looking for at least the following methods to be implemented

### `GetOutstanding`
- **Description**: Returns the current outstanding balance on a loan. Returns 0 if the loan is closed.

### `IsDelinquent`
- **Description**: Checks if the borrower has missed more than 2 consecutive repayments.

### `MakePayment`
- **Description**: Processes a payment of a certain amount on the loan.

