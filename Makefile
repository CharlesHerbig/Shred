default: testfile lockfile rootfile emptyfile

testfile: testfile-1024
	@cat testfile-1024 > $@
	@cat testfile-1024 >> $@
	@echo whee >> $@

lockfile: testfile-1024
	@cat testfile-1024 > $@
	@cat testfile-1024 >> $@
	@echo whee >> $@
	@chmod a-w $@

rootfile: testfile-1024
	@cat testfile-1024 > $@
	@cat testfile-1024 >> $@
	@echo whee >> $@
	sudo chown root:root $@

emptyfile:
	@touch $@

clean:
	@rm -f testfile lockfile rootfile emptyfile
