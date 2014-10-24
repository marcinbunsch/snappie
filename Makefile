all:
	go get
	go build

launchctl:
	launchctl load support/pl.bunsch.snappie.plist
