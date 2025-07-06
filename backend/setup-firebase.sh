#!/bin/bash

echo "🚀 Setting up Firebase for Project Phantom Backend"
echo "=================================================="

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo "❌ Google Cloud CLI is not installed."
    echo "Please install it first: https://cloud.google.com/sdk/docs/install"
    echo ""
    echo "For Arch Linux: sudo pacman -S google-cloud-cli"
    echo "For Ubuntu/Debian: curl https://sdk.cloud.google.com | bash"
    echo "For macOS: brew install google-cloud-sdk"
    exit 1
fi

echo "✅ Google Cloud CLI found"

# Check if user is authenticated
if ! gcloud auth list --filter=status:ACTIVE --format="value(account)" | grep -q .; then
    echo "🔐 You need to authenticate with Google Cloud"
    echo "This will open a browser window for authentication..."
    gcloud auth login
else
    echo "✅ Already authenticated with Google Cloud"
fi

# Set the project
echo "📁 Setting project to: project-phantom-7f30e"
gcloud config set project project-phantom-7f30e

# Set up application default credentials
echo "🔑 Setting up application default credentials..."
gcloud auth application-default login

echo ""
echo "✅ Firebase setup complete!"
echo ""
echo "You can now start the server with:"
echo "  yarn start:dev"
echo ""
echo "Test the API with:"
echo "  curl -X POST http://localhost:3000/user-settings \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"walletAddress\": \"56wXYarJLA8pFysi81syH6yABRKV3si4V75FQJHa3umi\", \"settings\": {\"theme\": \"dark\"}}'"
echo ""
echo "📚 Swagger docs will be available at: http://localhost:3000/api" 