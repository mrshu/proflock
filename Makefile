CC=go
INSTALL=install -c
DEST=/usr/bin

proflock: proflock.go
	$(CC) build ./proflock.go

install: proflock
	$(INSTALL) proflock $(DEST)/proflock

clean:
	rm -rf proflock
