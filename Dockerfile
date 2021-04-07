FROM mcr.microsoft.com/powershell:latest
LABEL maintainer="Darren <dwtaylornz@gmail.com>"

ADD plex-rm_dups.ps1 .

CMD powershell .\plex-rm_dups.ps1
