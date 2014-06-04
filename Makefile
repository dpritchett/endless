APPNAME:=endless

default: clean
	script/release

clean:
	rm -f ./$(APPNAME).*
