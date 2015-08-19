all:
	go get
	go build

launchctl:
	support/build_launchctl > ~/Library/LaunchAgents/pl.bunsch.snappie.plist
