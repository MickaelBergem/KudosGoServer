FROM scratch

ADD initdb.sql /
ADD main /

CMD ["/main"]
