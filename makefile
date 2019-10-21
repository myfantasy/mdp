
pre_test:
	if [ -d "test/" ]; then \
		rm -R test/; \
	fi;	
	mkdir test/	

t: pre_test
	go test .