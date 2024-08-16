# Reconciliation Service

## Overview
This project implements a transaction reconciliation service to identify unmatched and discrepant transactions between internal system data and external bank statements for xxxx. The service helps pinpoint errors, discrepancies, and missing transactions.

## Problem Statement
xxxx manages multiple bank accounts and requires a service to reconcile transactions within their system against corresponding transactions in bank statements. The objective is to detect and resolve discrepancies between the two sources.

## Data Model

### Transaction
- **trxID**: Unique identifier for the transaction (string)
- **amount**: Transaction amount (decimal)
- **type**: Transaction type (enum: DEBIT, CREDIT)
- **transactionTime**: Date and time of the transaction (datetime)

### Bank Statement
- **unique_identifier**: Unique identifier for the transaction in the bank statement (string, varies by bank; not necessarily equivalent to trxID)
- **amount**: Transaction amount (decimal; can be negative for debits)
- **date**: Date of the transaction (date)

## Assumptions
- System transactions and bank statements are provided as separate CSV files.
- Discrepancies only occur in amounts.

## Functionality

### Input Parameters
The service accepts the following input parameters:
- **System Transaction CSV File Path**: Path to the CSV file containing system transactions.
- **Bank Statement CSV File Path**: Path to one or more CSV files containing bank statements from different banks.
- **Start Date for Reconciliation Timeframe**: Start date for the reconciliation process (date).
- **End Date for Reconciliation Timeframe**: End date for the reconciliation process (date).

### Process
1. The service compares transactions within the specified timeframe across system and bank statement data.
2. It performs the reconciliation by identifying matches and discrepancies.

### Output
The reconciliation service provides a summary containing:
- **Total Number of Transactions Processed**
- **Total Number of Matched Transactions**
- **Total Number of Unmatched Transactions**
- **Details of Unmatched Transactions**:
  - System transaction details if missing in bank statements.
  - Bank statement details if missing in system transactions (grouped by bank).
- **Total Discrepancies**: Sum of absolute differences in amounts between matched transactions.
