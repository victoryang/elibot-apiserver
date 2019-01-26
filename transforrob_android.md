## armv7

### make
	make INSTALL_ROOT=${BUILD_ROOT}\adroid_build install

## modify code
	remove all code about virtualkeyboard
	clean up all install target code from *.pro

### jbiedit
	copy qmldir and libjbieditor.so to ${QT}\android_armv7\qml\Elibot\JobEditor

### robsimuitem
	remove c++ std math related dependency
	QString text = QString::fromStdString(pEle->GetText());
        color = text.toULong();

    copy qmldir and libqrobsimuitemplugin.so to ${QT}\android_armv7\qml\Elibot\QRobSimuItem
