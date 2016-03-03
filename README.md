## The-Startup-Stack cli

[The Startup Stack](http://http://the-startup-stack.com/) command line tool.

This project is a CLI for the startup stack. It helps you work better with all
the moving parts.

## Available commands

### Project

All commands related to the project directory and templates.

### Create

Creates the project template under the current directory.

#### Flags

* `--directory-name`
* `--project-name`

* Clones the [template repository](https://github.com/the-startup-stack/chef-repo-template) to `~/.the-startup-stack`
* Creates a directory called `{project-name}` under the current directory
* Copies the necesary directories and files from the template to the directory
  and renames to your name

`stack project create --directory-name chef --project-name your-company`

## License

The MIT License (MIT)
Copyright (c) 2016 Avi Tzurel

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
