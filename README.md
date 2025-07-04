# push2Storage

A simple Go package for uploading files to cloud storage providers. Currently supports AWS S3, with more providers coming soon.

## Features
- Minimal API: Upload files with a single function call
- Extensible: Support for additional cloud providers planned

## Quick Start
1. Install:
   ```sh
   go get github.com/ayush5588/push2Storage/pkg/upload
   ```
2. Import and use:
   ```go
   import "github.com/ayush5588/push2Storage/pkg/upload"

   // Prepare your AWS credentials and config
   creds := map[string]string{
       "accessKey":    "<your-access-key>",
       "secretKeyID":  "<your-secretKeyID>",
       "bucket":       "<bucket-name>",
       "region":       "<region>",
   }

   // Upload a file
   resp := upload.Upload("aws", creds, "full-path-to-file")
   fmt.Println(resp)
   ```
3. See the `examples/` folder for a complete usage example, including integration with an HTTP server.

## Supported Providers
- AWS S3

---
Contributions and feedback are welcome!

