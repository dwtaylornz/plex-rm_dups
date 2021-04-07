FROM mcr.microsoft.com/powershell:latest
LABEL maintainer="Darren <dwtaylornz@gmail.com>"

ADD plex-rm_dups.ps1 \script\plex-rm_dups.ps1

CMD pwsh -File \script\plex-rm_dups.ps1
