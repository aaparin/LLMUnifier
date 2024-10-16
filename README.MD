# LLM Unifier

LLM Unifier is a command-line tool written in Go that combines the content of multiple files into a single output, based on a configuration specified in a YAML file. It's designed to help prepare and consolidate code files for Large Language Model (LLM) training or analysis.

## Features

- Combine files from multiple directories
- Support for wildcard patterns to select files
- Ability to exclude specific files or directories
- Customizable comment symbol for different programming languages
- Flexible configuration through a YAML file
- Relative path comments for each included file in the output

## Installation

1. Ensure you have Go installed (version 1.16 or higher).
2. Clone the repository:
   ```
   git clone https://github.com/aaparin/LLMUnifier.git
   cd LLMUnifier
   ```
3. Initialize the Go module:
   ```
   go mod init github.com/yourusername/LLMUnifier
   ```
4. Install dependencies:
   ```
   go get gopkg.in/yaml.v2
   ```

## Configuration

Create a `config.yaml` file in the root directory of the project with the following structure:

```yaml
root_dir: /path/to/your/project
comment_symbol: "#"
output_file: unified_code.txt
paths:
  - app/*.php
  - index.php
  - !app/Console/
```

- `root_dir`: The root directory of your project
- `comment_symbol`: The symbol used for comments in the output file
- `output_file`: The name of the output file
- `paths`: List of paths to include/exclude files

## Usage

1. Configure the `config.yaml` according to your requirements.
2. Run the program:
   ```
   go run main.go
   ```
3. The program will create the file specified in `output_file`, containing the unified code from the specified files.

## Output Example

```
# app/index.php
<?php
echo 'Hello from index.php';
?>

# app/config.php
<?php
$config = [
    'debug' => true,
    'environment' => 'development'
];
?>
```

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change. Please ensure to update tests as appropriate.

## Support

If you encounter any problems or have any questions, please open an issue on the GitHub repository.