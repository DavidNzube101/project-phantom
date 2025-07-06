# Firebase Setup Guide

## Prerequisites
1. Firebase project created at https://console.firebase.google.com
2. Google Cloud CLI installed (optional, for local development)

## Setup Options

### Option 1: Service Account Key (Recommended for Production)
1. Go to Firebase Console > Project Settings > Service Accounts
2. Click "Generate new private key"
3. Download the JSON file
4. Set environment variable:
   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/service-account-key.json"
   ```

### Option 2: Google Cloud CLI (For Development)
1. Install Google Cloud CLI
2. Run authentication:
   ```bash
   gcloud auth application-default login
   ```
3. Set your project:
   ```bash
   gcloud config set project project-phantom-7f30e
   ```

### Option 3: Environment Variables
Set these environment variables:
```bash
export FIREBASE_PROJECT_ID="project-phantom-7f30e"
export FIREBASE_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n"
export FIREBASE_CLIENT_EMAIL="firebase-adminsdk-xxxxx@project-phantom-7f30e.iam.gserviceaccount.com"
```

## Firestore Database Setup
1. Go to Firebase Console > Firestore Database
2. Create database if not exists
3. Start in test mode (for development)
4. The collection `user_settings` will be created automatically when first user data is saved

## Testing the API
Once the server is running, you can test the endpoints:

### Save User Settings
```bash
curl -X POST http://localhost:3000/user-settings \
  -H "Content-Type: application/json" \
  -d '{
    "walletAddress": "0x1234567890abcdef",
    "settings": {
      "theme": "dark",
      "notifications": true,
      "language": "en"
    }
  }'
```

### Get User Settings
```bash
curl http://localhost:3000/user-settings/0x1234567890abcdef
```

### Update User Settings
```bash
curl -X PUT http://localhost:3000/user-settings/0x1234567890abcdef \
  -H "Content-Type: application/json" \
  -d '{
    "theme": "light",
    "newSetting": "value"
  }'
```

### Delete User Settings
```bash
curl -X DELETE http://localhost:3000/user-settings/0x1234567890abcdef
``` 