#!/usr/bin/env python3
"""
PDF Generator for Proposal & Contract Templates Pack
Tailored for Chbuzorikenna - Lagos, Nigeria
"""

import os
import sys
from pathlib import Path

def generate_pdf():
    """Generate PDF from markdown templates"""
    
    # Define paths
    repo_root = Path(__file__).parent
    templates_dir = repo_root / "templates"
    dist_dir = repo_root / "dist"
    output_pdf = dist_dir / "Proposal-Contract-Templates-Pack.pdf"
    
    # Ensure dist directory exists
    dist_dir.mkdir(exist_ok=True)
    
    print("🚀 Generating Proposal & Contract Templates Pack PDF...")
    print(f"📁 Templates directory: {templates_dir}")
    print(f"📄 Output PDF: {output_pdf}")
    
    # Check if markdown files exist
    proposal_file = templates_dir / "proposal-template.md"
    contract_file = templates_dir / "contract-template.md"
    
    if not proposal_file.exists():
        print(f"❌ Error: Proposal template not found at {proposal_file}")
        return False
    
    if not contract_file.exists():
        print(f"❌ Error: Contract template not found at {contract_file}")
        return False
    
    print("✅ Found both template files")
    
    # Read markdown content
    print("📖 Reading template files...")
    with open(proposal_file, 'r', encoding='utf-8') as f:
        proposal_content = f.read()
    
    with open(contract_file, 'r', encoding='utf-8') as f:
        contract_content = f.read()
    
    # Combine templates with separator
    combined_content = f"""{proposal_content}

---

<div style="page-break-before: always;"></div>

---

{contract_content}"""
    
    # Try to import markdown and pdf generation libraries
    try:
        import markdown
        print("✅ Markdown library available")
    except ImportError:
        print("⚙️  Installing markdown library...")
        os.system(f"{sys.executable} -m pip install markdown --quiet")
        import markdown
    
    # Convert markdown to HTML
    print("🔄 Converting markdown to HTML...")
    html_content = markdown.markdown(
        combined_content,
        extensions=['tables', 'nl2br', 'sane_lists']
    )
    
    # Add CSS styling for professional appearance
    styled_html = f"""<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Proposal & Contract Templates Pack - Chbuzorikenna</title>
    <style>
        @page {{
            size: A4;
            margin: 2cm;
        }}
        
        body {{
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }}
        
        h1 {{
            color: #1a1a1a;
            border-bottom: 3px solid #0066cc;
            padding-bottom: 10px;
            margin-top: 30px;
        }}
        
        h2 {{
            color: #0066cc;
            margin-top: 25px;
            border-bottom: 1px solid #ccc;
            padding-bottom: 5px;
        }}
        
        h3 {{
            color: #333;
            margin-top: 20px;
        }}
        
        table {{
            width: 100%;
            border-collapse: collapse;
            margin: 20px 0;
        }}
        
        table th {{
            background-color: #0066cc;
            color: white;
            padding: 12px;
            text-align: left;
            font-weight: bold;
        }}
        
        table td {{
            padding: 10px;
            border: 1px solid #ddd;
        }}
        
        table tr:nth-child(even) {{
            background-color: #f9f9f9;
        }}
        
        hr {{
            border: none;
            border-top: 2px solid #0066cc;
            margin: 30px 0;
        }}
        
        strong {{
            color: #0066cc;
        }}
        
        ul, ol {{
            margin: 15px 0;
            padding-left: 30px;
        }}
        
        li {{
            margin: 8px 0;
        }}
        
        .page-break {{
            page-break-before: always;
        }}
        
        p {{
            margin: 12px 0;
        }}
        
        code {{
            background-color: #f4f4f4;
            padding: 2px 6px;
            border-radius: 3px;
            font-family: 'Courier New', monospace;
        }}
        
        blockquote {{
            border-left: 4px solid #0066cc;
            padding-left: 20px;
            margin: 20px 0;
            font-style: italic;
            color: #666;
        }}
    </style>
</head>
<body>
{html_content}
</body>
</html>"""
    
    # Save HTML version
    html_file = dist_dir / "Proposal-Contract-Templates-Pack.html"
    with open(html_file, 'w', encoding='utf-8') as f:
        f.write(styled_html)
    print(f"✅ HTML version saved: {html_file}")
    
    # Try to generate PDF using weasyprint
    pdf_generated = False
    
    try:
        import weasyprint
        print("✅ WeasyPrint available, generating PDF...")
        weasyprint.HTML(string=styled_html).write_pdf(str(output_pdf))
        pdf_generated = True
        print(f"✅ PDF generated successfully: {output_pdf}")
    except ImportError:
        print("⚙️  Installing WeasyPrint for PDF generation...")
        try:
            os.system(f"{sys.executable} -m pip install weasyprint --quiet")
            import weasyprint
            weasyprint.HTML(string=styled_html).write_pdf(str(output_pdf))
            pdf_generated = True
            print(f"✅ PDF generated successfully: {output_pdf}")
        except Exception as e:
            print(f"⚠️  WeasyPrint installation failed: {e}")
    
    # If weasyprint failed, try pdfkit
    if not pdf_generated:
        try:
            import pdfkit
            print("✅ PDFKit available, generating PDF...")
            pdfkit.from_string(styled_html, str(output_pdf))
            pdf_generated = True
            print(f"✅ PDF generated successfully: {output_pdf}")
        except ImportError:
            print("⚙️  Installing pdfkit for PDF generation...")
            try:
                os.system(f"{sys.executable} -m pip install pdfkit --quiet")
                import pdfkit
                pdfkit.from_string(styled_html, str(output_pdf))
                pdf_generated = True
                print(f"✅ PDF generated successfully: {output_pdf}")
            except Exception as e:
                print(f"⚠️  PDFKit failed: {e}")
    
    # If both failed, try markdown-pdf via node
    if not pdf_generated:
        print("⚙️  Trying alternative PDF generation method...")
        try:
            # Save combined markdown
            combined_md = dist_dir / "combined-templates.md"
            with open(combined_md, 'w', encoding='utf-8') as f:
                f.write(combined_content)
            
            # Try using pandoc if available
            import subprocess
            result = subprocess.run(
                ['which', 'pandoc'],
                capture_output=True,
                text=True
            )
            
            if result.returncode == 0:
                print("✅ Pandoc available, generating PDF...")
                subprocess.run([
                    'pandoc',
                    str(combined_md),
                    '-o', str(output_pdf),
                    '--pdf-engine=xelatex',
                    '-V', 'geometry:margin=2cm'
                ], check=True)
                pdf_generated = True
                print(f"✅ PDF generated successfully: {output_pdf}")
            else:
                print("⚠️  Pandoc not available")
        except Exception as e:
            print(f"⚠️  Alternative method failed: {e}")
    
    if pdf_generated:
        print("\n" + "="*60)
        print("✨ SUCCESS! PDF generated successfully!")
        print("="*60)
        print(f"\n📄 PDF Location: {output_pdf}")
        print(f"📄 HTML Version: {html_file}")
        print(f"\n📦 File size: {output_pdf.stat().st_size / 1024:.2f} KB")
        print("\n✅ The Proposal & Contract Templates Pack is ready!")
        print("   Tailored for: Chbuzorikenna, Lagos, Nigeria")
        return True
    else:
        print("\n" + "="*60)
        print("⚠️  PDF generation failed with all methods")
        print("="*60)
        print(f"\n📄 HTML version is available at: {html_file}")
        print("\nYou can:")
        print("  1. Open the HTML file in a browser and print to PDF")
        print("  2. Use an online HTML-to-PDF converter")
        print("  3. Install wkhtmltopdf system package")
        return False

if __name__ == "__main__":
    success = generate_pdf()
    sys.exit(0 if success else 1)
