#!/bin/bash
# Convert the master Markdown to DOCX using Pandoc.
# Prerequisites: install pandoc. For PDF, install a PDF engine like wkhtmltopdf or TeX Live (pdflatex).

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
MASTER_MD="$ROOT_DIR/templates/contracts/Proposal-Contract-Templates-Pack_MASTER.md"
OUT_DIR="$ROOT_DIR/dist"
OUT_DOCX="$OUT_DIR/Proposal-Contract-Templates-Pack.docx"

mkdir -p "$OUT_DIR"

# DOCX (recommended)
pandoc -s "$MASTER_MD" -o "$OUT_DOCX" --resource-path="$ROOT_DIR" \
  --metadata title="Proposal & Contract Templates Pack"

echo "Built: $OUT_DOCX"
