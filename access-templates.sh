#!/bin/bash
#
# Quick Access Script for Proposal & Contract Templates Pack
# This script provides direct access to the PDF file
#

echo "════════════════════════════════════════════════════════════"
echo "   Proposal & Contract Templates Pack - Chbuzorikenna"
echo "   Lagos, Nigeria"
echo "════════════════════════════════════════════════════════════"
echo ""

PDF_PATH="dist/Proposal-Contract-Templates-Pack.pdf"
HTML_PATH="dist/Proposal-Contract-Templates-Pack.html"

if [ -f "$PDF_PATH" ]; then
    echo "✅ PDF Found: $PDF_PATH"
    echo ""
    echo "File Details:"
    ls -lh "$PDF_PATH"
    echo ""
    echo "File Type:"
    file "$PDF_PATH"
    echo ""
    echo "════════════════════════════════════════════════════════════"
    echo "📥 To access the PDF:"
    echo "════════════════════════════════════════════════════════════"
    echo ""
    echo "Option 1: Direct Path"
    echo "  → $PDF_PATH"
    echo ""
    echo "Option 2: GitHub Repository"
    echo "  → https://github.com/DavidNzube101/project-phantom/blob/copilot/export-proposal-contract-templates/dist/Proposal-Contract-Templates-Pack.pdf"
    echo ""
    echo "Option 3: Download Raw File"
    echo "  → https://raw.githubusercontent.com/DavidNzube101/project-phantom/copilot/export-proposal-contract-templates/dist/Proposal-Contract-Templates-Pack.pdf"
    echo ""
    echo "════════════════════════════════════════════════════════════"
    echo "📖 What's Inside:"
    echo "════════════════════════════════════════════════════════════"
    echo ""
    echo "✓ Business Proposal Template"
    echo "  • Executive Summary"
    echo "  • Company Overview (Chbuzorikenna)"
    echo "  • Scope of Work & Deliverables"
    echo "  • Project Timeline"
    echo "  • Investment & Pricing (NGN ₦)"
    echo "  • Terms & Conditions"
    echo "  • Signature Blocks"
    echo ""
    echo "✓ Service Agreement Contract"
    echo "  • Comprehensive legal terms"
    echo "  • Nigerian jurisdiction (Lagos State)"
    echo "  • Payment terms with VAT (7.5%)"
    echo "  • Withholding tax provisions"
    echo "  • IP rights & confidentiality"
    echo "  • Dispute resolution clauses"
    echo "  • Witness signature blocks"
    echo ""
    echo "════════════════════════════════════════════════════════════"
    echo ""
    
    # Check if HTML version exists
    if [ -f "$HTML_PATH" ]; then
        echo "💡 HTML version also available: $HTML_PATH"
        echo ""
    fi
    
    echo "📚 For more information, see:"
    echo "  • TEMPLATES_README.md - Complete generation guide"
    echo "  • dist/README.md - Template usage instructions"
    echo ""
    echo "✨ Ready to use! Download and customize for your needs."
    echo ""
else
    echo "❌ PDF not found at $PDF_PATH"
    echo ""
    echo "To generate the PDF, run:"
    echo "  python3 generate_pdf.py"
    echo ""
fi

echo "════════════════════════════════════════════════════════════"
