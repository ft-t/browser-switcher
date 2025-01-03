#!/bin/bash
RAW_VERSION=${RAW_VERSION:-1.0.1}
TARGET_FOLDER=browser-switcher_$RAW_VERSION

mkdir -p $TARGET_FOLDER/{DEBIAN,usr/local/bin,usr/share/applications}

cat <<EOF > $TARGET_FOLDER/DEBIAN/control
Package: browser-switcher
Version: $RAW_VERSION
Architecture: all
Maintainer: FT-T <your.email@example.com>
Description: A tool to switch between browsers.
Priority: optional
Depends: xdg-utils
EOF

cp ../dist/linux/browser-switcher $TARGET_FOLDER/usr/local/bin/browser-switcher
cp ../dist/linux/browser-switcher-proxied $TARGET_FOLDER/usr/local/bin/browser-switcher-proxied
cp ../dist/linux/browser-switcher.desktop $TARGET_FOLDER/usr/share/applications/browser-switcher.desktop

cat <<EOF > browser-switcher_$RAW_VERSION/DEBIAN/postinst
#!/bin/bash
set -e

# Create required directories
mkdir -p ~/BrowserSwitcher

xdg-mime default browser-switcher.desktop x-scheme-handler/http
xdg-mime default browser-switcher.desktop x-scheme-handler/https
xdg-settings set default-web-browser browser-switcher.desktop

update-alternatives --install /usr/bin/x-www-browser x-www-browser /usr/local/bin/browser-switcher 400
update-alternatives --set x-www-browser /usr/local/bin/browser-switcher

xdg-settings set default-web-browser browser-switcher.desktop
EOF

chmod 755 browser-switcher_$RAW_VERSION/DEBIAN/postinst

dpkg-deb --build $TARGET_FOLDER
