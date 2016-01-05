all:
	go get
	go build

install:
	support/build_launchctl > ~/Library/LaunchAgents/pl.bunsch.snappie.plist
	launchctl load ~/Library/LaunchAgents/pl.bunsch.snappie.plist

uninstall:
	launchctl remove pl.bunsch.snappie
	rm -rf ~/Library/LaunchAgents/pl.bunsch.snappie.plist
