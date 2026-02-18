# Breeze (formerly project-phantom) - just browse, privately

### _what? why?_ - 🔗 [here](https://www.notion.so/project-phantom-226b6acd0a1480ac9d73e5a0c26d8f5e)
### _infra? tech? architecture?_ - 🔗 [here](https://www.notion.so/project-phantom-Technical-Architecture-226b6acd0a148038a339fe16f0150114)

---

## Proposal & Contract Templates Pack (Nigeria/Lagos) — chbuzorikenna

This repo contains a master Markdown document for proposals and legal templates tailored to Nigeria/Lagos. You can export to DOCX using Pandoc or via the GitHub Action.

### Files
- `templates/contracts/Proposal-Contract-Templates-Pack_MASTER.md` — main document (logo embedded)
- `templates/notion/*` — Notion-friendly pages
- `scripts/convert.sh` — build DOCX
- `dist/Proposal-Contract-Templates-Pack.docx` — generated output (after build)
- `assets/chbuzorikenna-logo.svg` — replace with your own logo if desired

### Build locally
```bash
bash scripts/convert.sh
```

### CI build
Push to master or run the workflow manually to produce the DOCX and download it from Actions → Artifacts.