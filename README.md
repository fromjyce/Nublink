# Nublink 🔒

A secure, self-hosted file-sharing CLI with self-destructing links. Built with Go.

## Features ✨

- **End-to-end encryption** (AES-GCM)
- **Self-destructing links** (time-based or one-time download)
- **No cloud dependencies** - your machine is the server
- **TLS-secured transfers** (auto-generated certificates)
- **Lightweight** (~10MB binary)

## Installation ⚡

### From Source
```bash
git clone https://github.com/fromjyce/Nublink.git
cd Nublink
go build -o nublink ./cmd/nublink/
```

### Using `go install`
```bash
go install github.com/fromjyce/Nublink/cmd/nublink@latest
```

### Share a file
```bash
# Time-based expiration (1 hour)
nublink share --file secret.pdf --expire 1h

# One-time download
nublink share --file confidential.txt --once
```

### Access shared files
1. Open the generated link in any browser:
   ```
   https://[YOUR_IP]:8443/download/abc123-def456
   ```
2. Bypass the SSL warning (self-signed cert)
3. File auto-deletes after download/expiry

## Network Sharing 🌐
To allow others on your local network to access files:

1. Find your IP:
   ```bash
   # Mac/Linux
   ifconfig | grep "inet " | grep -v 127.0.0.1

   # Windows
   ipconfig | findstr "IPv4"
   ```
2. Share links using your IP instead of `localhost`:
   ```
   https://[YOUR_IP]:8443/download/abc123-def456
   ```

## Technical Details 🔧

### Security
| Component           | Implementation         |
|---------------------|------------------------|
| Transport Security  | TLS 1.3 (self-signed)  |
| Encryption at rest  | AES-256-GCM            |
| Key Management      | Per-file random keys   |

### File Storage
- Encrypted files stored in: `~/.nublink/files/`
- Metadata stored in: `~/.nublink/*.meta`
- Auto-cleaned after expiration

## Development 🛠️

### Prerequisites
- Go 1.21+
- OpenSSL (for cert generation)

## Contact
If you come across any mistakes in the programs or have any suggestions for improvement, please feel free to contact me <jaya2004kra@gmail.com>. I appreciate any feedback that can help me improve my coding skills

## License
All the programs in this repository are licensed under the MIT License. You can use them for educational purposes and modify them as per your requirements. ***However, I do not take any responsibility for the accuracy or reliability of the programs.***

## MY SOCIAL PROFILES:
### [LINKEDIN](https://www.linkedin.com/in/jayashrek/)