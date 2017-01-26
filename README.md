# Rebuild

Rebuild is a command line tool that builds and (re)starts your web application everytime you save a Go, template or specified file type.

## Installation

    go get -u github.com/ginodeis/rebuild

## Usage

    cd /path/to/myapp
    rebuild

Rebuild will watch for file events, and every time you create/modifiy/delete a file it will build and restart the application.
If `go build` returns an error, it will log it in the tmp folder.

`rebuild` uses `./runner.conf` for configuration by default, but you may specify an alternative config filepath using `-c`:

    rebuild -c other.conf

Here is a sample config file with the default settings:

    root:              .
    watch_paths        ., ../code
    tmp_path:          ./tmp
    build_name:        runner-build
    build_log:         runner-build-errors.log
    valid_ext:         .go, .tpl, .tmpl, .html
    build_delay:       600
    runner_args:       server
    colors:            1
    log_color_main:    cyan
    log_color_build:   yellow
    log_color_runner:  green
    log_color_watcher: magenta
    log_color_app:

## Thanks to

* [Andrea Franz](http://gravityblast.com) - Original fresh project
