# Proposal & Contract Templates Pack - Generation Guide

This guide explains how to generate the **Proposal & Contract Templates Pack** PDF for Chbuzorikenna (Lagos, Nigeria).

## 📦 What Gets Generated

The script generates a comprehensive PDF package containing:

1. **Business Proposal Template** - Professional proposal format with Nigerian business context
2. **Service Agreement Contract** - Legal contract template compliant with Nigerian laws

Both templates are tailored for:
- **Company:** Chbuzorikenna
- **Location:** Lagos, Nigeria
- **Currency:** Nigerian Naira (₦)
- **Legal Framework:** Federal Republic of Nigeria laws

## 🚀 Quick Start

### Generate the PDF

```bash
python3 generate_pdf.py
```

This will:
1. Read the markdown templates from `templates/` directory
2. Convert them to styled HTML
3. Generate a professional PDF
4. Output files to `dist/` directory

### Output Files

After running the script, you'll find:

- **dist/Proposal-Contract-Templates-Pack.pdf** - Main PDF file (ready to download)
- **dist/Proposal-Contract-Templates-Pack.html** - HTML version
- **dist/README.md** - Documentation for the templates

## 📂 Directory Structure

```
project-phantom/
├── generate_pdf.py              # PDF generation script
├── templates/                   # Source markdown templates
│   ├── proposal-template.md     # Business proposal template
│   └── contract-template.md     # Service contract template
└── dist/                        # Generated output files
    ├── Proposal-Contract-Templates-Pack.pdf
    ├── Proposal-Contract-Templates-Pack.html
    └── README.md
```

## 🔧 Requirements

### Python Dependencies

The script automatically installs required packages:
- `markdown` - Markdown to HTML conversion
- `weasyprint` - HTML to PDF generation

### System Requirements

- Python 3.6 or higher
- Internet connection (for first-time dependency installation)

## 📥 Downloading the PDF

### Option 1: Direct Download from Repository

1. Navigate to the `dist/` folder in the repository
2. Click on `Proposal-Contract-Templates-Pack.pdf`
3. Click the "Download" button

### Option 2: Clone and Generate Locally

```bash
# Clone the repository
git clone https://github.com/DavidNzube101/project-phantom.git
cd project-phantom

# Generate the PDF
python3 generate_pdf.py

# PDF will be in dist/Proposal-Contract-Templates-Pack.pdf
```

## ✏️ Customizing Templates

### Modify Template Content

1. Edit the source files in `templates/`:
   - `proposal-template.md` - For proposal changes
   - `contract-template.md` - For contract changes

2. Regenerate the PDF:
   ```bash
   python3 generate_pdf.py
   ```

### Customize Styling

Edit the CSS styles in `generate_pdf.py` (look for the `styled_html` section) to change:
- Colors
- Fonts
- Layout
- Spacing

## 🎨 Template Features

### Business Proposal Includes:
- Executive Summary
- Company Overview
- Scope of Work
- Project Timeline with milestones
- Investment & Pricing breakdown (₦)
- Payment Terms
- Terms & Conditions
- Signature blocks

### Service Contract Includes:
- Legal definitions and recitals
- Comprehensive scope of services
- Term and termination clauses
- Compensation and payment terms (₦)
- VAT (7.5%) and tax withholding provisions
- Intellectual property rights
- Confidentiality agreements
- Warranties and liability limitations
- Dispute resolution (Lagos jurisdiction)
- Multiple signature blocks with witnesses

## 🇳🇬 Nigerian Business Compliance

Templates include:
- **Currency:** Nigerian Naira (₦)
- **VAT:** 7.5% as per Nigerian tax regulations
- **Withholding Tax:** Compliance provisions
- **Jurisdiction:** Lagos State, Nigeria
- **Governing Law:** Federal Republic of Nigeria
- **Payment Methods:** Nigerian bank transfer details

## 📝 Usage Instructions

1. **Download the PDF** from `dist/Proposal-Contract-Templates-Pack.pdf`
2. **Open with a PDF editor** (Adobe Acrobat, Foxit, etc.)
3. **Fill in placeholders:**
   - Replace `[Client Name]` with actual client
   - Replace `[Date]` with current date
   - Fill in pricing amounts
   - Add project-specific details
   - Insert company logo where indicated
4. **Review legal terms** with a Nigerian legal professional
5. **Sign and execute** the documents

## 🔄 Troubleshooting

### PDF Generation Fails

If WeasyPrint fails to install:
```bash
# Try manual installation
pip install weasyprint

# Or use alternative
pip install markdown-pdf
```

### Missing Dependencies

```bash
# Install all dependencies manually
pip install markdown weasyprint
```

## ⚖️ Legal Notice

These templates are provided as starting points for business documentation. They should be:
- Reviewed by a qualified Nigerian legal professional
- Customized for your specific business needs
- Updated to reflect current Nigerian laws and regulations
- Adapted to your specific project requirements

## 📞 Support

For issues or questions about the templates:
- Check the `dist/README.md` for detailed template documentation
- Review the source templates in `templates/` directory
- Consult with a Nigerian business attorney for legal advice

## 📄 License

Templates are provided for business use by Chbuzorikenna and authorized parties.

---

**Generated for:** Chbuzorikenna  
**Location:** Lagos, Nigeria  
**Version:** 1.0  
**Last Updated:** February 2024
