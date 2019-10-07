gts-notion
---

Fetch race data from GTSport server and submit to my Notion page.
Just published code for fun, but this is useless for nobody but me.

## Usage for myself

(because I will forget after a week..)

Install Xcode11+ Swift 5.1
Install Golang

Set `NOTION_AUTH_TOKEN` as environment variable. ([seealso](https://presstige.io/p/Using-Notion-API-Go-client-2567fcfa8f7a4ed4bdf6f6ec9298d34a#6e20956c-20a2-4519-b91d-a9fb151d88d3))

```
make prepare
make run
```

Localize data (e.g. course name) is extracted automatically using a dumb shell script.

## Credits

- [kjk/notionapi](https://github.com/kjk/notionapi)

## LICENSE

Whatever
