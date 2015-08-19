all:
	go get
	go build

install:
	support/build_launchctl > ~/Library/LaunchAgents/pl.bunsch.snappie.plist

uninstall:
	rm -rf ~/Library/LaunchAgents/pl.bunsch.snappie.plist
