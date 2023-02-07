
# Telegram bot for Welly 1:60 toy cars collections

Bot keeps items in file (by default) or connect to postgres.

Telegram used as UI.

## Features

- add
- delete
- find item by itemID, Model or Manufacture name contains
- show description and photos

## Documentation

### Item:

- Model name
- Manufacture name
- Welly ID
- Color
- Comments
- Title photo
- []Photos 


## Usage/Examples

Telegram Bots Token is in ENV (WELLY_TOKEN), or in .env file (TOKEN) or use flag `--token`

Postgres URL is in ENV (DATABASE_URL) .env file (DATABASE_URL) or use flag `--pgurl`

Bot operates here:
```javascript
https://t.me/collectionist_bot
```

## Screenshots

| Search  | Show with photos | Add new |
| ------------- | ------------- | ------------- |
| ![App Screenshot](https://github.com/kormiltsev/tbot-welly/blob/main/etc/s1.png "Search") | ![App Screenshot](https://github.com/kormiltsev/tbot-welly/blob/main/etc/s2.png "Show photo") | ![App Screenshot](https://github.com/kormiltsev/tbot-welly/blob/main/etc/s3.png "Add new") |

