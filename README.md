# k2

## Installation

k2 requires the Go 1.18 compiler and toolchain in order to build properly.
In addition, it depends on the `google/uuid` library for generating UUIDs,
and `julienschmidt/httprouter` for handling incoming requests.

Before building, a locale must be choosen. Not doing so will result in an
error during compilation. The section 'Locale options' contains further
information and instructions for installing locales.

After having created a `locale.go` file, run

    go get k2
    go install

This will automatically resolve and install all dependencies, compile the
binary, as well as install it to your GOPATH.

Finally, k2 requires a database file to be initialized. This can be done by
creating `data.json` containing `{}`.

## Locale options

A list of preconfigured locale files can be found in `locales/`. To select one
for your local k2 install, you can copy or symlink the chosen file to `locale.go`.

For example, for symlinking (the recommended option), simply run:

    ln -sf locales/YOUR_LOCALE.go locale.go

After any changes to the locale configuration, k2 must be recompiled and restarted.
