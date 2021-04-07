FROM mcr.microsoft.com/powershell:latest
LABEL maintainer="Darren <dwtaylornz@gmail.com>"

ADD plex-rm_dups.ps1 /scripts 
RUN chmod +x /scripts/plex-rm_dups.ps1

CMD /scripts/plex-rm_dups.ps1
