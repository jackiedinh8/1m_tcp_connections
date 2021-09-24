CC          = cc
DEPTH       = src
TARGETDIR   = obj
CFLAGS      =
OBJDIR      = $(DEPTH)/$(TARGETDIR)
INCDIR      = $(DEPTH)/$(TARGETDIR)
LIBST       = $(OBJDIR)/libst.a
HEADER      = $(INCDIR)/st.h

LIBRESOLV   =
EXTRALIBS   =
ECHOSVR    = echosvr


all: $(ECHOSVR)

$(ECHOSVR): echosvr.c $(LIBST) $(HEADER)
	$(CC) $(CFLAGS) -I$(INCDIR) echosvr.c $(LIBST) $(EXTRALIBS) -o $@

$(OBJDIR)/%.o: %.c
	$(CC) $(CFLAGS) -I$(INCDIR) -c $< -o $@

#.DEFAULT:
#	@cd $(DEPTH); $(MAKE) $@
$(LIBST):
	@cd $(DEPTH); $(MAKE) linux-optimized

clean:
	rm echosvr
	@cd $(DEPTH); $(MAKE) clean
