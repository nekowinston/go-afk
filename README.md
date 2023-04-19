# go-afk

Rewrite of [@andreasgrafen](https://github.com/andreasgrafen)'s PHP script to
generate AFK image banners.

## API

This repository is hosted on [afk.winston.sh](https://afk.winston.sh).
The image is generated from query parameters:

![Example image](https://afk.winston.sh/?f=frappe&c=pink&t=Hello%20Gophers!&i=1)

```
https://afk.winston.sh/?f=frappe&c=pink&t=Hello%20Gophers!&i=1
```

- `?f=frappe`: flavour of [Catppuccin](https://github.com/catppuccin/catppuccin)
  to use for the foreground text colour.
- `&c=pink`: colour of Catppuccin to use
- `&t=Hello%20Gophers`: text to be rendered. has to be URL-encoded
  (e.g. `%20` for spaces)
- `&i=1`: when present, render the font in italics.
