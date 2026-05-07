all: android windows linux scanner
android:
	cd android && ./gradlew assembleRelease
windows:
	cd windows && go build -o ArgoTunnel.exe
linux:
	cd linux && go build -o ArgoTunnel
scanner:
	cd scanner-cli && go build -o argoscan
clean:
	rm -f windows/ArgoTunnel.exe linux/ArgoTunnel scanner-cli/argoscan
