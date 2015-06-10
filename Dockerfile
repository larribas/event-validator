FROM scratch
EXPOSE 8190
COPY event_validator /
ENTRYPOINT ["/event_validator"]