# Renamer Utility

The Renamer Utility is a command-line tool written in Go that allows you to rename files in a directory based on different patterns. It provides a flexible way to rename files in bulk without manual intervention.

## Getting Started

To use the Renamer Utility, follow these steps:

**Clone the Repository**: 
   ```sh
   git clone https://github.com/yourusername/renamer.git
   cd renamer
   ```
**Build the Project**:
```sh
go build
```
**Run the Utility**:
```sh
./renamer -pattern default
```
You can use the -dir flag to specify a different root directory.
**Available Patterns**: The utility supports two renaming patterns:
default: Renames files with numeric indices in the format "(index of total).ext".
custom: Renames files with numeric indices and a custom prefix in the format "index - prefix.ext".
