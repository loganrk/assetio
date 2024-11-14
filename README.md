# ASSETIO

## Overview

This project provides a set of services for managing accounts, securities, stocks, and mutual funds. The project defines a set of interfaces and methods for creating, updating, retrieving, and processing data related to these entities. The system allows interactions with the stock and mutual fund portfolios, including operations like buying/selling stocks, adding dividends, managing inventories, and more.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Features](#features)

## Project Structure

The project consists of the following main components:

### 1. Services

- **AccountSvr**: Interface for client account management.
- **SecuritySvr**: Interface for managing securities.
- **StockSvr**: Interface for managing stock-related operations.
- **MutualFundSvr**: Interface for managing mutual fund operations.

### 2. Request Types

- Structures that define the data required for each operation (e.g., `ClientAccountCreateRequest`, `ClientStockBuyRequest`).

### 3. Response Interface

- Handles responses with methods to set error codes, status codes, and data (e.g., `SetError`, `SetStatus`, `SetData`, `Send`).

### Main Service Interfaces

#### **AccountSvr**: Operations related to client account management.
- `AccountCreate`
- `AccountAll`
- `AccountGet`
- `AccountActivate`
- `AccountInactivate`
- `AccountUpdate`

#### **SecuritySvr**: Operations for managing security-related actions.
- `SecurityCreate`
- `SecurityAll`
- `SecurityGet`
- `SecuritySearch`
- `SecurityUpdate`

#### **StockSvr**: Operations for managing stock-related actions.
- `StockBuy`
- `StockSell`
- `StockDividendAdd`
- `StockSplit`
- `StockSummary`
- `StockInventories`
- `StockInventoryLedgers`

#### **MutualFundSvr**: Operations related to mutual fund management.
- `MutualFundBuy`
- `MutualFundAdd`
- `MutualFundSell`
- `MutualFundSummary`
- `MutualFundInventory`
- `MutualFundInventoryLedgers`

---

## Features

- **Account Management**: Allows creating, retrieving, updating, activating, and inactivating accounts.
- **Security Management**: Enables creating, updating, and retrieving securities.
- **Stock Transactions**: Facilitates buying, selling, and managing stocks, as well as adding dividends and performing stock splits.
- **Mutual Fund Management**: Allows users to buy and sell mutual funds, and track their summaries and inventories.

---