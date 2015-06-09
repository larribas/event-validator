FROM scratch
EXPOSE 7281
COPY event_validator /
ENTRYPOINT ["/event_validator"]