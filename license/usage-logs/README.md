### Metered License Key Usage Logs
When using unioffice on metered api key it is possible to see the usage logs of license key in every run. To enable this set the `license.SetMeteredKeyUsageLogVerboseMode` to `true` and to print out into console set the `logger.SetLogger` to Info or higher. as follows. This is available for metered api key only.

```go
// Set the log level to info or higher
logger.SetLogger(logger.NewConsoleLogger(logger.LogLevelInfo))

// Enable the verbose mode logging
license.SetMeteredKeyUsageLogVerboseMode(true)
```
Sample output 
```bash
[INFO]  metered.go:674 2024-09-26 20:32:22.607531 +0300 +03 m=+0.412134293 | File grocery_list.docx | Ref: dr17f8db | document.Read | 1 credit(s) used
[INFO]  metered.go:506 2024-09-26 20:32:22.653044 +0300 +03 m=+0.457648876 Ref: dr17f8db | document:d.Save | No credit used
```

#### Examples 
- [Modify docx file](main.go) This example shows how to display license usage logs for a editing simple docx file.


