all:

install:
	cd snappie && go get && go build && go install
	support/build_launchctl > ~/Library/LaunchAgents/pl.bunsch.snappie.plist

uninstall:
	rm -rf ~/Library/LaunchAgents/pl.bunsch.snappie.plist
