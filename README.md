# Malware Scanner (Go)
A directory fuzzer that scans each file in the target directory using Virus' total public api, 
returning in a PASS/FAIL result for each file.

Uses Go's standard net/http library

## Features
- Uploads a file to VirusTotal using their API  
- Retrieves the file's scan status  
- Command-line support for specifying file input  

## Prerequisites
- Go 1.18+ installed  
- A valid [VirusTotal API key](https://www.virustotal.com/gui/my-apikey)  

## Installation
Clone this repository and navigate to the directory:  
```sh
git clone https://github.com/yourusername/virustotal-scanner-go.git
cd virustotal-scanner-go
```

## Usage
1. **Set up your VirusTotal API key**  
   Replace the hardcoded API key in `trial.go` with your own:
   ```go
   req.Header.Add("x-apikey", "YOUR_API_KEY_HERE")
   ```

2. **Run the script**  
   ```sh
   go run trial.go <file_to_scan>
   ```
   If no file is provided, it defaults to `sample.txt`.  

## Example Output
```sh
File [example.txt] is being scanned
File hash ID: abc123xyz
Getting status of file...
{ "data": { "attributes": { "last_analysis_stats": { ... }}}}
```

## License
MIT License.

