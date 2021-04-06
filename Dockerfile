FROM mcr.microsoft.com/powershell:latest
LABEL maintainer="Darren <darren.taylor@concepts.co.nz>"

# ADD index.html /www/index.html

# EXPOSE 8000
# HEALTHCHECK CMD nc -z localhost 8000

# description
# CMD trap "exit 0;" TERM INT; httpd -p 8000 -h /www -f & wait
