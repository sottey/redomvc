# redomvc

`redomvc` is a simple CLI tool written in Go for checking domain availability using the [Name.com API](https://www.name.com/developer/documentation). It supports single-domain checks as well as bulk checking from a file, with configurable concurrency and rate limiting.

## 🚀 Features

- Check availability of a single domain
- Check multiple domains from a file
- Show registration price for available domains
- Supports API key/username config via flags or config file
- Configurable concurrency and request delay

---

## 📦 Installation

```bash
go build -o redomvc .
```

---

## 🔧 Configuration

You can either pass credentials and settings via flags or save them in a Viper-compatible config file (e.g., `config.yaml` or `.redomvc.yaml`).

### Supported Config Keys

- `username` – your Name.com username
- `token` – your Name.com API token
- `file` – path to file containing domains to check
- `workers` – number of concurrent checks (default: 5)
- `delay` – delay between requests in milliseconds (default: 250)
- `api` – Name.com API base URL (default: `https://api.name.com/v4/domains:checkAvailability`)

---

## 🧪 Usage

### Check a single domain:

```bash
./redomvc example.com
```

### Check domains from a file:

```bash
./redomvc -f domains.txt
```

### With flags:

```bash
./redomvc -f domains.txt -u your_username -t your_token -w 10 -d 500
```

---

## 📁 Example Config File (`.redomvc.yaml`)

```yaml
username: your_username
token: your_token
workers: 10
delay: 500
api: https://api.name.com/v4/domains:checkAvailability
```

---

## 📝 License

MIT
