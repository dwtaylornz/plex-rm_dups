FROM mcr.microsoft.com/powershell:latest
LABEL maintainer="Darren <dwtaylornz@gmail.com>"

ADD plex-rm_dups.ps1 . 

CMD Set-ExecutionPolicy Unrestricted
CMD ./plex-rm_dups.ps1
