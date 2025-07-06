# Project Phantom API Documentation

## Overview

Project Phantom is a privacy-first mobile browser backend API that provides user settings management with minimal data collection. Users are identified solely by their wallet address, ensuring maximum privacy and anonymity.

## Privacy-First Design Principles

- **No Personal Data**: Users are identified only by wallet address
- **Minimal Data Collection**: Only essential user preferences are stored
- **No Tracking**: No analytics, user behavior monitoring, or tracking
- **Flexible Schema**: Settings can accommodate future privacy features
- **User Control**: Users can delete their data at any time

## Base URL

- **Development**: `http://localhost:3000`
- **Production**: `https://api.project-phantom.com`

## Authentication

No traditional authentication is required. Users are identified by their wallet address, which serves as both the identifier and the document ID in the database.

## Interactive Documentation

Access the interactive Swagger documentation at:
- **Development**: `http://localhost:3000/api`
- **Production**: `https://api.project-phantom.com/api`

## Endpoints

### Health Check

#### GET /
Check if the API is running.

**Response:**
```json
"Hello World!"
```

### User Settings

#### POST /user-settings
Create or update user settings.

**Request Body:**
```json
{
  "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
  "settings": {
    "theme": "dark",
    "notifications": true,
    "language": "en",
    "privacyLevel": "high",
    "autoBlockTrackers": true
  }
}
```

**Response (201 Created):**
```json
{
  "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
  "settings": {
    "theme": "dark",
    "notifications": true,
    "language": "en",
    "privacyLevel": "high",
    "autoBlockTrackers": true
  },
  "createdAt": "2024-01-15T10:30:00.000Z",
  "updatedAt": "2024-01-15T10:30:00.000Z"
}
```

#### GET /user-settings/{walletAddress}
Retrieve user settings by wallet address.

**Parameters:**
- `walletAddress` (string, required): User's wallet address (42 characters)

**Response (200 OK):**
```json
{
  "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
  "settings": {
    "theme": "dark",
    "notifications": true,
    "language": "en"
  },
  "createdAt": "2024-01-15T10:30:00.000Z",
  "updatedAt": "2024-01-15T10:30:00.000Z"
}
```

**Response (404 Not Found):**
```json
{
  "statusCode": 404,
  "message": "User settings not found for wallet address: 0x1234567890abcdef1234567890abcdef12345678",
  "error": "Not Found"
}
```

#### PUT /user-settings/{walletAddress}
Update existing user settings (merges with existing settings).

**Parameters:**
- `walletAddress` (string, required): User's wallet address (42 characters)

**Request Body:**
```json
{
  "settings": {
    "theme": "light",
    "newSetting": "value"
  }
}
```

**Response (200 OK):**
```json
{
  "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
  "settings": {
    "theme": "light",
    "notifications": true,
    "language": "en",
    "newSetting": "value"
  },
  "createdAt": "2024-01-15T10:30:00.000Z",
  "updatedAt": "2024-01-15T10:35:00.000Z"
}
```

#### DELETE /user-settings/{walletAddress}
Permanently delete user settings.

**Parameters:**
- `walletAddress` (string, required): User's wallet address (42 characters)

**Response (204 No Content):**
No response body

## Error Responses

### 400 Bad Request
Invalid request data (e.g., invalid wallet address format, malformed JSON)

### 404 Not Found
User settings not found for the specified wallet address

### 500 Internal Server Error
Server error (check server logs for details)

## Data Schema

### User Settings Object
The settings object is flexible and can contain any user preferences. Common settings include:

```json
{
  "theme": "dark" | "light" | "auto",
  "notifications": boolean,
  "language": "en" | "es" | "fr" | "de" | "ja" | "ko" | "zh",
  "privacyLevel": "low" | "medium" | "high" | "maximum",
  "autoBlockTrackers": boolean,
  "javascriptEnabled": boolean,
  "cookiesEnabled": boolean,
  "defaultSearchEngine": "google" | "duckduckgo" | "bing" | "yahoo",
  "startPage": "homepage" | "blank" | "custom",
  "customStartPage": "https://example.com",
  "bookmarks": [
    {
      "title": "Example Site",
      "url": "https://example.com",
      "icon": "data:image/png;base64,..."
    }
  ]
}
```

## Rate Limiting

Currently, no rate limiting is implemented. Consider implementing rate limiting for production use.

## CORS

CORS is enabled for all origins in development. Configure appropriate CORS settings for production.

## Firebase Configuration

The API uses Firebase Firestore for data storage. Ensure proper Firebase configuration as described in `FIREBASE_SETUP.md`.

## Development

### Running the API
```bash
# Install dependencies
yarn install

# Start development server
yarn start:dev

# Build for production
yarn build

# Start production server
yarn start:prod
```

### Testing
```bash
# Run tests
yarn test

# Run tests in watch mode
yarn test:watch

# Run e2e tests
yarn test:e2e
```

## Privacy Considerations

1. **No Logging**: The API does not log user requests or personal data
2. **Data Minimization**: Only essential settings are stored
3. **User Control**: Users can delete their data at any time
4. **No Analytics**: No user behavior tracking or analytics
5. **Secure Storage**: Data is stored in Firebase Firestore with proper security rules

## Future Enhancements

- Encrypted settings storage
- Settings sync across devices
- Privacy-focused analytics (opt-in, anonymous)
- Rate limiting
- API versioning
- Webhook support for settings changes 