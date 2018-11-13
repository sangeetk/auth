FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD auth auth
EXPOSE 8080
ENTRYPOINT ["/auth", "--key=NEW", "--dbpath=/db"]
