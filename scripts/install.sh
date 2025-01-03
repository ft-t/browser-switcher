#!/bin/bash
sudo mkdir -p /etc/browser-switcher
mkdir -p ~/BrowserSwitcher

sudo cp -f browser-switcher /usr/local/bin/browser-switcher && sudo chmod 777 /usr/local/bin/browser-switcher
sudo cp -f browser-switcher-proxied /usr/local/bin/browser-switcher-proxied && sudo chmod 777 /usr/local/bin/browser-switcher-proxied

sudo cp -f browser-switcher.desktop /usr/share/applications/browser-switcher.desktop
sudo xdg-mime default browser-switcher.desktop x-scheme-handler/http
sudo xdg-mime default browser-switcher.desktop x-scheme-handler/https

sudo update-alternatives --install /usr/bin/x-www-browser x-www-browser /usr/local/bin/browser-switcher 400
sudo update-alternatives --config x-www-browser

sudo xdg-settings set default-web-browser browser-switcher.desktop
