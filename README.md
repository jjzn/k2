# k2

## Locale options

A list of preconfigured locale files can be found in `locals/`. To select one
for your local k2 install, you can copy or symlink the chosen file to `local.go`.

For example, for symlinking (the recommended option), simply run:

    ln -sf locals/YOUR_LOCALE.go local.go

After any changes to the locale configuration, k2 must be recompiled and restarted.
