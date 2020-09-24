CXX = gcc
CXXFLAGS = -std=c++11 -Wall -g -O3
INCLUDE  = -I ./l2
TARGETL2   = libl2.so
LIBPATH  = ./libs/
SRCL2 = ./l2/l2.cxx

.PHONY: buildl2 clean
buildl2:
	$(CXX) -shared -fPIC $(CXXFLAGS) -o $(TARGETL2) $(SRCL2)
	mv $(TARGETL2) $(LIBPATH)

clean:
	rm -f $(LIBPATH)*