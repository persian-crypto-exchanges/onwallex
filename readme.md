# OnWallex - Real-Time Currency Telegram Bot

## Overview

OnWallex is a real-time currency price tracking bot for Telegram. It uses the Wallex API to fetch and display the latest exchange rates for Tether (USDT), Bitcoin (BTC), and Ethereum (ETH) in USD and IRR. The bot automatically updates prices every hour and responds to user commands to provide on-demand updates.

## Features

Real-Time Currency Tracking: Fetches current prices for USDT, BTC, and ETH in both USD and IRR.
Scheduled Updates: Sends currency updates to a Telegram channel every hour.
Command-Based Updates: Responds to the /update command to fetch the latest prices on demand.
Formatted Output: Displays prices in a clean, user-friendly format.

## Installation

Prerequisites

1. Go installed on your system.
2. A valid Telegram Bot Token from BotFather.
3. A Telegram Channel to post the updates.
4. API access to Wallex.
5. Clone the repository:

```bash
git clone https://github.com/your-repo/onwallex.git
cd onwallex
```

6. Install dependencies:

```bash
go mod tidy
```

7. Replace placeholders in the code:

- Replace TELEGRAM_TOKEN with your actual Telegram Bot Token.
- Replace @YOUR_CHANNEL with your Telegram Channel's username.
- Build and run the application:

```bash
go run main.go
```

## Usage

Commands \
Hourly Updates: The bot automatically sends price updates to the configured Telegram channel every hour. \
Manual Updates: Users can type `/update` in the bot chat to fetch the latest prices.

## Bot Output Example

```bash
🔴 تتر
۸۱,۲۰۲ تومان

🟢 بیت کوین:
۹۶,۶۷۳ دلار
۷,۸۵۰,۱۰۲,۲۰۲ تومان

🟢 اتریوم:
۳,۴۶۴.۹ دلار
۲۸۱,۳۵۷,۴۰۵ تومان
```

## Dependencies

Gravel: A Go library for Telegram Bot API (https://github.com/gravelstone/gravel).

##

Enjoy real-time currency tracking with OnWallex! 🚀
