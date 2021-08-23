# MOOMIN

It uses Cocoa/WebKit on macOS, gtk-webkit2 on Linux and Edge on Windows 10.

## Description

Application start webview

## Install

Install golang https://golang.org/doc/install

**for linux:**

```
sudo apt-get install libwebkit2gtk-4.0-dev
sudo apt-get -y install make
```

## Build Setup frontend

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:3000
$ npm run dev

# build for production and launch server
$ npm run build
$ npm run start

# generate static project
$ npm run generate
```

## Build binary app

**linux:**

```bash
make build-deb
```

**darwin:**

Modify assets/Info.plist & icon.icns

```bash darwin
make build-darwin
```

**window:**

```
GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-H windowsgui" -v -o bin/$(BINARY_NAME)-win32-amd64.exe .
```


For detailed explanation on how things work, check out [Nuxt.js docs](https://nuxtjs.org).

## REFS:

- __nuxt/module:__ https://modules.nuxtjs.org/
- __nuxt/http:__ https://http.nuxtjs.org/
- __nuxt/pwa:__ https://pwa.nuxtjs.org/
- __nuxt/device:__ https://github.com/nuxt-community/device-module
- __@stephendltg/e-bus__ : https://www.npmjs.com/package/@stephendltg/e-bus
- __nuxt-vuex-localstorage__: https://www.npmjs.com/package/nuxt-vuex-localstorage
- __Vue i18n:__ https://kazupon.github.io/vue-i18n/

## INSTALL AND REMOVE DEB

```
sudo dpkg -i nom_du_paquet.deb
sudo apt-get remove nom_du_paquet
```
