# Examples

This example repository demonstrates many use cases for UniDoc's UniOffice Go library. Example code should make
it easy for users to get started with UniOffice and address specific use cases. Feel free to add to this by submitting
a pull request.

## Structure
The folder hierarchy is as follows:

- `document/` folder contains examples for creating and processing Word DOCX document files.
- `spreadsheet/` folder contains examples for creating and processing Excel XLSX spreadsheet files.
- `presentation/` folder contains examples for creating and processing Powerpoint PPTX presentation files.
- `license/` folder contains examples for using metered api key license and offline license files.

## License codes
UniOffice requires license codes to operate, there are two options:
- Metered License API keys: Free ones can be obtained at https://cloud.unidoc.io
- Offline Perpetual codes: Can be purchased at https://unidoc.io/pricing

Most of the examples demonstrate loading the Metered License API keys through an environment
variable `UNIDOC_LICENSE_API_KEY`.

Examples for Offline Perpetual License Key loading can be found in the license subdirectory.

# UniOffice License Loading.

The examples here illustrate how to work with UniOffice license codes and keys.
There are two types of licenses.

## Offline License
Offline licenses are cryptography based and contain full signed information that is verified based on signatures without making any outbound connections,
hence the name "offline". This kind of license is suitable for users deploying OEM products to their customers or where there are strict restrictions
on outbound connections due to firewalls and/or compliance requirements.

## Metered License (API keys)
The metered license is the most convenient way to get started with UniDoc products and the Free tier enables a powerful way to get started for free.
Anyone can get a free metered API key by signing up on http://cloud.unidoc.io/

## Examples

- [license/metered/main.go](license/metered/main.go) Demonstrates how to load the Metered API license key and how to print out relevant information.
- [license/metered-non-persistent-cache/main.go](license/metered-non-persistent-cache/main.go) Demonstrates how to load the Metered API license key for instances that not having persistent storage for usages cache and print out relevant information. When working with short-lived containers like docker or kubernetes instances usually it's doesn't have persistent storage location and will be destroyed after being idle.
- [license/offline/main.go](license/offline/main.go) Demonstrates how to print out information about the license after loading an offline license key.
- [license/usage-logs/main.go](license/usage-logs/main.go) This example shows how to display license usage logs for read and editing simple docx file.

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

### Build all examples

Simply run the build script which builds all the binaries to subfolder `bin/`

```bash
$ ./build_examples.sh
```
